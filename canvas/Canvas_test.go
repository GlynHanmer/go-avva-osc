package canvas

import (
	"github.com/hypebeast/go-osc/osc"
	"testing"
)

type TestSet struct {
	active   bool
	expected canvas
}

var CanvasTestSets = []TestSet{
	TestSet{true, canvas(true)},
	TestSet{false, canvas(false)},
}

func Test_createCanvas(t *testing.T) {
	for key, ts := range CanvasTestSets {
		actual := Canvas(ts.active)
		if actual != ts.expected {
			t.Fatalf("CanvasTestSets[%d}: Canvas returned not as expected.\n\tactive: %s\n\tActual  : %v\n\tExpected: %v", key, ts.active, actual, ts.expected)
		}
	}
}

type OscTestSet struct {
	active   bool
	expected *osc.Message
}

var CanvasOSCTestTests = []OscTestSet{
	OscTestSet{true, osc.NewMessage("/canvas", true)},
	OscTestSet{false, osc.NewMessage("/canvas", false)},
}

func Test_createCanvasOscMessages(t *testing.T) {
	for key, ts := range CanvasOSCTestTests {
		actual := Canvas(ts.active).Osc()
		if !actual.Equals(ts.expected) {
			t.Fatalf("CanvasOSCTestTests[%d]: OSC Message generated not as expected:\n\tactive: %b\n\tActual  : %v\n\tExpected: %v", key, ts.active, actual, ts.expected)
		}
	}
}
