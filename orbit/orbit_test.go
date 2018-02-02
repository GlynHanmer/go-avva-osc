package orbit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrbit(t *testing.T) {
	for _, test := range []struct {
		action string
		orbit
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
			orbit:  inc,
		},
		{
			action: "dec",
			orbit:  dec,
		},
		{
			action: "invert",
			orbit:  invert,
		},
	} {
		t.Run(test.action, func(t *testing.T) {
			orbit, err := NewOrbit(test.action)
			if test.err {
				assert.Error(t, err)
				assert.Nil(t, orbit)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.orbit, *orbit)
			}
		})
	}
}
