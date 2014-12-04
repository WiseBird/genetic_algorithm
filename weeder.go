package genetic_algorithm

import (
	log "github.com/cihub/seelog"
	"math"
)

// Weeder weeds part of population.
type WeederInterface interface {
	Weed([]ChromosomeInterface) []ChromosomeInterface
}

// SimpleWeeder weeds part of population proportional to rate.
type SimpleWeeder struct {
	rate float64
}

// Creates new SimpleWeeder.
// Rate must be in [0,100)
func NewSimpleWeeder(rate float64) *SimpleWeeder {
	if rate < 0 || rate >= 100 {
		panic("Rate must be in [0,100)")
	}

	weeder := new(SimpleWeeder)

	weeder.rate = rate

	return weeder
}
func (weeder *SimpleWeeder) Weed(pop []ChromosomeInterface) []ChromosomeInterface {
	popLen := len(pop)

	log.Tracef("Weed Rate=%f Population=%d\n", weeder.rate, popLen)

	toWeed := int(math.Floor(float64(popLen) / 100.0 * weeder.rate))
	return pop[:popLen-toWeed]
}
