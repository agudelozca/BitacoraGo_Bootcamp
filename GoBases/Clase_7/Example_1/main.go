package main

import "fmt"

func main() {
	// Se utiliza recover en conjunto con defer para manejar el panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	// Llamada a una función que causa un panic
	causePanic()
	fmt.Println("This line will not be executed.")
}

func causePanic() {
	// Causa deliberate un panic
	panic("Something went wrong!")
}
