package main

import (
	ga "github.com/WiseBird/genetic_algorithm"
	"math"
	log "github.com/cihub/seelog"

    "code.google.com/p/plotinum/plot"
    "code.google.com/p/plotinum/plotter"
    "code.google.com/p/plotinum/plotutil"
    "math/rand"
)

var (
	list = []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20}
	finalSum = 178
	finalProduct = 3120
)

func cost(c ga.ChromosomeInterface) float64 {
	bc := c.(*ga.BinaryChromosome)
	sum := 0
	prod := 1

	genes := bc.Genes().(ga.BinaryGenes)

	for i := 0; i < len(genes); i++ {
		if genes[i] {
			prod *= list[i]
		} else {
			sum += list[i]
		}
	}

	sumDiff := float64(sum - finalSum)
	prodDiff := float64(prod - finalProduct)

	return math.Abs(sumDiff) + math.Abs(prodDiff * prodDiff)
}

func main() {
	defer log.Flush()
	setupLogger()

	popSize := 32
	chromSize := len(list)
	weedRate := 50.0
	mutationProb := 0.2
	generations := 200

	optimizer := ga.NewEmptyIncrementalOptimizer().
		WithInitializer(ga.BinaryRandomInitializerInstance).
		WithWeeder(ga.NewSimpleWeeder(weedRate)).
		WithSelector(ga.NewRouletteWheelCostWeightingSelector()).
		WithBreeder(ga.NewOnePointBreeder(ga.NewEmptyBinaryChromosome)).
		WithMutator(ga.NewBinaryMutator(mutationProb)).
		WithCostFunction(cost).
		WithStopCriterion(ga.NewStopCriterionDefault().
			MaxGenerations(generations).
			MinCost(0).
			MaxMinCostAge(15)).
		WithPopSize(popSize).
		WithChromSize(chromSize)

	_, statistics := optimizer.Optimize()
	stats := statistics.(*ga.StatisticsDefault)

	log.Infof("Duration: %v", stats.Duration())

	//drawPlot()
}

func drawPlot() {
    rand.Seed(int64(0))

    p, err := plot.New()
    if err != nil {
            panic(err)
    }

    p.Title.Text = "Plotutil example"
    p.X.Label.Text = "X"
    p.Y.Label.Text = "Y"

    err = plotutil.AddLinePoints(p,
            "First", randomPoints(150),
            "Second", randomPoints(150),
            "Third", randomPoints(150))
    if err != nil {
            panic(err)
    }

    // Save the plot to a PNG file.
    if err := p.Save(4, 4, "points.png"); err != nil {
            panic(err)
    }
}

// randomPoints returns some random x, y points.
func randomPoints(n int) plotter.XYs {
        pts := make(plotter.XYs, n)
        for i := range pts {
                if i == 0 {
                        pts[i].X = rand.Float64()
                } else {
                        pts[i].X = pts[i-1].X + rand.Float64()
                }
                pts[i].Y = pts[i].X + 10*rand.Float64()
        }
        return pts
}

func setupLogger() {
	logger, err := log.LoggerFromConfigAsFile("seelog.xml")

	if err != nil {
	    panic(err)
	}

	log.ReplaceLogger(logger)
}