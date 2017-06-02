package pitterpatter_test

import (
	"testing"
	"github.com/GlynOwenHanmer/go-avva-osc/pitterpatter"
)

func Test_PitterPatterOsc(t *testing.T) {
	testSets := []struct{
		pitterpatter.PitterPatter
	}{
		{},
		{pitterpatter.PitterPatter{}},
		{pitterpatter.PitterPatter{
			Opacity:&pitterpatter.Opacity{},
		}},
	}

	for _, testSet := range testSets {
		t.Log(testSet.Osc())
	}
}
