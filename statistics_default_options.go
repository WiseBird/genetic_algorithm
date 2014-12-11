package genetic_algorithm

// Options for default statistics
type StatisticsDefaultOptions struct {
	trackMinCosts    bool
	trackGensWoImprv bool
	trackMeanCost    bool
	trackMeanCosts   bool
	trackWorstCost   bool
	trackWorstCosts  bool
	trackMinCostsVar bool
	trackDurations   bool
}

func NewStatisticsDefaultOptions() *StatisticsDefaultOptions {
	return new(StatisticsDefaultOptions)
}
func (options *StatisticsDefaultOptions) TrackMinCosts() *StatisticsDefaultOptions {
	options.trackMinCosts = true
	return options
}
func (options *StatisticsDefaultOptions) TrackGenerationsWithoutImprovements() *StatisticsDefaultOptions {
	options.trackGensWoImprv = true
	return options
}
func (options *StatisticsDefaultOptions) TrackMeanCost() *StatisticsDefaultOptions {
	options.trackMeanCost = true
	return options
}
func (options *StatisticsDefaultOptions) TrackMeanCosts() *StatisticsDefaultOptions {
	options.trackMeanCosts = true
	return options
}
func (options *StatisticsDefaultOptions) TrackWorstCost() *StatisticsDefaultOptions {
	options.trackWorstCost = true
	return options
}
func (options *StatisticsDefaultOptions) TrackWorstCosts() *StatisticsDefaultOptions {
	options.trackWorstCosts = true
	return options
}
func (options *StatisticsDefaultOptions) TrackMinCostsVar() *StatisticsDefaultOptions {
	options.TrackMinCosts()
	options.trackMinCostsVar = true
	return options
}
func (options *StatisticsDefaultOptions) TrackDurations() *StatisticsDefaultOptions {
	options.trackDurations = true
	return options
}

func (options *StatisticsDefaultOptions) Ensure(other StatisticsOptionsInterface) {
	opt, ok := other.(*StatisticsDefaultOptions)
	if !ok {
		panic("Expect *StatisticsDefaultOptions")
	}

	if options.trackMinCosts {
		opt.TrackMinCosts()
	}
	if options.trackGensWoImprv {
		opt.TrackGenerationsWithoutImprovements()
	}
	if options.trackMeanCost {
		opt.TrackMeanCost()
	}
	if options.trackMeanCosts {
		opt.TrackMeanCosts()
	}
	if options.trackWorstCost {
		opt.TrackWorstCost()
	}
	if options.trackWorstCosts {
		opt.TrackWorstCosts()
	}
	if options.trackMinCostsVar {
		opt.TrackMinCostsVar()
	}
	if options.trackDurations {
		opt.TrackDurations()
	}
}
func (options *StatisticsDefaultOptions) Copy() *StatisticsDefaultOptions {
	return &StatisticsDefaultOptions{
		options.trackMinCosts,
		options.trackGensWoImprv,
		options.trackMeanCost,
		options.trackMeanCosts,
		options.trackWorstCost,
		options.trackWorstCosts,
		options.trackMinCostsVar,
		options.trackDurations,
	}
}
