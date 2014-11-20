package genetic_algorithm

import (
	log "github.com/cihub/seelog"
	"fmt"
)

// Base class for selectors.
type SelectorBase struct {
	SelectorBaseVirtualMInterface

	population Chromosomes
	selectManyUnique bool
}

// SelectorBase's virtual methods
type SelectorBaseVirtualMInterface interface {
	SelectInd() int
}

// Constructor for SelectorBase
func NewSelectorBase(virtual SelectorBaseVirtualMInterface) *SelectorBase {
	selector := new(SelectorBase)

	selector.SelectorBaseVirtualMInterface = virtual
	selector.selectManyUnique = true

	return selector
}

// Sets whether or not SelectMany will return only unique chromosomes
func (selector *SelectorBase) SelectManyAreUnique(value bool) *SelectorBase {
	selector.selectManyUnique = value
	return selector
}

func (selector *SelectorBase) Prepare(population Chromosomes) {
	log.Tracef("Prepare Population=%d\n", len(population))

	selector.population = population
}

func (selector *SelectorBase) Select() ChromosomeInterface {
	return selector.population[selector.SelectorBaseVirtualMInterface.SelectInd()]
}
func (selector *SelectorBase) SelectMany(count int) Chromosomes {
	log.Tracef("SelectMany c=%d\n", count)

	if count < 0 {
		panic("Count must be greater than 0")
	}

	if len(selector.population) < count && selector.selectManyUnique {
		panic(fmt.Sprintf("Cant select %d unique chroms from %d chroms", count, len(selector.population)))
	}

	chroms := make(Chromosomes, count)
	selected := make(map[int]bool, count)

	for i := 0; i < count; i++ {
		ind := selector.SelectorBaseVirtualMInterface.SelectInd()
		if selector.selectManyUnique && selected[ind] {
			j := 1
			for ;; {
				if !selected[ind-j] && ind-j >= 0 {
					ind = ind-j
					break
				}
				if !selected[ind-j] && ind+j < len(selector.population) {
					ind = ind+j
					break
				}
				j++
			}
		}

		selected[ind] = true
		chrom := selector.population[ind]

		log.Debugf("Parent[%d] - %v", i, chrom)
		chroms[i] = chrom
	}

	return chroms
}