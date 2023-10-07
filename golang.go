package Powerbank

import (
	"fmt"
)

type SmartPowerBank struct {
	Capacity     int
	ChargeLevel  int
	Manufacturer string
}

func Powerbank() {
	powerBank := SmartPowerBank{
		Capacity:     10000,
		ChargeLevel:  50,
		Manufacturer: "Power Bank",
	}

	fmt.Printf("Manufacturer: %s\n", powerBank.Manufacturer)
	fmt.Printf("Capacity: %d mAh\n", powerBank.Capacity)
	fmt.Printf("Charge Level: %d%%\n", powerBank.ChargeLevel)
}
