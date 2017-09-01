package meshbend

import (
	"fmt"
	"github.com/hypebeast/go-osc/osc"
	"errors"
)

type gain string
const name string = "meshbend"
const address string = "/" + name

func NewGain(action string) gain {
	return gain(action)
}

func (g *gain) Osc() *osc.Message {
	return osc.NewMessage(address, "syncgain", string(*g))
}

// String makes gain adhere to the Stringer interface
func (e gain) String() string {
	return fmt.Sprintf("%s %s", name, string(e))
}


type threshold float64

func NewThreshold(t float64) (threshold, error) {
	switch {
	case t < 0:
		return threshold(0), errors.New("Threshold cannot be negative.")
	case t > 1:
		return threshold(1), errors.New("Threshold cannot be greater than 1.")
	}
	return threshold(t), nil
}

func (t *threshold) Osc() *osc.Message {
	return osc.NewMessage(address, "threshold", float64(*t))
}

// String makes gain adhere to the Stringer interface
func (t threshold) String() string {
	return fmt.Sprintf("%s %f", name, t)
}
