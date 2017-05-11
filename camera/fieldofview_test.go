package camera

import (
	"testing"
	"encoding/json"
)

func Test_createFieldOfView(t *testing.T) {
	uinti := uint8(0)
	_ = FieldOfView(&uinti)
}

func Test_FieldOfViewJson(t *testing.T) {
	var emptyUint8 FieldOfView
	jsonFov, err := json.Marshal(emptyUint8)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(jsonFov))
	validInts := []uint8{0, 5, 200, 255}
	for _, validInt := range validInts {
		fov := FieldOfView(&validInt)
		jsonFov, err = json.Marshal(fov)
		if err != nil {
			t.Error(err)
		}
		t.Log(string(jsonFov))
	}
}
