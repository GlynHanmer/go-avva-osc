package heightmapbounce

import (
	"fmt"
	"github.com/hypebeast/go-osc/osc"
)

type HMB string

func NewHMB(action string) HMB {
	return HMB(action)
}

func NewGainInc() HMB {
	return HMB("inc")
}

func NewGainDec() HMB {
	return HMB("dec")
}

func (h HMB) Osc() *osc.Message {
	return osc.NewMessage("/heightmapbounce", "gain", string(h))
}

// String makes HMB adhere to the Stringer interface
func (h HMB) String() string {
	return fmt.Sprintf("%s %s", "hmb", string(h))
}
