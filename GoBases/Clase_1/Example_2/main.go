package main

import "fmt"

func main() {
	/*
		Una empresa de meteorología quiere una aplicación donde pueda tener
		la temperatura, humedad y presión atmosférica de distintos lugares.
	*/
	var temperature, humidity, pressure float32 = 21, 77, 1023
	fmt.Println("El clima en sabaneta esta de la siguiente manera:\nLa temperatura es de:", temperature, "ºC", "\nLa humedad esta en:", humidity, "%", "\nLa presion esta en:", pressure, "mb")

}
