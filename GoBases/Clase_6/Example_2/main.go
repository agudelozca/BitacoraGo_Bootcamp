package main

import (
	"errors"
	"fmt"
)

/*
En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.

Creá un error personalizado con un struct que implemente “Error()”
con el mensaje “Error: salary is less than 10000" y lanzalo en caso de que “salary”

	sea menor o igual a  10000. La validación debe ser hecha con la función “Is()”
	dentro del “main”.
*/
type customError struct {
	Message string
}

func (c *customError) Error() string {
	return c.Message
}

func main() {
	var salary int = 100000
	if salary <= 10000 {
		err := &customError{
			Message: "Error: salary is less than 10000",
		}
		// Lanzar el error
		if errors.Is(err, &customError{}) {
			fmt.Println(err)
		}
	} else {
		fmt.Println("Salary valid")
	}
}
