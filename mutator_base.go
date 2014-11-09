package genetic_algorithm

import (
	"math"
	"math/rand"
	log "github.com/cihub/seelog"
)

const (
	// Will roll for every element in chromosome, mutate it if success
	MutatorOneByOneType = 0
	// Will mutete exactly (Npop * Nel * P) elements
	MutatorExactCountType = 1
)

type MutatorBase struct {
	MutatorBaseVirtualMInterface

	Probability float64
	Elitism int
	Type int
}

type MutatorBaseVirtualMInterface interface {
	MutateCromosome(chrom ChromosomeInterface, ind int)
}

func NewMutator(virtual MutatorBaseVirtualMInterface, probability float64) *MutatorBase {
	mutator := new(MutatorBase)

	mutator.MutatorBaseVirtualMInterface = virtual
	mutator.Probability = probability
	mutator.Elitism = 1
	mutator.Type = MutatorOneByOneType

	return mutator
}
// Will roll for every element in chromosome, mutate it if success
func (mutator *MutatorBase) OneByOne() *MutatorBase {
	mutator.Type = MutatorOneByOneType
	return mutator
}
// Will mutete exactly (Npop * Nel * P) elements
func (mutator *MutatorBase) ExactCount() *MutatorBase {
	mutator.Type = MutatorExactCountType
	return mutator
}
// The best chromosome can't be mutated
func (mutator *MutatorBase) WithElitism(count int) *MutatorBase {
	if count < 0 {
		panic("Elitism can't be negative")
	}

	mutator.Elitism = count
	return mutator
}
// All chromosomes can be mutated
func (mutator *MutatorBase) WithoutElitism() *MutatorBase {
	mutator.Elitism = 0
	return mutator
}

func (mutator *MutatorBase) Mutate(population Chromosomes) {
	switch mutator.Type {
		case MutatorOneByOneType:
			mutator.mutateOneByOne(population)
		case MutatorExactCountType:
			mutator.mutateExactCount(population)
	}
}
func (mutator *MutatorBase) mutateOneByOne(population Chromosomes) {
	m := 0
	for ind, chrom := range population {
		if mutator.Elitism > ind {
			continue
		}

		for i := 0; i < chrom.Genes().Len(); i++ {
			if rand.Float64() > mutator.Probability {
				continue
			}

			log.Tracef("Mutate: %v, at %d\n", chrom, i)
			mutator.MutateCromosome(chrom, i)
			m++
		}
	}

	log.Debugf("Elems mutated: %d", m)
}
func (mutator *MutatorBase) mutateExactCount(population Chromosomes) {
	popLen := len(population)
	if popLen == 0 {
		return
	}

	genesLen := population[0].Genes().Len()

	chromsToMutate := popLen - mutator.Elitism
	elementsToMutate := int(math.Floor(mutator.Probability * float64(chromsToMutate * genesLen)))
	log.Debugf("ElemsToMutate: %d", elementsToMutate)

	for i := 0; i < elementsToMutate; i++ {
		chromInd := rand.Intn(popLen - mutator.Elitism) + mutator.Elitism;
		elemInd := rand.Intn(genesLen)

		log.Tracef("Mutate: %v, at %d\n", population[chromInd], elemInd)
		mutator.MutateCromosome(population[chromInd], elemInd)
	}
}