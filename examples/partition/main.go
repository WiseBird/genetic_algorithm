package main

import (
	. "github.com/WiseBird/genetic_algorithm"
	"math"
	log "github.com/cihub/seelog"

    "code.google.com/p/plotinum/plot"
    "code.google.com/p/plotinum/plotter"
    "code.google.com/p/plotinum/plotutil"
)

var (
	list = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	finalSum = 178
	finalProduct = 3120
)

func cost(c ChromosomeInterface) float64 {
	bc := c.(*BinaryChromosome)
	sum := 0
	prod := 1

	genes := bc.Genes().(BinaryGenes)

	for i := 0; i < len(genes); i++ {
		if genes[i] {
			prod *= list[i]
		} else {
			sum += list[i]
		}
	}

	sumDiff := float64(sum - finalSum)
	prodDiff := float64(prod - finalProduct)

	return math.Abs(sumDiff) + math.Abs(prodDiff * prodDiff)
}

func main() {
	defer log.Flush()
	setupLogger()

	iterations := 1000

	statisticsOptions := NewStatisticsDefaultOptions().
			TrackMinCosts().
			TrackMeanCosts()
	statisticsAggregator := NewStatisticsDefaultAggregator(statisticsOptions)
	optimizer := createOptimizer(statisticsOptions)

	gatherStatistics(optimizer, statisticsAggregator, iterations)

	log.Warnf("Duration: %v", statisticsAggregator.Duration())
	log.Warnf("MinCosts: %v", statisticsAggregator.MinCosts())

	drawPlot(statisticsAggregator)
}

func createOptimizer(statisticsOptions StatisticsOptionsInterface) OptimizerInterface {
	popSize := 32
	chromSize := len(list)
	weedRate := 50.0
	mutationProb := 0.05
	generations := 200

	return NewIncrementalOptimizer().
		Initializer(BinaryRandomInitializerInstance).
		Weeder(NewSimpleWeeder(weedRate)).
		Selector(NewRouletteWheelCostWeightingSelector()).
		Breeder(NewOnePointBreeder(NewEmptyBinaryChromosome)).
		Mutator(NewBinaryMutator(mutationProb)).
		CostFunction(cost).
		StopCriterion(NewStopCriterionDefault().
			Max_Generations(generations).
			Min_Cost(0).
			Max_MinCostAge(15)).
		StatisticsOptions(statisticsOptions).
		PopSize(popSize).
		ChromSize(chromSize)
}
func gatherStatistics(optimizer OptimizerInterface, statisticsAggregator *StatisticsDefaultAggregator, iters int) {
	for i := 0; i < iters; i++ {
		_, statistics := optimizer.Optimize()
		statisticsAggregator.Aggregate(statistics)
	}
	statisticsAggregator.Compute()
}

func drawPlot(stats *StatisticsDefaultAggregator) {
    p, err := plot.New()
    if err != nil {
            panic(err)
    }

    p.Title.Text = "Plotutil example"
    p.X.Label.Text = "Gens"
    p.Y.Label.Text = "Cost"

    err = plotutil.AddLinePoints(p,
            "Mean", convertCostsToXYs(stats.MinCosts()),
            "Min", convertCostsToXYs(stats.MeanCosts()))
    if err != nil {
            panic(err)
    }

    if err := p.Save(8, 4, "points.png"); err != nil {
            panic(err)
    }
}

func convertCostsToXYs(costs []float64) plotter.XYs {
    pts := make(plotter.XYs, len(costs))
    for i, cost := range costs {
        pts[i].X = float64(i)
        if cost > 0 {
	        pts[i].Y = math.Log10(cost)
	    } else if cost < 0 {
	    	pts[i].Y = -1 * math.Log10(math.Abs(cost))
	    } else {
	    	pts[i].Y = 0
	    }
    }
    return pts
}

func setupLogger() {
	logger, err := log.LoggerFromConfigAsFile("seelog.xml")
	if err != nil {
	    panic(err)
	}

	log.ReplaceLogger(logger)
}