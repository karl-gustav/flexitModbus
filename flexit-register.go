package flexitModbus

import (
	"encoding/binary"
	"fmt"
	"strings"
)

type ReableRegister interface {
	GetAddress() uint16
	GetName() string
	GetFormat() string
	GetDescription() string
	GetUnit() string
	GetDefault() string
	GetValue() string
	GetMax() string
	GetMin() string
	GetNumberOfBytes() uint16
	GetNumberOfRegisters() uint16
	SetValueFromByteArray([]byte)
}

type WritableRegister interface {
	ReableRegister
	GetValueAsByteArray() []byte
	GetValueAsUInt16() uint16
	IsValueTooHigh() bool
	IsValueTooLow() bool
	SetValue(interface{}) error
}

type BaseRegister struct {
	Address     uint16 `json:"address"`
	Name        string `json:"name"`
	Format      string `json:"format"`
	Description string `json:"description"`
	Write       bool   `json:"write"`
	Unit        string `json:"unit"`
	ByteValue   string `json:"-"`
}

type Int16Register struct {
	BaseRegister
	Default int16 `json:"default"`
	Min     int16 `json:"min"`
	Max     int16 `json:"max"`
	Value   int16 `json:"value"`
}

type UInt16Register struct {
	BaseRegister
	Default uint16 `json:"default"`
	Min     uint16 `json:"min"`
	Max     uint16 `json:"max"`
	Value   uint16 `json:"value"`
}

type UInt32Register struct {
	BaseRegister
	Default uint32 `json:"default"`
	Min     uint32 `json:"min"`
	Max     uint32 `json:"max"`
	Value   uint32 `json:"value"`
}

type TemperatureRegister struct {
	BaseRegister
	Default float32 `json:"default"`
	Min     float32 `json:"min"`
	Max     float32 `json:"max"`
	Value   float32 `json:"value"`
}

func (f BaseRegister) GetAddress() uint16 {
	return f.Address
}

func (f BaseRegister) GetName() string {
	return f.Name
}

func (f BaseRegister) GetFormat() string {
	return f.Format
}

func (f BaseRegister) GetDescription() string {
	return f.Description
}

func (f BaseRegister) GetUnit() string {
	return f.Unit
}

// Returns how many 8 bits blocks the register consists of (remember that one modbus register consists of 16bits)
func (f BaseRegister) GetNumberOfBytes() uint16 {
	if strings.HasSuffix(f.Format, "32") {
		return 4
	}
	return 2
}

func (f BaseRegister) GetNumberOfRegisters() uint16 {
	return f.GetNumberOfBytes() / 2
}

func (f *Int16Register) SetValueFromByteArray(bytes []byte) {
	f.ByteValue = byteArrayToString(bytes)
	f.Value = byteArrayToInt16(bytes)
}

func (f *UInt16Register) SetValueFromByteArray(bytes []byte) {
	f.ByteValue = byteArrayToString(bytes)
	f.Value = binary.BigEndian.Uint16(bytes)
}

func (f *UInt32Register) SetValueFromByteArray(bytes []byte) {
	f.ByteValue = byteArrayToString(bytes)
	f.Value = binary.BigEndian.Uint32(bytes)
}

func (f *TemperatureRegister) SetValueFromByteArray(bytes []byte) {
	f.ByteValue = byteArrayToString(bytes)
	movedDecimalPoint := float32(byteArrayToInt16(bytes)) / float32(10)
	f.Value = movedDecimalPoint
}

func (f *Int16Register) SetValue(value interface{}) error {
	switch spesificValue := value.(type) {
	case int16:
		f.Value = spesificValue
	default:
		return fmt.Errorf("Invalid type used to set value: %T", spesificValue)
	}
	return nil
}

func (f *UInt16Register) SetValue(value interface{}) error {
	switch spesificValue := value.(type) {
	case uint16:
		f.Value = spesificValue
	default:
		return fmt.Errorf("Invalid type used to set value: %T", spesificValue)
	}
	return nil
}

