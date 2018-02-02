package emission

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmission(t *testing.T) {
	for _, test := range []struct {
		action string
		err    bool
		Emission
	}{
		{
			err: true,
		},
		{
			action: "bloop",
			err:    true,
		},
		{
			action:   "inc",
			Emission: inc,
		},
		{
			action:   "dec",
			Emission: dec,
		},
	} {
		t.Run(test.action, func(t *testing.T) {
			e, err := NewEmission(test.action)
			if test.err {
				assert.Error(t, err)
				assert.Nil(t, e)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, &test.Emission, e)
			}
		})
	}
}
