package main

var Processor processor

type processor struct {
}

func (p *processor) Process(a []int) []int {
	output := make([]int, len(a))

	for i, x := range a {
		output[i] = x * 3
	}
	return output
}

func main() {}
