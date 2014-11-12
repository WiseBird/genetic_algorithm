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
	generations := 100

	initializer := ga.BinaryRandomInitializerInstance
	weeder := ga.NewSimpleWeeder(weedRate)
	selector := ga.NewRouletteWheelCostWeightingSelector()
	breeder := ga.NewOnePointBreeder(ga.NewEmptyBinaryChromosome)
	mutator := ga.NewBinaryMutator(mutationProb)
	stopCriterion := ga.NewStopCriterionDefault().
		MaxGenerations(generations).
		MinCost(0).
		MaxMinCostAge(15)

	optimizer := ga.NewOptimizer(initializer, weeder, selector, breeder, mutator, cost, popSize, chromSize)
	_, statistics := optimizer.Optimize(stopCriterion)
	stats := statistics.(*ga.StatisticsDefault)

	log.Infof("Duration: %v", stats.Duration())
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
            "First", randomPoints(15),
            "Second", randomPoints(15),
            "Third", randomPoints(15))
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