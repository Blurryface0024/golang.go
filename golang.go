package main

import (
	"encoding/json"
	"fmt"
)

type SmartPowerBank struct {
	Capacity     int    json:"capacity"
	ChargeLevel  int    json:"charge_level"
	Manufacturer string json:"manufacturer"
}

func main() {
	powerBank := SmartPowerBank{
		Capacity:     10000,
		ChargeLevel:  75,
		Manufacturer: "Power Bank",
	}

	fmt.Printf("Manufacturer: %s\n", powerBank.Manufacturer)
	fmt.Printf("Capacity: %d mAh\n", powerBank.Capacity)
	fmt.Printf("Charge Level: %d%%\n", powerBank.ChargeLevel)

	jsonData, err := json.Marshal(powerBank)
	if err != nil {
		fmt.Println("Ошибка при кодировании в JSON:", err)
		return
	}
	fmt.Println("Информация о Power Bank в JSON:")
	fmt.Println(string(jsonData))

	var decodedPowerBank SmartPowerBank
	err = json.Unmarshal(jsonData, &decodedPowerBank)
	if err != nil {
		fmt.Println("Ошибка при декодировании JSON:", err)
		return
	}

	fmt.Println("Декодированная информация о Power Bank:")
	fmt.Printf("Manufacturer: %s\n", decodedPowerBank.Manufacturer)
	fmt.Printf("Capacity: %d mAh\n", decodedPowerBank.Capacity)
	fmt.Printf("Charge Level: %d%%\n", decodedPowerBank.ChargeLevel)
}