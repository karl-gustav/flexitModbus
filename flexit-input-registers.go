package flexitModbus

func GetAllInputRegisters() map[string]ReableRegister {
	return map[string]ReableRegister{
		"GWYVer": &UInt16Register{
			BaseRegister: BaseRegister{
				Address:     0,
				Name:        "GWYVer",
				Format:      "UINT 16",
				Description: "Protokollversion.",
				Write:       false,
				Unit:        "",
			},
			Default: 0,
			Min:     0,
			Max:     0xffff,
		},

		"CUHWType": &UInt16Register{
			BaseRegister: BaseRegister{
				Address:     1,
				Name:        "CUHWType",
				Format:      "UINT 16",
				Description: "Main board hardware type",
				Write:       false,
				Unit:        "",
			},
			Default: 0,
			Min:     0,
			Max:     0xffff,
		},

		"CUSWRev": &UInt16Register{
			BaseRegister: BaseRegister{
				Address:     2,
				Name:        "CUSWRev",
				Format:      "UINT 16",
				Description: "Main board software rev.",
				Write:       false,
				Unit:        "",
			},
			Default: 0,
			Min:     0,
			Max:     0xffff,
		},

		"CPASWRev": &UInt16Register{
			BaseRegister: BaseRegister{
				Address:     3,
				Name:        "CPASWRev",
				Format:      "UINT 16",
				Description: "CPA board software rev.",
				Write:       false,
				Unit:        "",
			},
			Default: 0,
			Min:     0,
			Max:     0xffff,
		},

		"CPB1SWRev": &UInt16Register{
			BaseRegister: BaseRegister{
				Address:     4,
				Name:        "CPB1SWRev",
				Format:      "UINT 16",
				Description: "CPB1 board software rev. (Norm. GB)",
				Write:       false,
				Unit:        "",
			},
			Default: 0,
			Min:     0,
			Max:     0xffff,
		},

		"CBPS2WRev": &UInt16Register{
			BaseRegister: BaseRegister{
				Address:     5,
				Name:        "CBPS2WRev",
				Format:      "UINT 16",
				Description: "CPB2 board software rev. (if installed)",
				Write:       false,
				Unit:        "",
			},
			Default: 0,
			Min:     0,
			Max:     0xffff,
		},

		"Time1": &UInt32Register{
			BaseRegister: BaseRegister{
				Address:     6,
				Name:        "Time1",
				Format:      "UINT 32",
				Description: "Real time clock value",
				Write:       false,
				Unit:        "",
			},
			Default: 0,
			Min:     0,
			Max:     0xffffffff,
		},

		"FilterTimer": &UInt16Register{
			BaseRegister: BaseRegister{
				Address:     8,
				Name:        "FilterTimer",
				Format:      "UINT 16",
				Description: "Filter timer",
				Write:       false,
				Unit:        "h",
			},
			Default: 0,
			Min:     0,
			Max:     0xffff,
		},

		"SupplyAirTemp": &TemperatureRegister{
			BaseRegister: BaseRegister{
				Address:     9,
				Name:        "SupplyAirTemp",
				Format:      "INT 16",
				Description: "Supply air temperature",
				Write:       false,
				Unit:        "1°C",
			},
			Default: 20,
			Min:     -45,
			Max:     50,
		},

		"ExtractAirTemp": &TemperatureRegister{
			BaseRegister: BaseRegister{
				Address:     10,
				Name:        "ExtractAirTemp",
				Format:      "INT 16",
				Description: "Extract air temperature",
				Write:       false,
				Unit:        "1°C",
			},
			Default: 20,
			Min:     -45,
			Max:     50,
		},

		"OutdoorAirTemp": &TemperatureRegister{
			BaseRegister: BaseRegister{
				Address:     11,
				Name:        "OutdoorAirTemp",
				Format:      "INT 16",
				Description: "Outdoor air temperature",
				Write:       false,
				Unit:        "1°C",
			},
			Default: 20,
			Min:     -45,
			Max:     50,
		},

		"ReturnWaterTemp": &TemperatureRegister{
			BaseRegister: BaseRegister{
				Address:     12,
				Name:        "ReturnWaterTemp",
				Format:      "INT 16",
				Description: "Return water temperature",
				Write:       false,
				Unit:        "1°C",
			},
			Default: 20,
			Min:     -45,
			Max:     50,
		},

		"Cooling": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     13,
				Name:        "Cooling",
				Format:      "INT 16",
				Description: "Cooling",
				Write:       false,
				Unit:        "%",
			},
			Default: 45,
			Min:     0,
			Max:     100,
		},

		"HeatExchanger": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     14,
				Name:        "HeatExchanger",
				Format:      "INT 16",
				Description: "Heat exchanger",
				Write:       false,
				Unit:        "%",
			},
			Default: 45,
			Min:     0,
			Max:     100,
		},

		"Heating": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     15,
				Name:        "Heating",
				Format:      "INT 16",
				Description: "Heating",
				Write:       false,
				Unit:        "%",
			},
			Default: 45,
			Min:     0,
			Max:     100,
		},

		"RegulationFanSpeed": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     16,
				Name:        "RegulationFanSpeed",
				Format:      "INT 16",
				Description: "Regulation fan speed",
				Write:       false,
				Unit:        "",
			},
			Default: 0,
			Min:     0,
			Max:     3,
		},

		"OperTime": &UInt16Register{
			BaseRegister: BaseRegister{
				Address:     17,
				Name:        "OperTime",
				Format:      "UINT 16",
				Description: "Operational time",
				Write:       false,
				Unit:        "h",
			},
			Default: 0,
			Min:     0,
			Max:     0xffff,
		},

		"FilterResetNo": &UInt16Register{
			BaseRegister: BaseRegister{
				Address:     18,
				Name:        "FilterResetNo",
				Format:      "UINT 16",
				Description: "The number of filter resets",
				Write:       false,
				Unit:        "",
			},
			Default: 0,
			Min:     0,
			Max:     0xffff,
		},

		"SupplyAirAlarm": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     19,
				Name:        "SupplyAirAlarm",
				Format:      "INT 16",
				Description: "Supply air sensor range alarm",
				Write:       false,
				Unit:        "bool",
			},
			Default: 0,
			Min:     0,
			Max:     1,
		},

		"ExtractAirAlarm": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     20,
				Name:        "ExtractAirAlarm",
				Format:      "INT 16",
				Description: "Extract air sensor range alarm",
				Write:       false,
				Unit:        "bool",
			},
			Default: 0,
			Min:     0,
			Max:     1,
		},

		"OutsideAirAlarm": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     21,
				Name:        "OutsideAirAlarm",
				Format:      "INT 16",
				Description: "Outside air sensor range alarm",
				Write:       false,
				Unit:        "bool",
			},
			Default: 0,
			Min:     0,
			Max:     1,
		},

		"ReturnWaterAlarm": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     22,
				Name:        "ReturnWaterAlarm",
				Format:      "INT 16",
				Description: "Return water sensor range alarm",
				Write:       false,
				Unit:        "bool",
			},
			Default: 0,
			Min:     0,
			Max:     1,
		},

		"FireThermostatAlarm": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     23,
				Name:        "FireThermostatAlarm",
				Format:      "INT 16",
				Description: "Fire thermostat active alarm",
				Write:       false,
				Unit:        "bool",
			},
			Default: 0,
			Min:     0,
			Max:     1,
		},

		"FireSmokeAlarm": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     24,
				Name:        "FireSmokeAlarm",
				Format:      "INT 16",
				Description: "Fire/smoke detector active alarm",
				Write:       false,
				Unit:        "bool",
			},
			Default: 0,
			Min:     0,
			Max:     1,
		},

		"FreezeProtectionAlarm": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     25,
				Name:        "FreezeProtectionAlarm",
				Format:      "INT 16",
				Description: "Freeze proctection return water alarm (low return water temperature)",
				Write:       false,
				Unit:        "bool",
			},
			Default: 0,
			Min:     0,
			Max:     1,
		},

		"RotorAlarm": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     26,
				Name:        "RotorAlarm",
				Format:      "INT 16",
				Description: "Rotor alarm",
				Write:       false,
				Unit:        "bool",
			},
			Default: 0,
			Min:     0,
			Max:     1,
		},

		"ReplaceFilterAlarm": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     27,
				Name:        "ReplaceFilterAlarm",
				Format:      "INT 16",
				Description: "Replace filters alarm",
				Write:       false,
				Unit:        "bool",
			},
			Default: 0,
			Min:     0,
			Max:     1,
		},

		"HeatingBatteryActive": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     28,
				Name:        "HeatingBatteryActive",
				Format:      "INT 16",
				Description: "Heating battery active",
				Write:       false,
				Unit:        "bool",
			},
			Default: 1,
			Min:     0,
			Max:     1,
		},

		"SchActive": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     29,
				Name:        "SchActive",
				Format:      "INT 16",
				Description: "Scheduler active",
				Write:       false,
				Unit:        "bool",
			},
			Default: 0,
			Min:     0,
			Max:     1,
		},

		"SP0TimeH": &UInt32Register{
			BaseRegister: BaseRegister{
				Address:     30,
				Name:        "SP0TimeH",
				Format:      "UINT 32",
				Description: "SP0 Time counter",
				Write:       false,
				Unit:        "s",
			},
			Default: 0,
			Min:     0,
			Max:     0xffffffff,
		},

		"SP1TimeH": &UInt32Register{
			BaseRegister: BaseRegister{
				Address:     32,
				Name:        "SP1TimeH",
				Format:      "UINT 32",
				Description: "SP1 Time counter",
				Write:       false,
				Unit:        "s",
			},
			Default: 0,
			Min:     0,
			Max:     0xffffffff,
		},

		"SP2TimeH": &UInt32Register{
			BaseRegister: BaseRegister{
				Address:     34,
				Name:        "SP2TimeH",
				Format:      "UINT 32",
				Description: "SP2 Time counter",
				Write:       false,
				Unit:        "s",
			},
			Default: 0,
			Min:     0,
			Max:     0xffffffff,
		},

		"SP3TimeH": &UInt32Register{
			BaseRegister: BaseRegister{
				Address:     36,
				Name:        "SP3TimeH",
				Format:      "UINT 32",
				Description: "SP3 Time counter",
				Write:       false,
				Unit:        "s",
			},
			Default: 0,
			Min:     0,
			Max:     0xffffffff,
		},

		"VVX1TimeH": &UInt32Register{
			BaseRegister: BaseRegister{
				Address:     38,
				Name:        "VVX1TimeH",
				Format:      "UINT 32",
				Description: "VVX1 Time counter",
				Write:       false,
				Unit:        "s",
			},
			Default: 0,
			Min:     0,
			Max:     0xffffffff,
		},

		"EV1TimeH": &UInt32Register{
			BaseRegister: BaseRegister{
				Address:     40,
				Name:        "EV1TimeH",
				Format:      "UINT 32",
				Description: "EV1 Time counter High",
				Write:       false,
				Unit:        "s",
			},
			Default: 0,
			Min:     0,
			Max:     0xffffffff,
		},

		"OperTimeH": &UInt32Register{
			BaseRegister: BaseRegister{
				Address:     42,
				Name:        "OperTimeH",
				Format:      "UINT 32",
				Description: "Operational counter High",
				Write:       false,
				Unit:        "s",
			},
			Default: 0,
			Min:     0,
			Max:     0xffffffff,
		},

		"FilterTimeH": &UInt32Register{
			BaseRegister: BaseRegister{
				Address:     44,
				Name:        "FilterTimeH",
				Format:      "UINT 32",
				Description: "Filter counter High",
				Write:       false,
				Unit:        "s",
			},
			Default: 0,
			Min:     0,
			Max:     0xffffffff,
		},

		"FilterAlarmPeriod": &UInt16Register{
			BaseRegister: BaseRegister{
				Address:     46,
				Name:        "FilterAlarmPeriod",
				Format:      "UINT 16",
				Description: "Filter alarm, time period",
				Write:       false,
				Unit:        "s",
			},
			Default: 180,
			Min:     0,
			Max:     360,
		},

		"ActualSetAirTemperature": &TemperatureRegister{
			BaseRegister: BaseRegister{
				Address:     47,
				Name:        "ActualSetAirTemperature",
				Format:      "INT 16",
				Description: "The set air temperature used on CU",
				Write:       false,
				Unit:        "1°C",
			},
			Default: 20.0,
			Min:     0,
			Max:     25.0,
		},

		"ActualSetAirSpeed": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     48,
				Name:        "ActualSetAirSpeed",
				Format:      "INT 16",
				Description: "The set air speed used on CU",
				Write:       false,
				Unit:        "",
			},
			Default: 2,
			Min:     0,
			Max:     3,
		},
	}
}
