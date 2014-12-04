package genetic_algorithm

type CrossoverInterface interface {
	ParentsCount() int
	Crossover(Chromosomes) Chromosomes
}
