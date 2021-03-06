package genetic_algorithm

import (
	"bytes"
	"fmt"
	"strconv"
)

type OrderedGenes []int

func (g OrderedGenes) Len() int                   { return len(g) }
func (g OrderedGenes) Swap(i, j int)              { g[i], g[j] = g[j], g[i] }
func (g OrderedGenes) Get(i int) interface{}      { return g[i] }
func (g OrderedGenes) Set(i int, val interface{}) { g[i] = val.(int) }
func (g OrderedGenes) Copy(genes GenesInterface, from1, from2, to2 int) int {
	bgenes, ok := genes.(OrderedGenes)
	if !ok {
		panic("Unexpected genes. Expected OrderedGenes")
	}

	return copy(g[from1:], bgenes[from2:to2])
}
func (g OrderedGenes) Ind(val int) int {
	for i := 0; i < len(g); i++ {
		if g[i] == val {
			return i
		}
	}
	return -1
}

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
func (chrom *OrderedChromosome) Ind(val int) int {
	return chrom.genes.Ind(val)
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
