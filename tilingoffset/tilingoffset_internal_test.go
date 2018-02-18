package tilingoffset

import "testing"

func TestName(t *testing.T) {
	to := NewIncrement()
	msg1 := to.Generate()
	t.Log(msg1)
	msg2 := *to.Generate()
	t.Log(msg2)
}
