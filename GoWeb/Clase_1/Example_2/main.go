/*
Vamos a crear un endpoint llamado /greetings.
Con una pequeña estructura con nombre y apellido
que al pegarle deberá responder en texto “Hello + nombre + apellido”
El endpoint deberá ser de método POST
Se deberá usar el package JSON para resolver el ejercicio
La respuesta deberá seguir esta estructura: “Hello Andrea Rivas”
La estructura deberá ser como esta:

	{
	                “firstName”: “Andrea”,

	                “lastName”: “Rivas”
	}
*/

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/greetings", func(w http.ResponseWriter, r *http.Request) {
		var person struct {
			FirstName string `json:"firstName"`
			LastName  string `json:"lastName"`
		}
		// Decodificamos el cuerpo JSON de la solicitud en la estructura person.
		err := json.NewDecoder(r.Body).Decode(&person)
		// Si hay un error al decodificar el JSON, respondemos con un código de error 400.
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// Construimos la respuesta con un saludo utilizando los campos FirstName y LastName.
		fmt.Fprintf(w, "Hello %s %s", person.FirstName, person.LastName)
	})
	// Iniciamos el servidor en el puerto 8080.
	http.ListenAndServe(":8080", nil)
}