func (f *UInt32Register) SetValue(value interface{}) error {
	switch spesificValue := value.(type) {
	case uint32:
		f.Value = spesificValue
	default:
		return fmt.Errorf("Invalid type used to set value: %T", spesificValue)
	}
	return nil
}

func (f *TemperatureRegister) SetValue(value interface{}) error {
	switch spesificValue := value.(type) {
	case float32:
		f.Value = spesificValue
	default:
		return fmt.Errorf("Invalid type used to set value: %T", spesificValue)
	}
	return nil
}

func (f Int16Register) GetValueAsByteArray() []byte {
	// Can use uint for int since the binary reprencentation is the same for int and uint
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, uint16(f.Value))
	return bs
}

func (f UInt16Register) GetValueAsByteArray() []byte {
	bs := make([]byte, 2)
	binary.BigEndian.PutUint16(bs, uint16(f.Value))
	return bs
}

func (f UInt32Register) GetValueAsByteArray() []byte {
	bs := make([]byte, 4)
	binary.BigEndian.PutUint32(bs, uint32(f.Value))
	return bs
}

func (f TemperatureRegister) GetValueAsByteArray() []byte {
	bs := make([]byte, 2)
	movedDecimalPoint := f.Value * float32(10)
	binary.BigEndian.PutUint16(bs, uint16(movedDecimalPoint))
	return bs
}

func (f Int16Register) IsValueTooHigh() bool {
	return f.Value > f.Max
}

func (f UInt16Register) IsValueTooHigh() bool {
	return f.Value > f.Max
}

func (f UInt32Register) IsValueTooHigh() bool {
	return f.Value > f.Max
}

func (f TemperatureRegister) IsValueTooHigh() bool {
	return f.Value > f.Max
}

func (f Int16Register) IsValueTooLow() bool {
	return f.Value < f.Min
}

func (f UInt16Register) IsValueTooLow() bool {
	return f.Value < f.Min
}

func (f UInt32Register) IsValueTooLow() bool {
	return f.Value < f.Min
}

func (f TemperatureRegister) IsValueTooLow() bool {
	return f.Value < f.Min
}

func (f Int16Register) GetValueAsUInt16() uint16 {
	// Can safely convert int16 --> uint16 because their binary reprecentations are the same
	return uint16(f.Value)
}

func (f UInt16Register) GetValueAsUInt16() uint16 {
	return uint16(f.Value)
}

func (f UInt32Register) GetValueAsUInt16() uint16 {
	panic("Can't safely convert uint32 to uint16, this only works for holding registers which are 16 bits")
}

func (f TemperatureRegister) GetValueAsUInt16() uint16 {
	return uint16(f.Value)
}

func (f Int16Register) GetValue() string {
	return fmt.Sprint(f.Value)
}

func (f UInt16Register) GetValue() string {
	return fmt.Sprint(f.Value)
}

func (f UInt32Register) GetValue() string {
	return fmt.Sprint(f.Value)
}

func (f TemperatureRegister) GetValue() string {
	return fmt.Sprint(f.Value)
}

func (f Int16Register) GetDefault() string {
	return fmt.Sprint(f.Default)
}

func (f UInt16Register) GetDefault() string {
	return fmt.Sprint(f.Default)
}

func (f UInt32Register) GetDefault() string {
	return fmt.Sprint(f.Default)
}

func (f TemperatureRegister) GetDefault() string {
	return fmt.Sprint(f.Default)
}

func (f Int16Register) GetMax() string {
	return fmt.Sprint(f.Max)
}

func (f UInt16Register) GetMax() string {
	return fmt.Sprint(f.Max)
}

func (f UInt32Register) GetMax() string {
	return fmt.Sprint(f.Max)
}

func (f TemperatureRegister) GetMax() string {
	return fmt.Sprint(f.Max)
}

func (f Int16Register) GetMin() string {
	return fmt.Sprint(f.Min)
}

func (f UInt16Register) GetMin() string {
	return fmt.Sprint(f.Min)
}

func (f UInt32Register) GetMin() string {
	return fmt.Sprint(f.Min)
}

func (f TemperatureRegister) GetMin() string {
	return fmt.Sprint(f.Min)
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
