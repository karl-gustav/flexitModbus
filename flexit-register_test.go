package flexitModbus

import (
	"bytes"
	"encoding/json"
	"log"
	"testing"
)

var (
	minusThreeHundredByteArray = []byte{254, 212}
)

func TestFlexitRegister(t *testing.T) {

	var register = TemperatureRegister{
		BaseRegister: BaseRegister{
			Address:     13,
			Name:        "ReturnWaterTemp",
			Format:      "INT 16",
			Description: "Return water temperature",
			Write:       false,
			Unit:        "1°C",
		},
		Default: 20,
		Min:     -45,
		Max:     50,
	}

	register.SetValueFromByteArray(minusThreeHundredByteArray)

	if register.Value != -30.0 {
		t.Errorf("setValueFromByteArray test value was incorrect, got: %d, want: %d.", register.Value, -3)
	}

	v := register.GetValueAsByteArray()
	expect := minusThreeHundredByteArray

	if !bytes.Equal(v, expect) {
		t.Errorf("getValueAsByteArray value was incorrect, got: %d, want: %d.", v, expect)
	}

	register.SetValueFromByteArray(minusThreeHundredByteArray)
	v = register.GetValueAsByteArray()
	expect = minusThreeHundredByteArray

	if !bytes.Equal(v, expect) {
		t.Errorf("get/set -ByteValue value was incorrect, got: %d, want: %d.", v, expect)
	}

	jsonString, _ := json.Marshal(&register)
	log.Println(string(jsonString))
	json.Unmarshal(jsonString, &register)
	log.Println("register.Value:", register.Value)

}
