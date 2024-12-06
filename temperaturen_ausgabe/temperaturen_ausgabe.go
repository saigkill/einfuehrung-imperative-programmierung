package main

import "fmt"

func fNachC(tempF float64) float64 {
	return (tempF - 32) / 1.8
}

func main() {
	var tempF0 float64 = -459.67
	var tempF1 float64 = 0
	var tempF2 float64 = 32
	var tempF3 float64 = 96

	fmt.Println(tempF0, "°F entsprechen", fNachC(tempF0), "°C")
	fmt.Println(tempF1, "°F entsprechen", fNachC(tempF1), "°C")
	fmt.Println(tempF2, "°F entsprechen", fNachC(tempF2), "°C")
	fmt.Println(tempF3, "°F entsprechen", fNachC(tempF3), "°C")
	fmt.Println()
	fmt.Printf("%v °F entsprechen %v °C.\n", tempF0, fNachC(tempF0))
	fmt.Printf("%v °F entsprechen %v °C.\n", tempF1, fNachC(tempF1))
	fmt.Printf("%v °F entsprechen %v °C.\n", tempF2, fNachC(tempF2))
	fmt.Printf("%v °F entsprechen %v °C.\n", tempF3, fNachC(tempF3))
	fmt.Println()
	fmt.Printf("%f °F entsprechen %f °C.\n", tempF0, fNachC(tempF0))
	fmt.Printf("%f °F entsprechen %f °C.\n", tempF1, fNachC(tempF1))
	fmt.Printf("%f °F entsprechen %f °C.\n", tempF2, fNachC(tempF2))
	fmt.Printf("%f °F entsprechen %f °C.\n", tempF3, fNachC(tempF3))
	fmt.Println()
	fmt.Printf("%7.2f °F entsprechen %7.2f °C.\n", tempF0, fNachC(tempF0))
	fmt.Printf("%7.2f °F entsprechen %7.2f °C.\n", tempF1, fNachC(tempF1))
	fmt.Printf("%7.2f °F entsprechen %7.2f °C.\n", tempF2, fNachC(tempF2))
	fmt.Printf("%7.2f °F entsprechen %7.2f °C.\n", tempF3, fNachC(tempF3))
	fmt.Println()
	fmt.Printf("Der Umsatz steigt um %v %% an.\n", 12.5)
	fmt.Printf("Ein Jahr hat %v Wochen.\n", 52)

	var a int = 78
	var b int = 0b1010_0011
	fmt.Printf("%v:  %[1]b  %[1]o  %[1]x\n", a)  // ohne Präfix
	fmt.Printf("%d:  %#[1]b %#[1]o %#[1]x\n", a) // mit Präfix
	fmt.Printf("%v:  %[1]b  %[1]O  %[1]X\n", b)  // Alternative Ausgabe
}
