package main

import (
	"encoding/json"
	"fmt"
)

type Persona struct {
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
}

func main() {
	// Serializar struct a JSON
	persona := Persona{Nombre: "Juan", Edad: 30}
	jsonBytes, err := json.Marshal(persona)
	if err != nil {
		fmt.Println("Error al serializar JSON:", err)
		return
	}
	fmt.Println("JSON serializado:", string(jsonBytes))

	// Deserializar JSON a struct
	var nuevaPersona Persona
	err = json.Unmarshal(jsonBytes, &nuevaPersona)
	if err != nil {
		fmt.Println("Error al deserializar JSON:", err)
		return
	}
	fmt.Println("Struct deserializado:", nuevaPersona)
}
