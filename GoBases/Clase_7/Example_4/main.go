package main

import (
	"errors"
	"fmt"
)

type Client struct {
	File        string
	Name        string
	ID          int
	PhoneNumber string
	Home        string
}

var clientDatabase []Client

func checkClientExistence(newClientID int) {
	for _, existingClient := range clientDatabase {
		if existingClient.ID == newClientID {
			// Genera un panic y lanza un mensaje
			panic("Error: client already exists")
		}
	}
}

func validateClientData(clientData Client) (Client, error) {
	if clientData.ID == 0 || clientData.PhoneNumber == "" || clientData.Home == "" || clientData.Name == "" || clientData.File == "" {
		return Client{}, errors.New("Error: all client data must have non-zero values")
	}
	return clientData, nil
}

func registerNewClient(newClient Client) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			fmt.Println("End of execution")
			fmt.Println("Several errors were detected at runtime")
		} else {
			fmt.Println("End of execution")
		}
	}()

	// Tarea 1: Verifica si el cliente ya existe
	checkClientExistence(newClient.ID)

	// Tarea 2: Valida que todos los datos del cliente tengan valores distintos de cero
	_, err := validateClientData(newClient)
	if err != nil {
		panic(err.Error())
	}

	// Agrega el nuevo cliente a la base de datos
	clientDatabase = append(clientDatabase, newClient)

	fmt.Println("Client registered successfully")
}

func main() {
	// Ejemplo de registro de nuevo cliente
	newClient := Client{
		File:        "A123",
		Name:        "John Doe",
		ID:          123,
		PhoneNumber: "555-1234",
		Home:        "123 Main St",
	}

	registerNewClient(newClient)

	newClient2 := newClient
	registerNewClient(newClient2)
	fmt.Println(clientDatabase)
}
