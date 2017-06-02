package pitterpatter

import (
	"fmt"
	"encoding/json"
	"log"
	"github.com/hypebeast/go-osc/osc"
)

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

type PitterPatter struct {
	Opacity *Opacity
}

func (pp PitterPatter) Osc() *osc.Message {
	json, err := json.Marshal(pp)
	if err != nil {
		log.Printf("Error creating PitterPatter json for osc message: %s", err.Error())
	}
	return osc.NewMessage("/pitterpatter", string(json))
}