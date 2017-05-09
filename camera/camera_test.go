package camera

import (
	"testing"
	"encoding/json"
)

func Test_createCamera(t *testing.T) {
	cf := SKYBOX
	t.Log(cf)
}

func Test_ClearFlagJson(t *testing.T) {
	json, err := json.Marshal(SKYBOX)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(json))
}

func Test_CameraJson(t *testing.T){
	json, err := json.Marshal(SKYBOX)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(json))
}

//todo test that this can then be parsed back into a camera vis unmarshalling json??
func Test_CameraOsc(t *testing.T) {
	cam := Camera{ClearFlags: SKYBOX}
	msg := cam.Osc()
	t.Log(msg)
}