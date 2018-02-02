package emission

import (
	"fmt"

	"github.com/hypebeast/go-osc/osc"
)

const (
	inc = Emission("inc")
	dec = Emission("dec")
)

type Emission string

func NewEmission(action string) (*Emission, error) {
	e := Emission(action)
	switch e {
	case inc, dec:
	default:
		return nil, fmt.Errorf("action not supported: %s", action)
	}
	return &e, nil
}

func NewIncrement() Emission {
	return inc
}

func NewDecrement() Emission {
	return dec
}

func (e *Emission) Osc() *osc.Message {
	return osc.NewMessage("/emission", "syncgain", string(*e))
}

// String makes Emission adhere to the Stringer interface
func (e Emission) String() string {
	return fmt.Sprintf("%s %s", "emission", string(e))
}
