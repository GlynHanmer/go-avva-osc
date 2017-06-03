package pitterpatter

import "fmt"

type normalisedFloat float64

const (
	minValue float64 = 0.0
	maxValue float64 = 1.0
)

func (a *normalisedFloat) equals(b *normalisedFloat) bool {
	switch {
	case a == b:
	case a == nil != (b == nil),
		*a != *b:
		return false
	}
	return true
}

func (g *normalisedFloat) String() string {
	if g != nil {
		return fmt.Sprintf("%f", *g)
	}
	return `<nil>`
}

func NormalisedFloat(value float64) (*normalisedFloat, error) {
	err := inRange(value)
	if err != nil {
		return nil, err
	}
	return normalisedFloatPointer(value), err
}


func normalisedFloatPointer(value float64) *normalisedFloat {
	l := normalisedFloat(value)
	return &l
}

type OutOfRangeError struct {
	value float64
	min float64
	max float64
}

func (e OutOfRangeError) Error() string {
	return fmt.Sprintf("Out of range. Value: %f, Min: %f, Max: %f", e.value, e.min, e.max)
}

func inRange(value float64) error {
	var err error
	if value < minValue || value > maxValue {
		err = OutOfRangeError{
			value:value,
			min:minValue,
			max:maxValue,
		}
	}
	return err
}