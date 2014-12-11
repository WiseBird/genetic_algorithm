package genetic_algorithm

import (
	"bytes"
	"fmt"
	log "github.com/cihub/seelog"
	"math"
	"time"
)

type StatisticsDataDefault interface {
	Generations() int
	Duration() time.Duration
	Durations() *HierarchicalDuration
	MinCost() float64
	MinCosts() []float64
	GenerationsWithoutImprovements() int
	MinCostsVar() float64
	MeanCost() float64
	MeanCosts() []float64
	WorstCost() float64
	WorstCosts() []float64
}

// Default realization of StatisticsInterface
type StatisticsDefault struct {
	started              bool
	durationTracker      *durationTracker
	durationTrackerStack [][]string
	hierarchy            *HierarchicalDuration

	generations int

	minCost              float64
	gensWoImprv          int
	prevDifferentMinCost float64
	minCosts             []float64
	minCostsVar          float64

	meanCost  float64
	meanCosts []float64

	worstCost  float64
	worstCosts []float64

	options *StatisticsDefaultOptions
}

func NewStatisticsDefault(options StatisticsOptionsInterface) StatisticsInterface {
	opts, ok := options.(*StatisticsDefaultOptions)
	if !ok {
		panic("Expects instance of StatisticsOptionsInterface")
	}

	statistics := new(StatisticsDefault)

	statistics.generations = -1
	statistics.options = opts

	return statistics
}
func (statistics *StatisticsDefault) Start(keys ...string) {
	if !statistics.options.trackDurations {
		return
	}

	if statistics.durationTracker == nil {
		statistics.durationTracker = newDurationTracker()
		statistics.durationTrackerStack = make([][]string, 0)
	}

	statistics.started = true
	statistics.hierarchy = nil

	if len(keys) != 0 {
		statistics.durationTrackerStack = append(statistics.durationTrackerStack, keys)
	}

	tracker := statistics.durationTracker
	for _, key := range keys {
		tracker = tracker.child(key)
	}
	tracker.start()
}
func (statistics *StatisticsDefault) End() {
	if !statistics.options.trackDurations {
		return
	}

	if !statistics.started {
		return
	}

	stackLen := len(statistics.durationTrackerStack)
	if stackLen == 0 {
		statistics.started = false
		statistics.durationTracker.end()
		return
	}

	keys := statistics.durationTrackerStack[stackLen-1]
	statistics.durationTrackerStack = statistics.durationTrackerStack[:stackLen-1]

	tracker := statistics.durationTracker
	for _, key := range keys {
		tracker = tracker.child(key)
	}
	tracker.end()
}

// Expects sorted population
func (statistics *StatisticsDefault) OnGeneration(population Chromosomes) {
	statistics.generations++

	if len(population) == 0 {
		panic("Population is empty")
	}

	statistics.minCost = population[0].Cost()
	log.Tracef("MinCost %v", statistics.minCost)

	if statistics.options.trackMinCosts {
		statistics.minCosts = append(statistics.minCosts, population[0].Cost())
		log.Tracef("MinCosts %v", statistics.minCosts)
	}
	if statistics.options.trackGensWoImprv {
		if statistics.generations == 0 || statistics.prevDifferentMinCost != statistics.minCost {
			statistics.prevDifferentMinCost = statistics.minCost
			statistics.gensWoImprv = 0
		} else {
			statistics.gensWoImprv++
		}
		log.Tracef("MinCostAge %v", statistics.gensWoImprv)
	}
	if statistics.options.trackMinCostsVar {
		if statistics.generations == 0 || statistics.generations == statistics.gensWoImprv {
			statistics.minCostsVar = math.NaN()
		} else {
			statistics.minCostsVar = pvarianceFloat64(statistics.minCosts)
		}
		log.Tracef("MinCostsVar %v", statistics.minCostsVar)
	}

	if statistics.options.trackMeanCost {
		statistics.meanCost = population.MeanCost()
		log.Tracef("MeanCost %v", statistics.meanCost)
	}
	if statistics.options.trackMeanCosts {
		var mean float64
		if statistics.options.trackMeanCost {
			mean = statistics.meanCost
		} else {
			mean = population.MeanCost()
		}

		statistics.meanCosts = append(statistics.meanCosts, mean)
		log.Tracef("MeanCosts %v", statistics.meanCosts)
	}

	if statistics.options.trackWorstCost {
		statistics.worstCost = population[len(population)-1].Cost()
		log.Tracef("WorstCost %v", statistics.worstCost)
	}
	if statistics.options.trackWorstCosts {
		statistics.worstCosts = append(statistics.worstCosts, population[len(population)-1].Cost())
		log.Tracef("WorstCosts %v", statistics.worstCosts)
	}
}

// Number of generations
func (statistics *StatisticsDefault) Generations() int {
	return statistics.generations
}

