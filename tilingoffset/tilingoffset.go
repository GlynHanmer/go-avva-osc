package tilingoffset

import (
"fmt"
"github.com/hypebeast/go-osc/osc"
)

type TilingOffset string

func NewTilingOffset(action string) TilingOffset {
	return TilingOffset(action)
}

func NewIncrement() TilingOffset {
	return TilingOffset("inc")
}

func NewDecrement() TilingOffset {
	return TilingOffset("dec")
}

func NewInvert() TilingOffset {
	return TilingOffset("invert")
}

func (o *TilingOffset) Osc() *osc.Message {
	return osc.NewMessage("/tilingoffset", "syncgain", string(*o))
}

// String makes Orbit adhere to the Stringer interface
func (o TilingOffset) String() string {
	return fmt.Sprintf("%s %s", "tilingoffset", string(o))
}

