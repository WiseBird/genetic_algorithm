package genetic_algorithm

import (
	"math/rand"
	"time"
)

type CostFunction func(ChromosomeInterface) float64

func init() {
	rand.Seed(time.Now().UnixNano())
}

type OptimizerInterface interface {
	Optimize() (ChromosomeInterface, StatisticsInterface)
}