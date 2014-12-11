package genetic_algorithm

import (
	"fmt"
	"math/rand"
)

const (
	mutatorInvertExactLen      = 0
	mutatorInvertPercentageLen = 1
)

// Mutator selects some part of the chromosome and inverts it
type InvertMutator struct {
	probability           float64
	chromosomeConstructor EmptyChromosomeConstructor
	kind                  int
	fromExact             int
	toExact               int
	fromPercent           float64
	toPercent             float64

	temp GenesInterface
}

func newInvertMutator(probability float64, chromosomeConstructor EmptyChromosomeConstructor) *InvertMutator {
	if probability > 1 || probability < 0 {
		panic(fmt.Sprintf("Incorrect probability %v", probability))
	}

	mutator := new(InvertMutator)

	mutator.probability = probability
	mutator.chromosomeConstructor = chromosomeConstructor

	return mutator
}

// Creates mutator that will invert part of chromosome that lays on interval [from:to]
// Probability is applied to each chromosome.
func NewInvertExactIntervalMutator(probability float64, chromosomeConstructor EmptyChromosomeConstructor, from, to int) *InvertMutator {
	if from > to || from < 0 || to < 1 {
		panic(fmt.Sprintf("Incorrect interval [%d:%d]", from, to))
	}

	mutator := newInvertMutator(probability, chromosomeConstructor)

	mutator.kind = mutatorInvertExactLen
	mutator.fromExact = from
	mutator.toExact = to

	return mutator
}

// Creates mutator that will invert part of chromosome.Part's len defined by genesToInvert.
// Probability is applied to each chromosome.
func NewInvertExactMutator(probability float64, chromosomeConstructor EmptyChromosomeConstructor, genesToInvert int) *InvertMutator {
	return NewInvertExactIntervalMutator(probability, chromosomeConstructor, genesToInvert, genesToInvert)
}

// Creates mutator that will invert part of chromosome that lays on interval [from*chromLen : to*chromLen]
// Probability is applied to each chromosome.
func NewInvertPercentageIntervalMutator(probability float64, chromosomeConstructor EmptyChromosomeConstructor, from, to float64) *InvertMutator {
	if from > to || from < 0 || to < 0 || from > 1 || to > 1 {
		panic(fmt.Sprintf("Incorrect percentage interval [%v:%v]", from, to))
	}

	mutator := newInvertMutator(probability, chromosomeConstructor)

	mutator.kind = mutatorInvertPercentageLen
	mutator.fromPercent = from
	mutator.toPercent = to

	return mutator
}

// Creates mutator that will invert part of chromosome. Part's len calcs as percentage * chromLen
// Probability is applied to each chromosome.
func NewInvertPercentageMutator(probability float64, chromosomeConstructor EmptyChromosomeConstructor, percentage float64) *InvertMutator {
	return NewInvertPercentageIntervalMutator(probability, chromosomeConstructor, percentage, percentage)
}

func (mutator *InvertMutator) Mutate(population Chromosomes) {
	for _, chrom := range population {
		if mutator.probability < rand.Float64() {
			continue
		}

		genesLen := chrom.Genes().Len()
		intervalLen := mutator.getIntervalLen(genesLen)
		from, to := mutator.getInterval(genesLen, intervalLen)

		mutator.mutate(chrom.Genes(), from, to)
	}
}
func (mutator *InvertMutator) getIntervalLen(genesLen int) int {
	if mutator.kind == mutatorInvertExactLen {
		return rand.Intn(mutator.toExact-mutator.fromExact+1) + mutator.fromExact
	}

	percent := rand.Float64()*(mutator.toPercent-mutator.fromPercent) + mutator.fromPercent
	return round(percent * float64(genesLen))
}
func (mutator *InvertMutator) getInterval(genesLen, intervalLen int) (int, int) {
	if intervalLen > genesLen {
		panic(fmt.Sprintf("Interval bigger than chromosome. %d > %d", intervalLen, genesLen))
	}

	firstPoint := rand.Intn(genesLen - intervalLen + 1)
	return firstPoint, firstPoint + intervalLen
}
func (mutator *InvertMutator) mutate(genes GenesInterface, from, to int) {
	genesLen := genes.Len()

	if mutator.temp == nil || mutator.temp.Len() != genesLen {
		mutator.temp = mutator.chromosomeConstructor(genesLen).Genes()
	}

	mutator.temp.Copy(genes, from, from, to)

	for i := from; i < to; i++ {
		ind := to + from - i - 1
		genes.Set(i, mutator.temp.Get(ind))
	}
}
