package flexitModbus

import (
	"encoding/binary"
	"fmt"
	"time"

	"github.com/goburrow/modbus"
)

type registerExecutorFunc func(client modbus.Client) ([]byte, error)

const (
	serialDevice      = "/dev/ttyUSB0"
	numberOfRegisters = 49
)

func ReadAllHoldingRegisters() (registers map[string]FlexitRegister, err error) {
	fn := func(client modbus.Client) ([]byte, error) {
		return client.ReadHoldingRegisters(0, numberOfRegisters)
	}
	bytesFromAllRegisters, err := registerExecutor(fn)
	if err != nil {
		return nil, err
	}
	registers = GetAllHoldingRegisters()
	parseRegistersInto(bytesFromAllRegisters, registers)
	return
}

func ReadAllInputRegisters() (registers map[string]FlexitRegister, err error) {
	fn := func(client modbus.Client) ([]byte, error) {
		return client.ReadInputRegisters(0, numberOfRegisters)
	}
	bytesFromAllRegisters, err := registerExecutor(fn)
	if err != nil {
		return nil, err
	}
	registers = GetAllInputRegisters()
	parseRegistersInto(bytesFromAllRegisters, registers)
	return
}

func ReadHoldingRegister(registerName string) (holdingRegister *FlexitRegister, err error) {
	register, found := GetAllHoldingRegisters()[registerName]
	if !found {
		return nil, fmt.Errorf("%v is not a holding register", registerName)
	}
	fn := func(client modbus.Client) ([]byte, error) {
		return client.ReadHoldingRegisters(register.Address, register.GetNumberOfRegisters())
	}
	registerBytes, err := registerExecutor(fn)
	if err != nil {
		return nil, err
	}
	register.SetValueFromByteArray(registerBytes)
	return &register, nil
}

func GetHoldingRegister(registerName string) (holdingRegister *FlexitRegister, err error) {
	register, found := GetAllHoldingRegisters()[registerName]
	if !found {
		return nil, fmt.Errorf("%v is not a holding register", registerName)
	}
	return &register, nil
}

func ReadInputRegister(registerName string) (inputRegister *FlexitRegister, err error) {
	register, found := GetAllInputRegisters()[registerName]
	if !found {
		return nil, fmt.Errorf("%v is not a input register", registerName)
	}
	fn := func(client modbus.Client) ([]byte, error) {
		return client.ReadInputRegisters(register.Address, register.GetNumberOfRegisters())
	}
	registerBytes, err := registerExecutor(fn)
	if err != nil {
		return nil, err
	}
	register.SetValueFromByteArray(registerBytes)
	return &register, nil
}

func GetInputRegister(registerName string) (inputRegister *FlexitRegister, err error) {
	register, found := GetAllInputRegisters()[registerName]
	if !found {
		return nil, fmt.Errorf("%v is not a input register", registerName)
	}
	return &register, nil
}

func WriteHoldingRegister(register FlexitRegister) error {
	writeValue := uint16(register.Value)
	fn := func(client modbus.Client) ([]byte, error) {
		return client.WriteSingleRegister(register.Address, writeValue)
	}
	bytes, err := registerExecutor(fn)
	if err != nil {
		return err
	}
	readValue := binary.BigEndian.Uint16(bytes)
	if readValue != writeValue {
		return fmt.Errorf("Response value %v is not equal sendt value %v", readValue, writeValue)
	}
	return err
}

func registerExecutor(fn registerExecutorFunc) ([]byte, error) {
	handler := newUsbModbusClient()

	err := handler.Connect()
	if err != nil {

		return nil, fmt.Errorf("ERROR Failed to connect: %v", err)
	}
	defer handler.Close()

	client := modbus.NewClient(handler)
	return fn(client)
}

func parseRegistersInto(registers []byte, target map[string]FlexitRegister) {
	for key, register := range target {
		startIndex := register.Address * 2
		endIndex := startIndex + register.GetNumberOfBytes()
		registerBytes := registers[startIndex:endIndex]
		register.SetValueFromByteArray(registerBytes)
		target[key] = register
	}
	return
}

func newUsbModbusClient() *modbus.RTUClientHandler {
	handler := modbus.NewRTUClientHandler(serialDevice)
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "E" // "E"ven, "O"dd, "N"o parity
	handler.StopBits = 1
	handler.SlaveId = 1
	handler.Timeout = 1 * time.Second
	// handler.Logger = log.New(os.Stdout, "", log.LstdFlags)
	return handler
}
