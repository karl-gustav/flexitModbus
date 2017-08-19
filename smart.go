package main

import (
	"fmt"
	"time"
)

func main() {
	inputRegisters, _ := ReadAllInputRegisters()
	fmt.Println("Input registers:")
	for _, r := range inputRegisters {
		fmt.Printf("%v %v value %v min/max %v/%v bytearray %v\n", r.Address, r.Name, r.Value, r.Min, r.Max, r.ByteValue)
	}

	holdingRegisters, _ := ReadAllHoldingRegisters()
	fmt.Println("Holding registers:")
	for _, r := range holdingRegisters {
		fmt.Printf("%v %v value %v min/max %v/%v bytearray %v\n", r.Address, r.Name, r.Value, r.Min, r.Max, r.ByteValue)
	}

	writeValue("SetAirTemperature", 200)
	writeValue("SetAirSpeed", 2)
	register, err := ReadHoldingRegister("SupplyAirSpeed1")
	fmt.Println("register", register.Address, register.Value, "read ", register, err)
}

func writeValue(key string, value int64) {
	register, err := ReadHoldingRegister(key)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	register.Value = value
	fmt.Println("Writing", register.Name, "value", register.Value)
	err = WriteHoldingRegister(*register)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	time.Sleep(100 * time.Millisecond)
	register, err = ReadHoldingRegister(key)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	if register.Value != value {
		fmt.Println("Value was not stored/written, expected", value, "got", register.Value)
	} else {
		fmt.Println("Success, expected", value, "got", register.Value)
	}
}
