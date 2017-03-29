package orbit

import (
	"fmt"
	"github.com/hypebeast/go-osc/osc"
)

type Orbit string

func NewOrbit(action string) Orbit {
	return Orbit(action)
}

func NewIncrement() Orbit {
	return Orbit("inc")
}

func NewDecrement() Orbit {
	return Orbit("dec")
}

func (o Orbit) Osc() *osc.Message {
	return osc.NewMessage("/orbit", "syncgain", string(o))
}

// String makes Orbit adhere to the Stringer interface
func (o Orbit) String() string {
	return fmt.Sprintf("%s %s", "orbit", string(o))
}
