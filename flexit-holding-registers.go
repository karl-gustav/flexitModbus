package flexitModbus

func GetAllHoldingRegisters() map[string]WritableRegister {
	return map[string]WritableRegister{
		"SupplyAirSpeed1": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     0,
				Name:        "SupplyAirSpeed1",
				Format:      "INT 16",
				Description: "Speed 1 (sa)",
				Write:       true,
				Unit:        "%",
			},
			Default: 50,
			Min:     20,
			Max:     100,
		},

		"SupplyAirSpeed2": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     1,
				Name:        "SupplyAirSpeed2",
				Format:      "INT 16",
				Description: "Speed 2 (sa)",
				Write:       true,
				Unit:        "%",
			},
			Default: 75,
			Min:     20,
			Max:     100,
		},

		"SupplyAirSpeed3": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     2,
				Name:        "SupplyAirSpeed3",
				Format:      "INT 16",
				Description: "Speed 3 (sa)",
				Write:       true,
				Unit:        "%",
			},
			Default: 100,
			Min:     20,
			Max:     100,
		},

		"SupplyAirSpeed4": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     3,
				Name:        "SupplyAirSpeed4",
				Format:      "INT 16",
				Description: "Speed 4 (sa)",
				Write:       true,
				Unit:        "",
			},
			Default: 3,
			Min:     1,
			Max:     3,
		},

		"ExtractAirSpeed1": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     4,
				Name:        "ExtractAirSpeed1",
				Format:      "INT 16",
				Description: "Speed 1 (ea)",
				Write:       true,
				Unit:        "%",
			},
			Default: 50,
			Min:     20,
			Max:     100,
		},

		"ExtractAirSpeed2": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     5,
				Name:        "ExtractAirSpeed2",
				Format:      "INT 16",
				Description: "Speed 2 (ea)",
				Write:       true,
				Unit:        "%",
			},
			Default: 75,
			Min:     20,
			Max:     100,
		},

		"ExtractAirSpeed3": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     6,
				Name:        "ExtractAirSpeed3",
				Format:      "INT 16",
				Description: "Speed 3 (ea)",
				Write:       true,
				Unit:        "%",
			},
			Default: 100,
			Min:     20,
			Max:     100,
		},

		"ExtractAirSpeed4": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     7,
				Name:        "ExtractAirSpeed4",
				Format:      "INT 16",
				Description: "Speed 4 (ea)",
				Write:       true,
				Unit:        "",
			},
			Default: 1,
			Min:     1,
			Max:     3,
		},

		"SetAirTemperature": &TemperatureRegister{
			BaseRegister: BaseRegister{
				Address:     8,
				Name:        "SetAirTemperature",
				Format:      "INT 16",
				Description: "Air temperature",
				Write:       true,
				Unit:        "1°C",
			},
			Default: 20.0,
			Min:     10.0,
			Max:     30.0,
		},

		"SupplyAirMinTemp": &TemperatureRegister{
			BaseRegister: BaseRegister{
				Address:     9,
				Name:        "SupplyAirMinTemp",
				Format:      "INT 16",
				Description: "Min supply air temp.",
				Write:       true,
				Unit:        "1°C",
			},
			Default: 16.0,
			Min:     5.0,
			Max:     25.0,
		},

		"SupplyAirMaxTemp": &TemperatureRegister{
			BaseRegister: BaseRegister{
				Address:     10,
				Name:        "SupplyAirMaxTemp",
				Format:      "INT 16",
				Description: "Max supply air temp.",
				Write:       true,
				Unit:        "1°C",
			},
			Default: 35.0,
			Min:     15.0,
			Max:     45.0,
		},

		"CoolingOutdoorAirMinTemp": &TemperatureRegister{
			BaseRegister: BaseRegister{
				Address:     11,
				Name:        "CoolingOutdoorAirMinTemp",
				Format:      "INT 16",
				Description: "CO, Min outdoor temp.",
				Write:       true,
				Unit:        "1°C",
			},
			Default: 17.0,
			Min:     5.0,
			Max:     25.0,
		},

		"ForcedVentSpeed": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     12,
				Name:        "ForcedVentSpeed",
				Format:      "INT 16",
				Description: "Speed to set during forced ventilation (max timer)",
				Write:       true,
				Unit:        "",
			},
			Default: 3,
			Min:     1,
			Max:     3,
		},

		"ForcedVentTime": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     13,
				Name:        "ForcedVentTime",
				Format:      "INT 16",
				Description: "Forced ventilation (max timer) activation time",
				Write:       true,
				Unit:        "min",
			},
			Default: 30,
			Min:     0,
			Max:     360,
		},

		"AirRegulationType": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     14,
				Name:        "AirRegulationType",
				Format:      "INT 16",
				Description: "Air regulation type",
				Write:       true,
				Unit:        "bool",
			},
			Default: 1,
			Min:     0,
			Max:     1,
		},

		"CoolingActive": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     15,
				Name:        "CoolingActive",
				Format:      "INT 16",
				Description: "Cooling (CO)",
				Write:       true,
				Unit:        "bool",
			},
			Default: 0,
			Min:     0,
			Max:     1,
		},

		"ForcedVentilation": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     16,
				Name:        "ForcedVentilation",
				Format:      "INT 16",
				Description: "Forced ventilation activate/deactivate",
				Write:       true,
				Unit:        "bool",
			},
			Default: 0,
			Min:     0,
			Max:     1,
		},

		"SetAirSpeed": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     17,
				Name:        "SetAirSpeed",
				Format:      "INT 16",
				Description: "A set value that panels use to write wanted air speed to CU.",
				Write:       true,
				Unit:        "",
			},
			Default: 2,
			Min:     0,
			Max:     3,
		},

		"Time": &UInt32Register{
			BaseRegister: BaseRegister{
				Address:     18,
				Name:        "Time",
				Format:      "UINT 32",
				Description: "Real time clock value",
				Write:       false,
				Unit:        "s",
			},
			Default: 0,
			Min:     0,
			Max:     0xffffffff,
		},

		"FireSmokeMode": &Int16Register{
			BaseRegister: BaseRegister{
				Address:     21,
				Name:        "FireSmokeMode",
				Format:      "INT 16",
				Description: "Fire/Smoke mode",
				Write:       true,
				Unit:        "",
			},
			Default: 1,
			Min:     1,
			Max:     4,
		},
	}
}
