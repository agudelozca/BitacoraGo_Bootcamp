package main

import (
	"fmt"
	"time"
)

type Alumnos struct {
	Name     string
	Apellido string
	DNI      int
	Fecha    time.Time
}

func (a *Alumnos) detalle() {
	fmt.Printf("Nombre: %s\nApellido: %s\nDNI: %d\nFecha: %s\n", a.Name, a.Apellido, a.DNI, a.Fecha)
}

func main() {
	Alumno1 := Alumnos{
		Name:     "Juan",
		Apellido: "Perez",
		DNI:      12345678,
		Fecha:    time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	Alumno1.detalle()

	Alumno2 := Alumnos{
		Name:     "Cristian",
		Apellido: "Agudelo",
		DNI:      11111111,
		Fecha:    time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
	}
	Alumno2.detalle()
}
