package camera

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewClearFlags(t *testing.T) {
	for _, test := range []struct{
		input string
		err bool
		ClearFlags
	}{
		{
			err: true,
		},
		{
			input: "oaihsd",
			err: true,
		},
		{
			input:      "skybox",
			ClearFlags: skybox,
		},
		{
			input:      "solid",
			ClearFlags: solid,
		},
		{
			input:      "depth",
			ClearFlags: depthOnly,
		},
	}{
		t.Run(test.input, func(t *testing.T){
			cf, err := NewClearFlag(test.input)
			if test.err {
				assert.Error(t, err)
				assert.Nil(t, cf)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.ClearFlags, *cf)
			}
		})
	}
}
