package main

import "fmt"

func main() {
	var tempC float64
	const nullpunkt float64 = 273.15

	tempC = 10
	fmt.Println("Temperatur zum Start     (°C):", tempC)
	tempC = tempC + 12
	fmt.Println("Temperatur nach Änderung (°C):", tempC)

	var tempF float64
	tempF = tempC*1.8 + 32

	var tempK float64 = tempC - nullpunkt

	var tempText string = "Temperatur in "
	fmt.Println()
	fmt.Println(tempText+"°C:", tempC)
	fmt.Println(tempText+"°F:", tempF)
	fmt.Println(tempText+"°K:", tempK)
}
