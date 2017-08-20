package flexitModbus

import (
	"bytes"
	"testing"
)

func TestFlexitRegister(t *testing.T) {

	var register = FlexitRegister{
		Address:     13,
		Name:        "ReturnWaterTemp",
		Format:      "INT 16",
		Description: "Return water temperature",
		Write:       false,
		Unit:        "0.1°C",
		Default:     20,
		Min:         -45,
		Max:         50,
	}

	register.SetValueFromByteArray([]byte{255, 253})

	if register.Value != -3 {
		t.Errorf("setValueFromByteArray test value was incorrect, got: %d, want: %d.", register.Value, -3)
	}

	v := register.GetValueAsByteArray()
	expect := []byte{255, 253}

	if !bytes.Equal(v, expect) {
		t.Errorf("getValueAsByteArray value was incorrect, got: %d, want: %d.", v, expect)
	}

	register.SetValueFromByteArray([]byte{255, 253})
	v = register.GetValueAsByteArray()
	expect = []byte{255, 253}

	if !bytes.Equal(v, expect) {
		t.Errorf("get/set -ByteValue value was incorrect, got: %d, want: %d.", v, expect)
	}

}
