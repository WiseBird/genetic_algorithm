package genetic_algorithm

import (
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
	mutator.Elitism = count
	return mutator
}
// All chromosomes can be mutated
func (mutator *MutatorBase) WithoutElitism() *MutatorBase {
	mutator.Elitism = 0
	return mutator
}

func (mutator *MutatorBase) Mutate(population []ChromosomeInterface) {
	switch mutator.Type {
		case MutatorOneByOneType:
			mutator.mutateOneByOne(population)
		case MutatorExactCountType:
			mutator.mutateExactCount(population)
	}
}
func (mutator *MutatorBase) mutateOneByOne(population []ChromosomeInterface) {
	for ind, chrom := range population {
		if mutator.Elitism > ind {
			continue
		}

		for i := 0; i < chrom.Size(); i++ {
			if rand.Float64() > mutator.Probability {
				continue
			}

			log.Tracef("Mutate: %v, at %d\n", chrom, i)
			mutator.MutateCromosome(chrom, i)
		}
	}
}
func (mutator *MutatorBase) mutateExactCount(population []ChromosomeInterface) {
	panic("Doesn't implemented")
}