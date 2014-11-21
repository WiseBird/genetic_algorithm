package plotting

import (
	. "github.com/WiseBird/genetic_algorithm"
	"math"
    pplot "code.google.com/p/plotinum/plot"
    "code.google.com/p/plotinum/plotter"
    "code.google.com/p/plotinum/plotutil"
)

var (
	Log10 = func(f float64) float64 {
        if f > 0 {
	        return math.Log10(f)
	    } else if f < 0 {
	    	return -1 * math.Log10(math.Abs(f))
	    } else {
	    	return 0
	    }
	}
	CostsConverter = func(costs []float64) plotter.XYs {
	    pts := make(plotter.XYs, len(costs))
	    for i, cost := range costs {
	        pts[i].X = float64(i)
	        pts[i].Y = cost
	    }
	    return pts
	}
)

type Plotter struct {
	plots []*plot
}
func NewPlotter() *Plotter {
	plotter := new(Plotter)

	plotter.plots = make([]*plot, 0, 1)

	return plotter
}
func (plotter *Plotter) AddPlot(statisticsAggregator StatisticsAggregatorInterface) *plot {
	plot := newPlot(plotter, nil, statisticsAggregator, 1)
	plotter.plots = append(plotter.plots, plot)

	return plot
}
func (plotter *Plotter) AddPlotWithComputation(optimizer OptimizerInterface, statisticsAggregator StatisticsAggregatorInterface) *plot {
	plot := newPlot(plotter, optimizer, statisticsAggregator, 1)
	plotter.plots = append(plotter.plots, plot)

	return plot
}
func (plotter *Plotter) AddPlotWithComputations(optimizer OptimizerInterface, statisticsAggregator StatisticsAggregatorInterface, iterations int) *plot {
	plot := newPlot(plotter, optimizer, statisticsAggregator, iterations)
	plotter.plots = append(plotter.plots, plot)

	return plot
}
func (plotter *Plotter) Draw(widthInch, heightInch float64, fileName string) {
	for _, p := range plotter.plots {
		if p.optimizer != nil {
			for i := 0; i < p.iterations; i++ {
				_, statistics := p.optimizer.Optimize()
				p.statisticsAggregator.Aggregate(statistics)
			}
			p.statisticsAggregator.Compute()
		}

		for _, dataSet := range p.dataSets {
		    err := plotutil.AddLinePoints(
		    	p.plot,
		    	dataSet.name,
	    		dataSet.values(p.statisticsAggregator))

		    if err != nil {
	            panic(err)
		    }
		}
	}

	for _, p := range plotter.plots {
	    if err := p.plot.Save(widthInch, heightInch, fileName); err != nil {
            panic(err)
	    }
	}
}

type plot struct {
	plotter *Plotter

	optimizer OptimizerInterface
	statisticsAggregator StatisticsAggregatorInterface
	iterations int

	plot *pplot.Plot
	dataSets []*plotDataSet
}
func newPlot(plotter *Plotter, optimizer OptimizerInterface, statisticsAggregator StatisticsAggregatorInterface, iterations int) *plot {
	p := new(plot)

	p.plotter = plotter
	p.optimizer = optimizer
	p.statisticsAggregator = statisticsAggregator
	p.iterations = iterations

	p.dataSets = make([]*plotDataSet, 0, 1)

	var err error
	p.plot, err = pplot.New()
    if err != nil {
            panic(err)
    }

    p.plot.Title.Text = "Plot"
    p.plot.X.Label.Text = "Generations"
    p.plot.Y.Label.Text = "Cost"

	return p
}
func (p *plot) Title(title string) *plot {
	p.plot.Title.Text = title
	return p
}
func (p *plot) XLabel(label string) *plot {
	p.plot.X.Label.Text = label
	return p
}
func (p *plot) YLabel(label string) *plot {
	p.plot.Y.Label.Text = label
	return p
}
func (p *plot) AddDataSet(name string, extracter DataExtracter) *plotDataSet {
	dataSet := newPlotDataSet(p, name, extracter)
	p.dataSets = append(p.dataSets, dataSet)

	return dataSet
}
func (p *plot) AddMinCostDataSet() *plotDataSet {
	return p.AddDataSet("Min", func(sa StatisticsAggregatorInterface) plotter.XYs {
		sda, ok := sa.(*StatisticsDefaultAggregator)
		if !ok {
			panic("Expects StatisticsDefaultAggregator")
		}

		return CostsConverter(sda.MinCosts())
	});
}
func (p *plot) AddMeanCostDataSet() *plotDataSet {
	return p.AddDataSet("Mean", func(sa StatisticsAggregatorInterface) plotter.XYs {
		sda, ok := sa.(*StatisticsDefaultAggregator)
		if !ok {
			panic("Expects StatisticsDefaultAggregator")
		}

		return CostsConverter(sda.MeanCosts())
	});
}
func (p *plot) Done() *Plotter {
	return p.plotter
}
func (p *plot) InnerPlot(title string) *pplot.Plot {
	return p.plot
}

type DataExtracter func(StatisticsAggregatorInterface) plotter.XYs
type ValueConverter func(float64) float64
type plotDataSet struct {
	plot *plot

	name string
	extracter DataExtracter
	xConverter ValueConverter
	yConverter ValueConverter
}
func newPlotDataSet(plot *plot, name string, extracter DataExtracter) *plotDataSet {
	dataSet := new(plotDataSet)

	dataSet.plot = plot
	dataSet.name = name
	dataSet.extracter = extracter

	return dataSet
}
func (dataSet *plotDataSet) XConverter(converter ValueConverter) *plotDataSet {
	dataSet.xConverter = converter
	return dataSet
}
func (dataSet *plotDataSet) YConverter(converter ValueConverter) *plotDataSet {
	dataSet.yConverter = converter
	return dataSet
}
func (dataSet *plotDataSet) values(statisticsAggregator StatisticsAggregatorInterface) plotter.XYs {
	xys := dataSet.extracter(statisticsAggregator)
	if dataSet.xConverter != nil {
		for i := 0; i < len(xys); i++ {
			xys[i].X = dataSet.xConverter(xys[i].X)
		}
	}
	if dataSet.yConverter != nil {
		for i := 0; i < len(xys); i++ {
			xys[i].Y = dataSet.yConverter(xys[i].Y)
		}
	}
	return xys
}
func (dataSet *plotDataSet) Done() *plot {
	return dataSet.plot
}