package heightmapbounce

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmission(t *testing.T) {
	for _, test := range []struct {
		action string
		err    bool
		HMB
	}{
		{
			err: true,
		},
		{
			action: "bloop",
			err:    true,
		},
		{
			action: "inc",
			HMB:    inc,
		},
		{
			action: "dec",
			HMB:    dec,
		},
	} {
		t.Run(test.action, func(t *testing.T) {
			e, err := NewHMB(test.action)
			if test.err {
				assert.Error(t, err)
				assert.Nil(t, e)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &test.HMB, e)
			}
		})
	}
}
