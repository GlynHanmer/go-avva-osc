package pitterpatter

import (
	"testing"
	"encoding/json"
	"bytes"
	"fmt"
)

func TestGenerateJsons(t *testing.T) {
	present := 0.5
	floatValues := []*float64{ nil, &present }

	opacities := generateOpacities(floatValues)
	fills := generateFills(floatValues)
	frequencies := generateFrequencies(floatValues)

	var pitterPatters []PitterPatter

	for _, opacity := range opacities {
		for _, fill := range fills {
			for _, freq := range frequencies {
				pitterPatters = append(pitterPatters,
					PitterPatter{Opacity: opacity, Fill: fill, Frequency: freq},
				)
			}
		}
	}

	for _, pp := range pitterPatters {
		jsonBytes, err := json.Marshal(pp)
		if err != nil {
			t.Fatal("Unable to marshal PitterPatter into json: Error: %s", err.Error())
		}
		var escapedBytes bytes.Buffer
		escapedBytes.WriteByte('"')
		for _, b := range jsonBytes {
			if b == '"' {escapedBytes.WriteByte(92)}
			escapedBytes.WriteByte(b)
		}
		escapedBytes.WriteString(`",`)
		t.Log(escapedBytes.String())
		fmt.Println(escapedBytes.String())
	}
}

func generateOpacities(values []*float64) []*Opacity {
	var opacities []*Opacity
	opacities = append(opacities, nil)
	for _, gain := range values {
		for _, lift := range values {
			opacity := Opacity{Gain: gain, Lift: lift}
			opacities = append(opacities, &opacity)
		}
	}
	return opacities
}

func generateFills(values []*float64) []*Fill {
	var fills []*Fill
	fills = append(fills, nil)
	for _, min := range values {
		for _, max := range values {
			fill := Fill{Min: min, Max: max}
			fills = append(fills, &fill)
		}
	}
	return fills
}

func generateFrequencies(values []*float64) []*Frequency {
	var freqs []*Frequency
	freqs = append(freqs, nil)
	for _, min := range values {
		for _, max := range values {
			freq := Frequency{Min: min, Max: max}
			freqs = append(freqs, &freq)
		}
	}
	return freqs
}