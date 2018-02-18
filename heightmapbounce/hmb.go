package heightmapbounce

import (
	"fmt"

	"github.com/hypebeast/go-osc/osc"
)

const (
	inc = HMB("inc")
	dec = HMB("dec")
)

type HMB string

func NewHMB(action string) (*HMB, error) {
	e := HMB(action)
	switch e {
	case inc, dec:
	default:
		return nil, fmt.Errorf("action not supported: %s", action)
	}
	return &e, nil
}

func NewIncrement() HMB {
	return HMB("inc")
}

func NewDecrement() HMB {
	return HMB("dec")
}

func (h HMB) Generate() *osc.Message {
	return osc.NewMessage("/heightmapbounce", "gain", string(h))
}

// String makes HMB adhere to the Stringer interface
func (h HMB) String() string {
	return fmt.Sprintf("%s %s", "hmb", string(h))
}
