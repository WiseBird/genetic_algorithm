package genetic_algorithm

import (
	"fmt"
	log "github.com/cihub/seelog"
	"math"
	"math/rand"
)

const (
	// Will roll for every element in chromosome, mutate it if success
	MutatorOneByOneType = 0
	// Will mutete exactly (Npop * Nel * P) elements
	MutatorExactCountType = 1
)

// Base class for mutators that mutate separate genes
type MutatorGeneBase struct {
	MutatorGeneBaseVirtualMInterface

	probability float64
	elitism     int
	kind        int
}

// MutatorGeneBase's virtual methods
type MutatorGeneBaseVirtualMInterface interface {
	MutateCromosome(chrom ChromosomeInterface, ind int)
}

func NewGeneBaseMutator(virtual MutatorGeneBaseVirtualMInterface, probability float64) *MutatorGeneBase {
	if probability > 1 || probability < 0 {
		panic(fmt.Sprintf("Incorrect probability %v", probability))
	}

	mutator := new(MutatorGeneBase)

	mutator.MutatorGeneBaseVirtualMInterface = virtual
	mutator.probability = probability
	mutator.elitism = 1
	mutator.kind = MutatorOneByOneType

	return mutator
}

// Will roll for every element in chromosome, mutate it if success
func (mutator *MutatorGeneBase) OneByOne() *MutatorGeneBase {
	mutator.kind = MutatorOneByOneType
	return mutator
}

// Will mutete exactly (Npop * Nel * P) elements
func (mutator *MutatorGeneBase) ExactCount() *MutatorGeneBase {
	mutator.kind = MutatorExactCountType
	return mutator
}

// The best chromosome[s] can't be mutated
func (mutator *MutatorGeneBase) WithElitism(count int) *MutatorGeneBase {
	if count < 0 {
		panic("Elitism can't be negative")
	}

	mutator.elitism = count
	return mutator
}

// All chromosomes can be mutated
func (mutator *MutatorGeneBase) WithoutElitism() *MutatorGeneBase {
	mutator.elitism = 0
	return mutator
}

func (mutator *MutatorGeneBase) Mutate(population Chromosomes) {
	switch mutator.kind {
	case MutatorOneByOneType:
		mutator.mutateOneByOne(population)
	case MutatorExactCountType:
		mutator.mutateExactCount(population)
	}
}
func (mutator *MutatorGeneBase) mutateOneByOne(population Chromosomes) {
	m := 0
	for ind, chrom := range population {
		if mutator.elitism > ind {
			continue
		}

		for i := 0; i < chrom.Genes().Len(); i++ {
			if rand.Float64() > mutator.probability {
				continue
			}

			log.Tracef("Mutate: %v, at %d\n", chrom, i)
			mutator.MutateCromosome(chrom, i)
			m++
		}
	}

	log.Debugf("Elems mutated: %d", m)
}
func (mutator *MutatorGeneBase) mutateExactCount(population Chromosomes) {
	popLen := len(population)
	if popLen == 0 {
		return
	}

	genesLen := population[0].Genes().Len()

	chromsToMutate := popLen - mutator.elitism
	elementsToMutate := int(math.Floor(mutator.probability * float64(chromsToMutate*genesLen)))
	log.Debugf("ElemsToMutate: %d", elementsToMutate)

	for i := 0; i < elementsToMutate; i++ {
		chromInd := rand.Intn(popLen-mutator.elitism) + mutator.elitism
		elemInd := rand.Intn(genesLen)

		log.Tracef("Mutate: %v, at %d\n", population[chromInd], elemInd)
		mutator.MutateCromosome(population[chromInd], elemInd)
	}
}
