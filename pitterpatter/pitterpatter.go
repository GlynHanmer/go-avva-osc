package pitterpatter

import "fmt"

const (
	minValue float64 = 0.0
	maxValue float64 = 1.0
)

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