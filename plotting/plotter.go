package plotting

import (
	. "github.com/WiseBird/genetic_algorithm"
	"math"

	pplot "code.google.com/p/plotinum/plot"
	"code.google.com/p/plotinum/plotter"
	"code.google.com/p/plotinum/plotutil"

	"io"
	"os"
	"path/filepath"
	"strings"

	"code.google.com/p/plotinum/vg"
	"code.google.com/p/plotinum/vg/vgeps"
	"code.google.com/p/plotinum/vg/vgimg"
	"code.google.com/p/plotinum/vg/vgpdf"
	"code.google.com/p/plotinum/vg/vgsvg"
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

type canvas interface {
	vg.Canvas
	Size() (w, h vg.Length)
	io.WriterTo
}

type Plotter struct {
	plots []*plot
}

func NewPlotter() *Plotter {
	plotter := new(Plotter)

	plotter.plots = make([]*plot, 0, 1)

	return plotter
}
func (plotter *Plotter) AddPlotWithData(data StatisticsDataInterface) *plot {
	plot := newPlot(plotter, nil, data)
	plotter.plots = append(plotter.plots, plot)

	return plot
}
func (plotter *Plotter) AddPlot(optimizer OptimizerInterface) *plot {
	plot := newPlot(plotter, optimizer, nil)
	plotter.plots = append(plotter.plots, plot)

	return plot
}
func (plotter *Plotter) Draw(widthInch, heightInch float64, fileName string) []StatisticsDataInterface {
	for _, p := range plotter.plots {
		if p.optimizer != nil {
			_, p.statisticsData = p.optimizer.Optimize()
		}

		for _, dataSet := range p.dataSets {
			err := plotutil.AddLinePoints(
				p.plot,
				dataSet.name,
				dataSet.values(p.statisticsData))

			if err != nil {
				panic(err)
			}
		}
	}

	w, h := vg.Inches(widthInch), vg.Inches(heightInch)
	c := plotter.createCanvas(fileName, len(plotter.plots), w, h)
	for i, p := range plotter.plots {
		plotter.draw(p.plot, i, c, w, h)
	}

	if err := plotter.saveFile(c, fileName); err != nil {
		panic(err)
	}

	data := make([]StatisticsDataInterface, len(plotter.plots))
	for i, p := range plotter.plots {
		data[i] = p.statisticsData
	}
	return data
}
func (plotter *Plotter) createCanvas(fileName string, plots int, w, h vg.Length) canvas {
	h *= vg.Length(plots)

	switch ext := strings.ToLower(filepath.Ext(fileName)); ext {

	case ".eps":
		return vgeps.NewTitle(w, h, fileName)

	case ".jpg", ".jpeg":
		return vgimg.JpegCanvas{Canvas: vgimg.New(w, h)}

	case ".pdf":
		return vgpdf.New(w, h)

	case ".png":
		return vgimg.PngCanvas{Canvas: vgimg.New(w, h)}

	case ".svg":
		return vgsvg.New(w, h)

	case ".tiff":
		return vgimg.TiffCanvas{Canvas: vgimg.New(w, h)}

	default:
		panic("Unsupported file extension: " + ext)
	}
}
func (plotter *Plotter) draw(plot *pplot.Plot, ind int, c canvas, w, h vg.Length) {
	_, canvasHeight := c.Size()
	da := pplot.DrawArea{
		Canvas: c,
		Rect: pplot.Rect{
			Min:  pplot.Point{0, canvasHeight - h*vg.Length(ind+1)},
			Size: pplot.Point{w, h},
		},
	}

	plot.Draw(da)
}
func (plotter *Plotter) saveFile(c canvas, fileName string) (err error) {
	f, err := os.Create(fileName)
	if err != nil {
		return err
	}
	if _, err = c.WriteTo(f); err != nil {
		return err
	}
	return f.Close()
}

type plot struct {
	plotter *Plotter

	optimizer      OptimizerInterface
	statisticsData StatisticsDataInterface

	plot     *pplot.Plot
	dataSets []*plotDataSet
}

func newPlot(plotter *Plotter, optimizer OptimizerInterface, statisticsData StatisticsDataInterface) *plot {
	p := new(plot)

	p.plotter = plotter
	p.optimizer = optimizer
	p.statisticsData = statisticsData

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
	return p.AddDataSet("Min", func(sa StatisticsDataInterface) plotter.XYs {
		sda, ok := sa.(StatisticsDataDefault)
		if !ok {
			panic("Expects StatisticsDefaultAggregator")
		}

		return CostsConverter(sda.MinCosts())
	})
}
func (p *plot) AddMeanCostDataSet() *plotDataSet {
	return p.AddDataSet("Mean", func(sa StatisticsDataInterface) plotter.XYs {
		sda, ok := sa.(StatisticsDataDefault)
		if !ok {
			panic("Expects StatisticsDefaultAggregator")
		}

		return CostsConverter(sda.MeanCosts())
	})
}
func (p *plot) Done() *Plotter {
	return p.plotter
}
func (p *plot) InnerPlot(title string) *pplot.Plot {
	return p.plot
}

type DataExtracter func(StatisticsDataInterface) plotter.XYs
type ValueConverter func(float64) float64
type plotDataSet struct {
	plot *plot

	name       string
	extracter  DataExtracter
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
func (dataSet *plotDataSet) values(statisticsData StatisticsDataInterface) plotter.XYs {
	xys := dataSet.extracter(statisticsData)
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
