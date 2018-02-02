package orbit

import (
	"fmt"

	"github.com/hypebeast/go-osc/osc"
)

type orbit string

const (
	inc    = orbit("inc")
	dec    = orbit("dec")
	invert = orbit("invert")
)

func NewOrbit(action string) (*orbit, error) {
	e := orbit(action)
	switch e {
	case inc, dec, invert:
	default:
		return nil, fmt.Errorf("action not supported: %s", action)
	}
	return &e, nil
}

func NewIncrement() orbit {
	return inc
}

func NewDecrement() orbit {
	return dec
}

func NewInvert() orbit {
	return invert
}

func (o orbit) Osc() *osc.Message {
	return osc.NewMessage("/orbit", "syncgain", string(o))
}

// String makes orbit adhere to the Stringer interface
func (o orbit) String() string {
	return fmt.Sprintf("%s %s", "orbit", string(o))
}
