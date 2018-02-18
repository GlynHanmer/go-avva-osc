package meshbend

import (
	"fmt"

	"github.com/hypebeast/go-osc/osc"
)

const (
	name    = "meshbend"
	address = "/" + name

	gainInc    = gain("inc")
	gainDec    = gain("dec")
	gainInvert = gain("invert")
)

type gain string

func NewGain(action string) (*gain, error) {
	e := gain(action)
	switch e {
	case gainInc, gainDec, gainInvert:
	default:
		return nil, fmt.Errorf("action not supported: %s", action)
	}
	return &e, nil
}

func (g *gain) Generate() *osc.Message {
	return osc.NewMessage(address, "syncgain", string(*g))
}

// String makes gain adhere to the Stringer interface
func (e gain) String() string {
	return fmt.Sprintf("%s %s", name, string(e))
}

type threshold float64

func NewThreshold(t float64) (*threshold, error) {
	switch {
	case t < 0:
		return newThreshold(0), fmt.Errorf("threshold cannot be negative, received %f", t)
	case t > 1:
		return newThreshold(1), fmt.Errorf("threshold cannot be greater than 1, received %f", t)
	}
	return newThreshold(t), nil
}

func newThreshold(t float64) *threshold {
	return (*threshold)(&t)
}

func (t *threshold) Generate() *osc.Message {
	return osc.NewMessage(address, "threshold", float64(*t))
}

// String makes gain adhere to the Stringer interface
func (t threshold) String() string {
	return fmt.Sprintf("%s %f", name, t)
}
