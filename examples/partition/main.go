package main

import (
	ga "github.com/WiseBird/genetic_algorithm"
	"math"
	log "github.com/cihub/seelog"

    "code.google.com/p/plotinum/plot"
    "code.google.com/p/plotinum/plotter"
    "code.google.com/p/plotinum/plotutil"
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
		WithStatisticsOptions(ga.NewStatisticsDefaultOptions().
			TrackMinCosts().
			TrackMeanCosts()).
		WithPopSize(popSize).
		WithChromSize(chromSize)

	_, statistics := optimizer.Optimize()
	stats := statistics.(*ga.StatisticsDefault)

	log.Infof("Duration: %v", stats.Duration())
	log.Infof("MinCosts: %v", stats.MinCosts())
	log.Infof("MeanCosts: %v", stats.MeanCosts())

	drawPlot(stats)
}

func drawPlot(stats *ga.StatisticsDefault) {
    p, err := plot.New()
    if err != nil {
            panic(err)
    }

    p.Title.Text = "Plotutil example"
    p.X.Label.Text = "Gens"
    p.Y.Label.Text = "Cost"

    err = plotutil.AddLinePoints(p,
            "Mean", convertCostsToXYs(stats.MinCosts()),
            "Min", convertCostsToXYs(stats.MeanCosts()))
    if err != nil {
            panic(err)
    }

    //p.Y.Max = math.Max(1e4, p.Y.Min * 10)

    if err := p.Save(4, 4, "points.png"); err != nil {
            panic(err)
    }
}

func convertCostsToXYs(costs []float64) plotter.XYs {
    pts := make(plotter.XYs, len(costs))
    for i, cost := range costs {
        pts[i].X = float64(i)
        if cost != 0 {
	        pts[i].Y = math.Log10(cost)
	    } else {
	    	pts[i].Y = 0
	    }
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