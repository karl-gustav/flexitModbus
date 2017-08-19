package main

import (
	"encoding/binary"
	"fmt"
	"strings"
)

type FlexitRegister struct {
	Address     uint16 `json:"address"`
	Name        string `json:"name"`
	Format      string `json:"format"`
	Description string `json:"description"`
	Write       bool   `json:"write"`
	Unit        string `json:"unit"`
	Default     int64  `json:"default"`
	Min         int64  `json:"min"`
	Max         int64  `json:"max"`
	Value       int64  `json:"value"`
	ByteValue   string `json:"byteValue"`
}

// Returns how many 8 bits blocks the register consists of (remember that one modbus register consists of 16bits)
func (f *FlexitRegister) GetNumberOfBytes() uint16 {
	if strings.HasSuffix(f.Format, "32") {
		return 4
	}
	return 2
}

func (f *FlexitRegister) GetNumberOfRegisters() uint16 {
	return f.GetNumberOfBytes() / 2
}

func (f *FlexitRegister) SetValueFromByteArray(bytes []byte) {
	f.ByteValue = byteArrayToString(bytes)
	switch f.Format {
	case "UINT 16":
		f.Value = int64(binary.BigEndian.Uint16(bytes))
	case "UINT 32":
		f.Value = int64(binary.BigEndian.Uint32(bytes))
	case "INT 16":
		f.Value = int64(byteArrayToInt16(bytes))
	case "INT 32":
		f.Value = int64(byteArrayToInt32(bytes))
	}
}

func (f *FlexitRegister) GetValueAsByteArray() []byte {
	// Can use Uint for Int since the binary reprencentation is the same for int and uint
	var bs []byte
	switch f.Format {
	case "UINT 16", "INT 16":
		bs = make([]byte, 2)
		binary.BigEndian.PutUint16(bs, uint16(f.Value))
	case "UINT 32", "INT 32":
		bs = make([]byte, 4)
		binary.BigEndian.PutUint32(bs, uint32(f.Value))
	}
	return bs
}

func byteArrayToInt16(bytes []byte) int16 {
	data := int16(0)
	for _, b := range bytes {
		data = (data << 8) | int16(b)
	}
	return data
}

func byteArrayToInt32(bytes []byte) int32 {
	data := int32(0)
	for _, b := range bytes {
		data = (data << 8) | int32(b)
	}
	return data
}

func byteArrayToString(bytes []byte) string {
	var output []string
	for _, e := range bytes {
		output = append(output, fmt.Sprintf("%d", e))
	}
	return "[" + strings.Join(output, ", ") + "]"
}
