package genetic_algorithm

import (
	"bytes"
	"fmt"
)

type BinaryGenes []bool

func (g BinaryGenes) Len() int                   { return len(g) }
func (g BinaryGenes) Swap(i, j int)              { g[i], g[j] = g[j], g[i] }
func (g BinaryGenes) Get(i int) interface{}      { return g[i] }
func (g BinaryGenes) Set(i int, val interface{}) { g[i] = val.(bool) }
func (g BinaryGenes) Copy(genes GenesInterface, from1, from2, to2 int) int {
	bgenes, ok := genes.(BinaryGenes)
	if !ok {
		panic("Unexpected genes. Expected BinaryGenes")
	}

	return copy(g[from1:], bgenes[from2:to2])
}

type BinaryChromosome struct {
	*ChromosomeBase
	genes BinaryGenes
}

func NewBinaryChromosome(genes BinaryGenes) *BinaryChromosome {
	chrom := new(BinaryChromosome)

	chrom.ChromosomeBase = NewChromosomeBase()
	chrom.genes = genes

	return chrom
}
func NewEmptyBinaryChromosome(genesLen int) ChromosomeInterface {
	return NewBinaryChromosome(make(BinaryGenes, genesLen))
}
func (chrom *BinaryChromosome) Genes() GenesInterface {
	return chrom.genes
}
func (chrom *BinaryChromosome) BinaryGenes() BinaryGenes {
	return chrom.genes
}
func (chrom *BinaryChromosome) String() string {
	var buffer bytes.Buffer

	for _, b := range chrom.genes {
		if b {
			buffer.WriteString("1")
		} else {
			buffer.WriteString("0")
		}
	}
	return fmt.Sprintf("BC genes:[%v], cost: %f", buffer.String(), chrom.costVal)
}
