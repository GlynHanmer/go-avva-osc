package canvas

import (
	"testing"

	"github.com/hypebeast/go-osc/osc"
	"github.com/magiconair/properties/assert"
)

func Test_createCanvas(t *testing.T) {
	for _, test := range []struct {
		active   bool
		expected Canvas
	}{
		{true, Canvas(true)},
		{false, Canvas(false)},
	} {
		assert.Equal(t, test.expected, NewCanvas(test.active))
	}
}

func Test_createCanvasOscMessages(t *testing.T) {
	for _, test := range []struct {
		active   bool
		expected *osc.Message
	}{
		{true, osc.NewMessage("/canvas", int32(1))},
		{false, osc.NewMessage("/canvas", int32(0))},
	} {
		actual := NewCanvas(test.active).Osc()
		assert.Equal(t, test.expected, actual)
	}
}

func Test_String(t *testing.T) {
	var testSets = []struct {
		state    bool
		expected string
	}{
		{true, "canvas on"},
		{false, "canvas off"},
	}
	for _, ts := range testSets {
		actual := NewCanvas(ts.state).String()
		assert.Equal(t, ts.expected, actual)
	}
}
