package genetic_algorithm

// Base class for chromosomes.
type ChromosomeBase struct {
	costVal    float64
	fitnessVal float64
}
func NewChromosomeBase() *ChromosomeBase {
	return &ChromosomeBase{}
}
func (chrom *ChromosomeBase) SetCost(cost float64) {
	chrom.costVal = cost
	chrom.fitnessVal = 0
}
func (chrom *ChromosomeBase) Cost() float64 {
	return chrom.costVal
}