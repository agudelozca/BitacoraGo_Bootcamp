package main

import "fmt"

func main() {
	// Mapa de empleados
	employees := map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	// Obtener e imprimir la edad de Benjamin
	benjaminAge, benjaminExists := employees["Benjamin"]
	if benjaminExists {
		fmt.Printf("La edad de Benjamin es: %d\n", benjaminAge)
	} else {
		fmt.Println("Benjamin no está en la lista de empleados.")
	}

	// Contar empleados mayores de 21 años
	count := 0
	for _, age := range employees {
		if age > 21 {
			count++
		}
	}
	fmt.Printf("Hay %d empleados mayores de 21 años.\n", count)

	// Agregar un nuevo empleado llamado Federico
	employees["Federico"] = 25
	fmt.Println("Se ha agregado a Federico a la lista de empleados.")

	// Eliminar a Pedro del mapa
	delete(employees, "Pedro")
	fmt.Println("Se ha eliminado a Pedro de la lista de empleados.")

	// Imprimir el mapa actualizado
	fmt.Println("Empleados actualizados:")
	for name, age := range employees {
		fmt.Printf("%s: %d años\n", name, age)
	}
}