// Returns duration of optimization.
// If optimization are in progress then returns time elapsed from start.
func (statistics *StatisticsDefault) Duration() time.Duration {
	if !statistics.options.trackDurations {
		return time.Duration(0)
	}

	if !statistics.started {
		return statistics.durationTracker.elapsed[0]
	}

	return time.Since(statistics.durationTracker.startTime)
}

// Returns entire tracked hierarchy of duration
func (statistics *StatisticsDefault) Durations() *HierarchicalDuration {
	if statistics.hierarchy == nil {
		statistics.hierarchy = statistics.durationTracker.toHierarchy()
	}

	return statistics.hierarchy
}

// Min cost for last iteration
func (statistics *StatisticsDefault) MinCost() float64 {
	return statistics.minCost
}

// Min cost for each iteration
// Len would be `Generations() + 1` because of initial value
func (statistics *StatisticsDefault) MinCosts() []float64 {
	return statistics.minCosts
}

// Number of generations during which the min cost remains unchanged
func (statistics *StatisticsDefault) GenerationsWithoutImprovements() int {
	return statistics.gensWoImprv
}

// Variance of min costs
// Variance equals NaN until two different values of MinCost are obtained
func (statistics *StatisticsDefault) MinCostsVar() float64 {
	return statistics.minCostsVar
}

// Mean cost of last iteration
func (statistics *StatisticsDefault) MeanCost() float64 {
	return statistics.meanCost
}

// Mean cost of each iteration
// Len would be `Iterations() + 1` because of initial value
func (statistics *StatisticsDefault) MeanCosts() []float64 {
	return statistics.meanCosts
}

// Worst cost for last iteration
func (statistics *StatisticsDefault) WorstCost() float64 {
	return statistics.worstCost
}

// Worst cost for each iteration
// Len would be `Iterations() + 1` because of initial value
func (statistics *StatisticsDefault) WorstCosts() []float64 {
	return statistics.worstCosts
}

func (statistics *StatisticsDefault) Data() StatisticsDataInterface {
	return statistics
}

type durationTracker struct {
	startTime time.Time
	elapsed   []time.Duration

	children map[string]*durationTracker
}

func newDurationTracker() *durationTracker {
	tracker := new(durationTracker)

	return tracker
}
func (tracker *durationTracker) start() {
	tracker.startTime = time.Now()
}
func (tracker *durationTracker) end() {
	if tracker.elapsed == nil {
		tracker.elapsed = make([]time.Duration, 0, 1)
	}

	tracker.elapsed = append(tracker.elapsed, time.Since(tracker.startTime))
}
func (tracker *durationTracker) child(name string) *durationTracker {
	if tracker.children == nil {
		tracker.children = make(map[string]*durationTracker, 1)
	}

	child, ok := tracker.children[name]
	if !ok {
		child = newDurationTracker()
		tracker.children[name] = child
	}

	return child
}
func (tracker *durationTracker) toHierarchy() *HierarchicalDuration {
	return tracker.toNamedHierarchy("total")
}
func (tracker *durationTracker) toNamedHierarchy(name string) *HierarchicalDuration {
	hierarchy := newHierarchicalDuration(name)

	for _, elapsed := range tracker.elapsed {
		hierarchy.add(elapsed)
	}

	for childName, child := range tracker.children {
		hierarchy.addChild(child.toNamedHierarchy(childName))
	}

	return hierarchy
}

type HierarchicalDuration struct {
	Name     string
	Duration time.Duration
	Calls    int

	Children map[string]*HierarchicalDuration
}

func newHierarchicalDuration(name string) *HierarchicalDuration {
	hierarchy := new(HierarchicalDuration)

	hierarchy.Name = name

	return hierarchy
}
func (hierarchy *HierarchicalDuration) add(duration time.Duration) {
	hierarchy.Calls++
	hierarchy.Duration += duration
}
func (hierarchy *HierarchicalDuration) addChild(child *HierarchicalDuration) {
	if hierarchy.Children == nil {
		hierarchy.Children = make(map[string]*HierarchicalDuration, 1)
	}

	hierarchy.Children[child.Name] = child
}
func (hierarchy *HierarchicalDuration) String() string {
	var buffer bytes.Buffer

	hierarchy.string(0, &buffer)

	return buffer.String()
}
func (hierarchy *HierarchicalDuration) string(indent int, buffer *bytes.Buffer) {
	for i := 0; i < indent; i++ {
		buffer.WriteString("\t")
	}

	buffer.WriteString(fmt.Sprintf("%s: %v", hierarchy.Name, hierarchy.Duration))

	if hierarchy.Calls != 0 && hierarchy.Calls != 1 {
		buffer.WriteString(fmt.Sprintf(" [%d x %v]",
			hierarchy.Calls, hierarchy.Duration/time.Duration(hierarchy.Calls)))
	}

	buffer.WriteString("\n")

	for _, child := range hierarchy.Children {
		child.string(indent+1, buffer)
	}
}
