package pitterpatter

import (
	"bytes"
	"fmt"
	"testing"
)

func Test_inRange(t *testing.T) {
	testSets := []struct {
		float64
		expectError bool
	}{
		{-1.0, true},
		{0, false},
		{1.0, false},
		{20.0, true},
	}
	for _, testSet := range testSets {

		err := inRange(testSet.float64)
		if err == nil != testSet.expectError == false {
			t.Errorf("Expect error: %t\nReturned error: %s\nInput: %f", testSet.expectError, err, testSet.float64)
		}
	}
}

func Test_normalisedFloatEquals(t *testing.T) {
	testSets := []struct {
		A     *normalisedFloat
		B     *normalisedFloat
		equal bool
	}{
		{
			A:     nil,
			B:     nil,
			equal: true,
		},
		{
			A:     normalisedFloatPointer(0.5),
			B:     nil,
			equal: false,
		},
		{
			A:     normalisedFloatPointer(0.5),
			B:     normalisedFloatPointer(0.5),
			equal: true,
		},
		{
			A:     normalisedFloatPointer(0.0),
			B:     normalisedFloatPointer(1.0),
			equal: false,
		},
		{
			A:     nil,
			B:     normalisedFloatPointer(0.5),
			equal: false,
		},
	}
	for _, testSet := range testSets {
		equal := testSet.A.equals(testSet.B)
		if testSet.equal != equal {
			var message bytes.Buffer
			fmt.Fprintf(&message, "Expecting: %t\nActual  : %t", testSet.equal, equal)
			fmt.Fprintf(&message, "\nA: %s", testSet.A)
			fmt.Fprintf(&message, "\nB: %s\n", testSet.B)
			t.Error(message.String())
		}
	}
}

func Test_NormalisedFloat(t *testing.T) {
	testSets := []struct {
		float64
		gain  *normalisedFloat
		error bool
	}{
		{
			float64: -1,
			gain:    nil,
			error:   true,
		},
		{
			float64: 0,
			gain:    normalisedFloatPointer(0),
			error:   false,
		},
		{
			float64: 0.5,
			gain:    normalisedFloatPointer(0.5),
			error:   false,
		},
		{
			float64: 1,
			gain:    normalisedFloatPointer(1),
			error:   false,
		},
		{
			float64: 1.5,
			gain:    nil,
			error:   true,
		},
	}

	for _, testSet := range testSets {
		gain, err := NormalisedFloat(testSet.float64)
		if err == nil != testSet.error == false {
			t.Errorf("Unexpected error.\nExpect error: %t\nActual: %s", testSet.error, err)
		}
		if !gain.equals(testSet.gain) {
			t.Errorf("Unexpected value.\nExpected: %s\nActual  : %s", testSet.gain, gain)
		}
	}
}
