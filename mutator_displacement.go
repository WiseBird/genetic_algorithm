package genetic_algorithm

import (
	"math/rand"
)

// Mutator selects some part of the chromosome and places it in random position
type DisplacementMutator struct {
	*MutatorIntervalBase
}

// Probability is applied to each chromosome
func NewDisplacementMutator(probability float64, chromosomeConstructor EmptyChromosomeConstructor) *DisplacementMutator {
	mutator := new(DisplacementMutator)

	mutator.MutatorIntervalBase = NewMutatorIntervalBase(mutator, probability, chromosomeConstructor)

	return mutator
}

// Mutator selects random element from chromosome and places it in random position
// Probability is applied to each chromosome
func NewInsertionMutator(probability float64, chromosomeConstructor EmptyChromosomeConstructor) *MutatorIntervalBase {
	return NewDisplacementMutator(probability, chromosomeConstructor).ExactInterval(1, 1)
}

func (mutator *DisplacementMutator) MutateGenes(genes GenesInterface, from, to int) {
	insertPoint := mutator.chooseInsertPoint(genes.Len(), from, to)
	mutator.insert(genes, from, to, insertPoint)
}
func (mutator *DisplacementMutator) chooseInsertPoint(genesLen, from, to int) int {
	possiblePoints := genesLen - (to - from)
	point := rand.Intn(possiblePoints)
	if point >= from {
		point += to - from + 1
	}
	return point
}
func (mutator *DisplacementMutator) insert(genes GenesInterface, from, to, insertPoint int) {
	temp := mutator.getIntervalCopy(genes, from, to)

	if insertPoint < from {
		copyLen := from - insertPoint
		genes.Copy(genes, to-copyLen, insertPoint, from)
		genes.Copy(temp, insertPoint, from, to)
	} else {
		copyLen := insertPoint - to
		genes.Copy(genes, from, to, insertPoint)
		genes.Copy(temp, from+copyLen, from, to)
	}
}
