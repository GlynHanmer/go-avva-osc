package pitterpatter_test

import (
	"testing"

	"github.com/glynternet/go-avva-osc/pitterpatter"
)

func Test_PitterPatterOsc(t *testing.T) {
	testSets := []struct {
		pitterpatter.PitterPatter
	}{
		{},
		{pitterpatter.PitterPatter{}},
		{pitterpatter.PitterPatter{
			Opacity: &pitterpatter.Opacity{},
		}},
	}

	for _, testSet := range testSets {
		t.Log(testSet.Generate())
	}
}
