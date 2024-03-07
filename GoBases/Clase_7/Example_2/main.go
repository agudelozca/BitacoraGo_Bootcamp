package main

import (
	"fmt"
	"os"
)

func main() {
	path := "customer.txt"
	ReadFilePath(path)
	fmt.Println("Ejecucion finalizada")
}

func ReadFilePath(path string) *os.File {
	file, err := os.Open(path)
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		panic("The indicated file was not found or is damaged")
	}
	return file
}
