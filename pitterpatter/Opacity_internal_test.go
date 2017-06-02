package pitterpatter

import (
	"testing"
	"bytes"
	"fmt"
)

func Test_gainEquals(t *testing.T) {
		testSets := []struct{
			A     *gain
			B     *gain
			equal bool
		}{
			{
				A:     nil,
				B:     nil,
				equal: true,
			},
			{
				A:     gainPointer(0.5),
				B:     nil,
				equal: false,
			},
			{
				A:     gainPointer(0.5),
				B:     gainPointer(0.5),
				equal: true,
			},
			{
				A:     gainPointer(0.0),
				B:     gainPointer(1.0),
				equal: false,
			},
			{
				A:     nil,
				B:     gainPointer(0.5),
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

func Test_Gain(t *testing.T) {
	testSets := []struct {
		float64
		gain  *gain
		error bool
	}{
		{
			float64: -1,
			gain:    nil,
			error: true,
		},
		{
			float64: 0,
			gain:    gainPointer(0),
			error:   false,
		},
		{
			float64: 0.5,
			gain: gainPointer(0.5),
			error: false,
		},
		{
			float64: 1,
			gain: gainPointer(1),
			error: false,
		},
		{
			float64: 1.5,
			gain: nil,
			error: true,
		},
	}

	for _, testSet := range testSets {
		gain, err := Gain(testSet.float64)
		if err == nil != testSet.error == false {
			t.Errorf("Unexpected error.\nExpect error: %t\nActual: %s", testSet.error, err)
		}
		if !gain.equals(testSet.gain) {
			t.Errorf("Unexpected value.\nExpected: %s\nActual  : %s", testSet.gain, gain)
		}
	}
}

func Test_liftEquals(t *testing.T) {
	testSets := []struct{
		A     *lift
		B     *lift
		equal bool
	}{
		{
			A:     nil,
			B:     nil,
			equal: true,
		},
		{
			A:     liftPointer(0.5),
			B:     nil,
			equal: false,
		},
		{
			A:     liftPointer(0.5),
			B:     liftPointer(0.5),
			equal: true,
		},
		{
			A:     liftPointer(0.0),
			B:     liftPointer(1.0),
			equal: false,
		},
		{
			A:     nil,
			B:     liftPointer(0.5),
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

func Test_Lift(t *testing.T) {
	testSets := []struct {
		float64
		lift  *lift
		error bool
	}{
		{
			float64: -1,
			lift:    nil,
			error:   true,
		},
		{
			float64: 0,
			lift:    liftPointer(0),
			error:   false,
		},
		{
			float64: 0.5,
			lift:    liftPointer(0.5),
			error:   false,
		},
		{
			float64: 1,
			lift:    liftPointer(1),
			error:   false,
		},
		{
			float64: 1.5,
			lift:    nil,
			error:   true,
		},
	}

	for _, testSet := range testSets {
		lift, err := Lift(testSet.float64)
		if err == nil != testSet.error == false {
			t.Errorf("Unexpected error.\nExpect error: %t\nActual: %s", testSet.error, err)
		}
		if !lift.equals(testSet.lift) {
			t.Errorf("Unexpected value.\nExpected: %s\nActual  : %s", testSet.lift, lift)
		}
	}
}