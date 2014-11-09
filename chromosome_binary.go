package genetic_algorithm

import (
	"fmt"
)

type BinaryChromosome struct {
	*ChromosomeBase
	Genes []bool
}
func NewBinaryChromosome(genes []bool) *BinaryChromosome {
	chrom := new(BinaryChromosome)

	chrom.ChromosomeBase = NewChromosomeBase()
	chrom.Genes = genes

	return chrom
}
func (chrom *BinaryChromosome) Size() int {
	return len(chrom.Genes)
}
func (chrom *BinaryChromosome) String() string {
	genes := ""
	for _, b := range chrom.Genes {
		if b {
			genes += "1"
		} else {
			genes += "0"
		}
	}
	return fmt.Sprintf("BC genes:%v, cost: %f", genes, chrom.costVal)
}