package main

import "fmt"

func main() {
	// Variable de entrada
	numeroMes := 7

	// Validación del número del mes
	if numeroMes < 1 || numeroMes > 12 {
		fmt.Println("Número de mes no válido.")
	} else {
		// Nombres de los meses
		meses := [12]string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio", "Julio", "Agosto", "Septiembre", "Octubre", "Noviembre", "Diciembre"}

		// Imprimir el mes correspondiente
		fmt.Printf("%d, %s\n", numeroMes, meses[numeroMes-1])
	}
}
