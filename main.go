package flexitModbus

import (
	"encoding/binary"
	"fmt"
	"sync"

	"github.com/goburrow/modbus"
)

type registerExecutorFunc func(client modbus.Client) ([]byte, error)

const (
	numberOfRegisters = 49
)

var clientHandlerFactory func() *modbus.RTUClientHandler

var lock sync.Mutex

func Setup(rtuClientHandlerFactory func() *modbus.RTUClientHandler) {
	clientHandlerFactory = rtuClientHandlerFactory
}

func ReadAllHoldingRegisters() (registers map[string]WritableRegister, err error) {
	fn := func(client modbus.Client) ([]byte, error) {
		return client.ReadHoldingRegisters(0, numberOfRegisters)
	}
	bytesFromAllRegisters, err := registerExecutor(fn)
	if err != nil {
		return nil, err
	}
	registers = GetAllHoldingRegisters()
	for _, register := range registers {
		parseRegistersInto(bytesFromAllRegisters, register)
	}
	return
}

func ReadAllInputRegisters() (registers map[string]ReableRegister, err error) {
	fn := func(client modbus.Client) ([]byte, error) {
		return client.ReadInputRegisters(0, numberOfRegisters)
	}
	bytesFromAllRegisters, err := registerExecutor(fn)
	if err != nil {
		return nil, err
	}
	registers = GetAllInputRegisters()
	for _, register := range registers {
		parseRegistersInto(bytesFromAllRegisters, register)
	}
	return
}

func ReadHoldingRegister(registerName string) (holdingRegister WritableRegister, err error) {
	register, found := GetAllHoldingRegisters()[registerName]
	if !found {
		return nil, fmt.Errorf("%v is not a holding register", registerName)
	}
	fn := func(client modbus.Client) ([]byte, error) {
		return client.ReadHoldingRegisters(register.GetAddress(), register.GetNumberOfRegisters())
	}
	registerBytes, err := registerExecutor(fn)
	if err != nil {
		return nil, err
	}
	register.SetValueFromByteArray(registerBytes)
	return register, nil
}

func GetHoldingRegister(registerName string) (holdingRegister WritableRegister, err error) {
	register, found := GetAllHoldingRegisters()[registerName]
	if !found {
		return nil, fmt.Errorf("%v is not a holding register", registerName)
	}
	return register, nil
}

func ReadInputRegister(registerName string) (inputRegister ReableRegister, err error) {
	register, found := GetAllInputRegisters()[registerName]
	if !found {
		return nil, fmt.Errorf("%v is not a input register", registerName)
	}
	fn := func(client modbus.Client) ([]byte, error) {
		return client.ReadInputRegisters(register.GetAddress(), register.GetNumberOfRegisters())
	}
	registerBytes, err := registerExecutor(fn)
	if err != nil {
		return nil, err
	}
	register.SetValueFromByteArray(registerBytes)
	return register, nil
}

func GetInputRegister(registerName string) (inputRegister ReableRegister, err error) {
	register, found := GetAllInputRegisters()[registerName]
	if !found {
		return nil, fmt.Errorf("%v is not a input register", registerName)
	}
	return register, nil
}

func WriteHoldingRegister(register WritableRegister) error {
	writeValue := register.GetValueAsUInt16()
	fn := func(client modbus.Client) ([]byte, error) {
		return client.WriteSingleRegister(register.GetAddress(), writeValue)
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

func registerExecutor(fn registerExecutorFunc) (result []byte, err error) {
	if clientHandlerFactory == nil {
		panic("rtuClientHandlerFactory was not set, did you forget to run the Setup() function?")
	}
	clientHandler := clientHandlerFactory()

	err = clientHandler.Connect()
	if err != nil {
		return nil, fmt.Errorf("ERROR Failed to connect: %v", err)
	}
	defer clientHandler.Close()

	client := modbus.NewClient(clientHandler)
	lock.Lock()
	result, err = fn(client)
	lock.Unlock()
	return
}

func parseRegistersInto(registers []byte, register ReableRegister) {
	startIndex := register.GetAddress() * 2
	endIndex := startIndex + register.GetNumberOfBytes()
	registerBytes := registers[startIndex:endIndex]
	register.SetValueFromByteArray(registerBytes)
	return
}
