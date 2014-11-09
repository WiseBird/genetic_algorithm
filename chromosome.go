package genetic_algorithm

import (
	"fmt"
	log "github.com/cihub/seelog"
)

type ChromosomeInterface interface {
	Size() int
	SetCost(float64)
	Cost() float64
	// Calcs fitness=(0,1] from normalized(0 - best case) cost.
	// Must throw error if cost is negative.
	Fitness() float64
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
func (c Chromosomes) String() string {
	s := ""
	for i := 0; i < len(c); i++ {
		if i != 0 {
			s += "\n"
		}
		s += fmt.Sprintf("%v", c[i])
	}
	return s
}