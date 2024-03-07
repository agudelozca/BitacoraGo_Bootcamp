package main

import "fmt"

func main() {
	// Variables de entrada
	edad := 25
	empleado := true
	antiguedad := 2
	sueldo := 120000

	// Verificación de elegibilidad para préstamo
	if edad > 22 && empleado && antiguedad > 1 {
		fmt.Println("El cliente es elegible para un préstamo.")

		// Verificación de exención de interés
		if sueldo > 100000 {
			fmt.Println("¡Además, este cliente no pagará intereses!")
		} else {
			fmt.Println("Este cliente pagará intereses.")
		}
	} else {
		fmt.Println("El cliente no cumple con los requisitos para un préstamo.")
	}
}
