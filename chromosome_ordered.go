package genetic_algorithm

import (
	"fmt"
	"strconv"
	"bytes"
)

type OrderedGenes []int
func (b OrderedGenes) Len() int { return len(b) }
/*func (b OrderedGenes) Copy(genes GenesInterface, from1, from2, to2 int) int {
	bgenes, ok := genes.(OrderedGenes)
	if !ok {
		panic("Unexpected genes. Expected OrderedGenes")
	}

	return copy(b[from1:], bgenes[from2:to2])
}*/

type OrderedChromosome struct {
	*ChromosomeBase
	genes OrderedGenes
}
func NewOrderedChromosome(genes OrderedGenes) *OrderedChromosome {
	chrom := new(OrderedChromosome)

	chrom.ChromosomeBase = NewChromosomeBase()
	chrom.genes = genes

	return chrom
}
func NewEmptyOrderedChromosome(genesLen int) ChromosomeInterface {
	genes := make(OrderedGenes, genesLen)
	for i := 0; i < len(genes); i++ {
		genes[i] = -1
	}
	return NewOrderedChromosome(genes)
}
func (chrom *OrderedChromosome) Genes() GenesInterface {
	return chrom.genes
}
func (chrom *OrderedChromosome) OrderedGenes() OrderedGenes {
	return chrom.genes
}
func (chrom *OrderedChromosome) String() string {
	var buffer bytes.Buffer

	for i, g := range chrom.genes {
		if i != 0 {
			buffer.WriteString(" ")
		}

		buffer.WriteString(strconv.Itoa(g))
	}
	return fmt.Sprintf("BC genes:[%v], cost: %f", buffer.String(), chrom.costVal)
}