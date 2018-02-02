package meshbend

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGainActions(t *testing.T) {
	for _, test := range []struct {
		action string
		gain
		err bool
	}{
		{
			err: true,
		},
		{
			action: "boop",
			err:    true,
		},
		{
			action: "inc",
			gain:   gainInc,
		},
		{
			action: "dec",
			gain:   gainDec,
		},
		{
			action: "invert",
			gain:   gainInvert,
		},
	} {
		t.Run(test.action, func(t *testing.T) {
			gain, err := NewGain(test.action)
			if test.err {
				assert.Error(t, err)
				assert.Nil(t, gain)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.gain, *gain)
			}
		})
	}
}
