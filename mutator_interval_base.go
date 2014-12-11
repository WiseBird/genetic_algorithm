package genetic_algorithm

import (
	"fmt"
	"math/rand"
)

const (
	mutatorInvertExactLen      = 0
	mutatorInvertPercentageLen = 1
)

// Base class for mutators that mutate some interval of genes
type MutatorIntervalBase struct {
	MutatorIntervalBaseVirtualMInterface

	probability           float64
	chromosomeConstructor EmptyChromosomeConstructor
	kind                  int
	fromExact             int
	toExact               int
	fromPercent           float64
	toPercent             float64

	temp GenesInterface
}

// MutatorIntervalBase's virtual methods
type MutatorIntervalBaseVirtualMInterface interface {
	MutateGenes(genes GenesInterface, from, to int)
}

// Probability is applied to each chromosome
// By default interval will be equal one third of the chromosome
func NewMutatorIntervalBase(virtual MutatorIntervalBaseVirtualMInterface, probability float64, chromosomeConstructor EmptyChromosomeConstructor) *MutatorIntervalBase {
	if probability > 1 || probability < 0 {
		panic(fmt.Sprintf("Incorrect probability %v", probability))
	}

	mutator := new(MutatorIntervalBase)

	mutator.MutatorIntervalBaseVirtualMInterface = virtual
	mutator.probability = probability
	mutator.chromosomeConstructor = chromosomeConstructor
	mutator.kind = mutatorInvertPercentageLen
	mutator.fromPercent = 0.33
	mutator.toPercent = 0.33

	return mutator
}

// Sets interval len to be randomly choosen from [from:to]
func (mutator *MutatorIntervalBase) ExactInterval(from, to int) *MutatorIntervalBase {
	if from > to || from < 0 || to < 1 {
		panic(fmt.Sprintf("Incorrect interval [%d:%d]", from, to))
	}

	mutator.kind = mutatorInvertExactLen
	mutator.fromExact = from
	mutator.toExact = to

	return mutator
}

// Sets interval len to be randomly choosen from [from*chromLen : to*chromLen]
func (mutator *MutatorIntervalBase) PercentageInterval(from, to float64) *MutatorIntervalBase {
	if from > to || from < 0 || to < 0 || from > 1 || to > 1 {
		panic(fmt.Sprintf("Incorrect percentage interval [%v:%v]", from, to))
	}

	mutator.kind = mutatorInvertPercentageLen
	mutator.fromPercent = from
	mutator.toPercent = to

	return mutator
}

func (mutator *MutatorIntervalBase) Mutate(population Chromosomes) {
	for _, chrom := range population {
		if mutator.probability < rand.Float64() {
			continue
		}

		genesLen := chrom.Genes().Len()
		intervalLen := mutator.getIntervalLen(genesLen)
		if intervalLen == 0 || intervalLen == genesLen {
			continue
		}

		from, to := mutator.getInterval(genesLen, intervalLen)

		mutator.MutateGenes(chrom.Genes(), from, to)
	}
}
func (mutator *MutatorIntervalBase) getIntervalLen(genesLen int) int {
	if mutator.kind == mutatorInvertExactLen {
		return rand.Intn(mutator.toExact-mutator.fromExact+1) + mutator.fromExact
	}

	percent := rand.Float64()*(mutator.toPercent-mutator.fromPercent) + mutator.fromPercent
	return round(percent * float64(genesLen))
}
func (mutator *MutatorIntervalBase) getInterval(genesLen, intervalLen int) (int, int) {
	if intervalLen > genesLen {
		panic(fmt.Sprintf("Interval bigger than chromosome. %d > %d", intervalLen, genesLen))
	}

	firstPoint := rand.Intn(genesLen - intervalLen + 1)
	return firstPoint, firstPoint + intervalLen
}
func (mutator *MutatorIntervalBase) getIntervalCopy(genes GenesInterface, from, to int) GenesInterface {
	genesLen := genes.Len()

	if mutator.temp == nil || mutator.temp.Len() != genesLen {
		mutator.temp = mutator.chromosomeConstructor(genesLen).Genes()
	}

	mutator.temp.Copy(genes, from, from, to)

	return mutator.temp
}
