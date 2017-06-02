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

type lift float64

func (a *lift) equals(b *lift) bool {
	switch {
	case a == b:
	case a == nil != (b == nil),
		*a != *b:
		return false
	}
	return true
}

func (l *lift) String() string {
	if l != nil {
		return fmt.Sprintf("%f", *l)
	}
	return `<nil>`
}

func Lift(value float64) (*lift, error) {
	err := inRange(value)
	if err != nil {
		return nil, err
	}
	return liftPointer(value), err
}

func liftPointer(value float64) *lift {
	l := lift(value)
	return &l
}

type opacity struct {
	Gain *gain
	Lift *lift
}
