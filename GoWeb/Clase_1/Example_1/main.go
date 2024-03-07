package main

/*
Vamos a crear una aplicación web con el package net/http nativo de go,
que tenga un endpoint /ping que al pegarle responda un texto que diga “pong”
El endpoint deberá ser de método GET
La respuesta de “pong” deberá ser enviada como texto, NO como JSON
*/

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})
	http.ListenAndServe(":8080", nil)
}
