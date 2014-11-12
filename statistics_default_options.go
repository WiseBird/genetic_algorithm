package genetic_algorithm

// Options for default statistics
type StatisticsDefaultOptions struct {
	trackMinCosts bool
	trackMinCostAge bool
	trackMeanCost bool
	trackMeanCosts bool
	trackWorstCost bool
	trackWorstCosts bool
	trackMinCostsVar bool
}
func NewStatisticsDefaultOptions() *StatisticsDefaultOptions {
	return new(StatisticsDefaultOptions)
}
func (options *StatisticsDefaultOptions) TrackMinCosts() {
	options.trackMinCosts = true
}
func (options *StatisticsDefaultOptions) TrackMinCostAge() {
	options.trackMinCostAge = true
}
func (options *StatisticsDefaultOptions) TrackMeanCost() {
	options.trackMeanCost = true
}
func (options *StatisticsDefaultOptions) TrackMeanCosts() {
	options.trackMeanCosts = true
}
func (options *StatisticsDefaultOptions) TrackWorstCost() {
	options.trackWorstCost = true
}
func (options *StatisticsDefaultOptions) TrackWorstCosts() {
	options.trackWorstCosts = true
}
func (options *StatisticsDefaultOptions) TrackMinCostsVar() {
	options.TrackMinCosts()
	options.trackMinCostsVar = true
}