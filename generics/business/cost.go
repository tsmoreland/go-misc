package business

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer
}

func (s Solar) Cost() float64 {
	return s.Netto * 0.4
}

func (w Wind) Cost() float64 {
	return w.Netto * 0.4
}

func Cost[T Number](usage, netto T) T {
	cost := usage * netto
	return cost
}
