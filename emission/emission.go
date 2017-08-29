package emission

import (
	"fmt"
	"github.com/hypebeast/go-osc/osc"
)

type Emission string

func NewEmission(action string) Emission {
	return Emission(action)
}

func NewIncrement() Emission {
	return Emission("inc")
}

func NewDecrement() Emission {
	return Emission("dec")
}

func (e *Emission) Osc() *osc.Message {
	return osc.NewMessage("/emission", "syncgain", string(*e))
}

// String makes Emission adhere to the Stringer interface
func (e Emission) String() string {
	return fmt.Sprintf("%s %s", "emission", string(e))
}

