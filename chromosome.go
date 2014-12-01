package genetic_algorithm

import (
	"fmt"
	log "github.com/cihub/seelog"
	"bytes"
)

type EmptyChromosomeConstructor func(genesLen int) ChromosomeInterface

type ChromosomeInterface interface {
	Genes() GenesInterface
	SetCost(float64)
	Cost() float64
}

type Chromosomes []ChromosomeInterface
func (c Chromosomes) Len() int           { return len(c) }
func (c Chromosomes) Less(i, j int) bool { return c[i].Cost() < c[j].Cost() }
func (c Chromosomes) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }
func (c Chromosomes) SetCost(cost CostFunction) {
	log.Tracef("Setting cost for %d chroms", len(c))
	
	for i := 0; i < len(c); i++ {
		var chrom = c[i]
		chrom.SetCost(cost(chrom))
	}
}
func (c Chromosomes) MeanCost() float64 {
	if len(c) == 0 {
		return 0
	}

	var sum float64
	for _, chrom := range c {
		sum += chrom.Cost()
	}
	return sum / float64(len(c))
}
func (c Chromosomes) String() string {
	var buffer bytes.Buffer

	for i := 0; i < len(c); i++ {
		if i != 0 {
			buffer.WriteString("\n")
		}
		buffer.WriteString(fmt.Sprintf("%v", c[i]))
	}
	return buffer.String()
}

type GenesInterface interface {
	Len() int
}
type CopyableGenesInterface interface {
	Copy(genes GenesInterface, from1, from2, to2 int) int
}