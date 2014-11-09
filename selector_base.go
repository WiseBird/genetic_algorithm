package genetic_algorithm

import (
	log "github.com/cihub/seelog"
	"fmt"
)

// Base class for selectors.
type SelectorBase struct {
	SelectorBaseVirtualMInterface

	populationLen int
	selectManyUnique bool
}

// Interface for SelectorBase's subclasses to implement
type SelectorBaseVirtualMInterface interface {
	Prepare(Chromosomes)
	Select() ChromosomeInterface
	//SelectExcept(map[int]bool) ChromosomeInterface
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

	selector.populationLen = len(population)

	selector.SelectorBaseVirtualMInterface.Prepare(population)
}

func (selector *SelectorBase) SelectMany(count int) Chromosomes {
	log.Tracef("SelectMany c=%d\n", count)

	if count < 0 {
		panic("Count must be greater than 0")
	}

	if selector.populationLen < count && selector.selectManyUnique {
		panic(fmt.Sprintf("Cant select %d unique chroms from %d chroms", count, selector.populationLen))
	}

	chroms := make(Chromosomes, count)
	//selected := make(map[int]bool, count)

	for i := 0; i < count; i++ {
		/*var chrom ChromosomeInterface
		if selector.selectManyUnique {
			chrom = selector.SelectorBaseVirtualMInterface.SelectExcept(selected)
			selected[]
		} else {
			chrom = selector.SelectorBaseVirtualMInterface.Select()
		}*/
		
		chrom := selector.SelectorBaseVirtualMInterface.Select()
		if selector.selectManyUnique {
			for ;; {
				unique := true
				for j := 0; j < i; j++ {
					if chroms[j] == chrom {

						log.Tracef("Collision try another")

						unique = false
						break
					}
				}
				if unique {
					break
				}

				chrom = selector.SelectorBaseVirtualMInterface.Select()
			}
		}

		log.Debugf("Parent[%d] - %v", i, chrom)
		chroms[i] = chrom
	}

	return chroms
}