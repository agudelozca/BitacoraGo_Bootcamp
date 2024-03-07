package main

import (
	"errors"
	"fmt"
)

/*
Hacé lo mismo que en el ejercicio anterior pero reformulando el código para que,

	en reemplazo de “Error()”,  se implemente “errors.New()”
*/
type CustomError struct {
	Message string
}

func main() {
	var salary int = 100

	if salary < 10000 {
		err := errors.New("Error: salary is less than 10000")
		fmt.Println(err)
	} else {
		fmt.Println("Salary is valid")
	}
}
