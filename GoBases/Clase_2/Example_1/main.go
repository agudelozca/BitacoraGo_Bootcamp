package main

import "fmt"

func main() {
	word := "cristian"
	lon := len(word)
	fmt.Printf("La palabra %s contiene %d letras\n", word, lon)

	for _, j := range word {
		fmt.Printf("%c ", j)
	}
}
