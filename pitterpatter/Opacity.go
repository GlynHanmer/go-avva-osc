package pitterpatter

import "fmt"

type gain float64

func (a *gain) equals(b *gain) bool {
	switch {
	case a == b:
	case a == nil != (b == nil),
		*a != *b:
		return false
	}
	return true
}

func (g *gain) String() string {
	if g != nil {
		return fmt.Sprintf("%f", *g)
	}
	return `<nil>`
}
//type lift *float64

type opacity struct {
	Gain *gain
	//Lift *lift
}

func Gain(value float64) (*gain, error) {
	err := inRange(value)
	if err != nil {
		return nil, err
	}
	return gainPointer(value), err
}

func gainPointer(value float64) *gain {
	g := gain(value)
	return &g
}

//
//func Lift(value float64) (gain, error) {
//	return gain(&value), inRange(value)
//}
