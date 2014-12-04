package genetic_algorithm

type InitializerInterface interface {
	Init(count, chromSize int) Chromosomes
}
