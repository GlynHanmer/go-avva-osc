package scene

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCutType(t *testing.T) {
	for _, test := range []struct {
		input string
		err   bool
		ct    cutType
	}{
		{
			err: true,
		},
		{
			input: "bloopedydoop",
			err:   true,
		},
		{
			input: "auto",
			ct:    auto,
		},
		{
			input: "manual",
			ct:    manual,
		},
		{
			input: "end_of_scene",
			ct:    endOfScene,
		},
	} {
		t.Run(test.input, func(t *testing.T) {
			ct, err := NewCutType(test.input)
			if test.err {
				assert.Error(t, err)
				assert.Nil(t, ct)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.ct, *ct)
			}
		})
	}
}
