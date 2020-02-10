package model

type Options struct {
	Value int
}

type Processor interface {
	Process([]int) []int
	WithOptions(Options)
}
