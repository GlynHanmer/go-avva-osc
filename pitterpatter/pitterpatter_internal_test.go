package pitterpatter

import (
	"testing"
)

func Test_inRange(t *testing.T) {
	testSets := []struct{
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