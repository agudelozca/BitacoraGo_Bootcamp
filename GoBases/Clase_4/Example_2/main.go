package main

/*
Ejercicio 2 - Employee
Una empresa necesita realizar una buena gestión de sus empleados, para esto
realizaremos un pequeño programa nos ayudará a gestionar correctamente dichos
empleados. Los objetivos son:
Crear una estructura Person con los campos ID, Name, DateOfBirth.
Crear una estructura Employee con los campos: ID, Position y una composición con la estructura Person.
Realizar el método a la estructura Employe que se llame PrintEmployee(), lo que hará es realizar
la impresión de los campos de un empleado.
Instanciar en la función main() tanto una Person como un Employee cargando sus respectivos campos
y por último ejecutar el método PrintEmployee().
Si logras realizar este pequeño programa pudiste ayudar a la empresa a solucionar la gestión de los empleados.
*/

import (
	"fmt"
	"time"
)

// Definir la estructura Person
type Person struct {
	ID          int
	Name        string
	DateOfBirth time.Time
}

// Definir la estructura Employee con composición de la estructura Person
type Employee struct {
	ID       int
	Position string
	Person   Person
}

// Método PrintEmployee asociado a la estructura Employee
func (e *Employee) PrintEmployee() {
	fmt.Printf("Employee ID: %d\nPosition: %s\nPerson ID: %d\nName: %s\nDate of Birth: %s\n",
		e.ID, e.Position, e.Person.ID, e.Person.Name, e.Person.DateOfBirth.Format("2006-01-02"))
}

func main() {
	// Instanciar una Person
	person := Person{
		ID:          1,
		Name:        "John Doe",
		DateOfBirth: time.Date(1990, time.January, 15, 0, 0, 0, 0, time.UTC),
	}

	// Instanciar un Employee con composición de la Person
	employee := Employee{
		ID:       101,
		Position: "Software Engineer",
		Person:   person,
	}

	// Ejecutar el método PrintEmployee() para imprimir los campos del empleado
	employee.PrintEmployee()
}
