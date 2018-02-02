package meshbend_test

import (
	"testing"

	"github.com/glynternet/go-avva-osc/meshbend"
	"github.com/hypebeast/go-osc/osc"
)

func TestGainActions(t *testing.T) {
	client := osc.NewClient("192.168.2.2", 9000)
	for _, action := range []string{"inc", "dec", "invert", "inc"} {
		_ = action
		gain := meshbend.NewGain(action)
		err := client.Send(gain.Osc())
		if err != nil {
			t.Error(err)
		}
	}
}
