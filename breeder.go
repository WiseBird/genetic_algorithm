package genetic_algorithm

type BreederInterface interface {
	ParentsCount() int
	Crossover(Chromosomes) (Chromosomes)
}