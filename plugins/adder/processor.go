package main

import (
	"github.com/Pungyeon/plugin/model"
)

var _ model.Processor = &processor{}

var Processor processor

type processor struct {
	value int
}

func (p *processor) Process(a []int) []int {
	output := make([]int, len(a))

	for i, x := range a {
		output[i] = x + p.value
	}
	return output
}

func (p *processor) WithOptions(options model.Options) {
	p.value = options.Value
}

func main() {}
