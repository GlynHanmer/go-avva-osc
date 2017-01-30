package canvas

import (
	"github.com/hypebeast/go-osc/osc"
	"testing"
)

type TestSet struct {
	active   bool
	expected Canvas
}

var CanvasTestSets = []TestSet{
	TestSet{true, Canvas(true)},
	TestSet{false, Canvas(false)},
}

func Test_createCanvas(t *testing.T) {
	for key, ts := range CanvasTestSets {
		actual := NewCanvas(ts.active)
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
		actual := NewCanvas(ts.active).Osc()
		if !actual.Equals(ts.expected) {
			t.Fatalf("CanvasOSCTestTests[%d]: OSC Message generated not as expected:\n\tactive: %b\n\tActual  : %v\n\tExpected: %v", key, ts.active, actual, ts.expected)
		}
	}
}

func Test_StringerInterface(t *testing.T) {
	var testSets = []struct {
		state    bool
		expected string
	}{
		{true, "canvas on"},
		{false, "canvas off"},
	}
	for _, ts := range testSets {
		c := NewCanvas(ts.state)
		actual := c.String()
		if actual != ts.expected {
			t.Fatalf("Actual and expected String() output are not as expected\n\tActual  :%s\n\tExpected: %s\n\tState: %+v", actual, ts.expected, ts.state)
		}
	}
}
