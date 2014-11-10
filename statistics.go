package genetic_algorithm

type StatisticsConstructor func() StatisticsInterface

type StatisticsInterface interface {
	Start()
	OnInteration(Chromosomes)
	End()

	Iterations() int

	MinCost() float64
	// Min cost for each iteration
	// Len would be `Iterations() + 1` because of initial value
	MinCosts() []float64
	// Number of iterations during which the value remains unchanged 
	MinCostAge() int

	MeanCost() float64
	MeanCosts() []float64

	WorstCost() float64
	WorstCosts() []float64

	TrackMinCost()
	TrackMinCosts()
	TrackMinAge()
	TrackMeanCost()
	TrackMeanCosts()
	TrackWorstCost()
	TrackWorstCosts()
}