package millisperbeat

import (
	"github.com/hypebeast/go-osc/osc"
	"testing"
)

type TestSet struct {
	expected MillisPerBeat
}

var MillisPerBeatTestSets = []TestSet{
	TestSet{MillisPerBeat{}},
}

func Test_createMillisPerBeat(t *testing.T) {
	for key, ts := range MillisPerBeatTestSets {
		actual := NewMillisPerBeat()
		if actual != ts.expected {
			t.Fatalf("MillisPerBeatTestSets[%d]: MillisPerBeat returned not as expected.\n\tActual  : %v\n\tExpected: %v", key, actual, ts.expected)
		}
	}
}

type OscTestSet struct {
	expected *osc.Message
}

var MillisPerBeatOSCTestTests = []OscTestSet{
	OscTestSet{osc.NewMessage("/audio", "millisperbeat", float32(250))},
}

func Test_createCanvasOscMessages(t *testing.T) {
	for key, ts := range MillisPerBeatOSCTestTests {
		actual := NewMillisPerBeat().Osc()
		if !actual.Equals(ts.expected) {
			t.Fatalf("MillisPerBeatOSCTestTests[%d]: OSC Message generated not as expected:\n\t\n\tActual  : %v\n\tExpected: %v", key, actual, ts.expected)
		}
	}
}

func Test_StringerInterface(t *testing.T) {
	var testSets = []struct {
		expected string
	}{
		{"/audio millisperbeat"},
	}
	for _, ts := range testSets {
		c := NewMillisPerBeat()
		actual := c.String()
		if actual != ts.expected {
			t.Fatalf("Actual and expected String() output are not as expected\n\tActual  :%s\n\tExpected: %s", actual, ts.expected)
		}
	}
}
