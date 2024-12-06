package main

import "fmt"

// Deklarqation einer Funktion zur Berechnung der Fläche eines Rechtecksaus den Seitenlängen
func rechteck(seiteA int, seiteB int) int {
	var flaeche int = seiteA * seiteB
	return flaeche
}

// Deklaration einer Funktion zur Berechnung des Volumens eines Quaders aus Grundfläche und Höhe
func quader(grundflaeche int, hoehe int) int {
	var volumen int = grundflaeche * hoehe
	return volumen
}

func main() {
	var a int = 3
	var b int = 4
	var c int = 5
	fmt.Println("Fläche eines", a, "x", b, "Rechtecks:", rechteck(a, b))
	fmt.Println("Volumen eines Quaders mit Grundfläche", a, "x", b, "und Höhe", c, ":", quader(rechteck(a, b), c))
}
