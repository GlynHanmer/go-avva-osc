package pitterpatter

import (
	"testing"
	"bytes"
	"fmt"
)

func Test_gainEquals(t *testing.T) {
		testSets := []struct{
			gainA *gain
			gainB *gain
			equal bool
		}{
			{
				gainA:nil,
				gainB:nil,
				equal:true,
			},
			{
				gainA:gainPointer(0.5),
				gainB:nil,
				equal:false,
			},
			{
				gainA:gainPointer(0.5),
				gainB:gainPointer(0.5),
				equal:true,
			},
			{
				gainA:gainPointer(0.0),
				gainB:gainPointer(1.0),
				equal:false,
			},
			{
				gainA:nil,
				gainB:gainPointer(0.5),
				equal:false,
			},
		}
	for _, testSet := range testSets {
		equal := testSet.gainA.equals(testSet.gainB)
		if testSet.equal != equal {
			var message bytes.Buffer
			fmt.Fprintf(&message, "Expecting: %t\nActual  : %t", testSet.equal, equal)
			fmt.Fprintf(&message, "\nA: %s", testSet.gainA)
			fmt.Fprintf(&message, "\nB: %s\n", testSet.gainB)
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