package tilingoffset

import "testing"

func TestName(t *testing.T) {
	to := NewIncrement()
	msg1 := to.Osc()
	t.Log(msg1)
	msg2 := *to.Osc()
	t.Log(msg2)
}