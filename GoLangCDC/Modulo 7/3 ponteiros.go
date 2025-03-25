package main

import (
	"fmt"
)

// Ponteiros: Um ponteiro nada mais é que uma variável que ao invés
// de amarzenar um valor (1, olá, "true", ...) armazena um endereço de memória

func main() {
	// Memória -> essa memória tem um endereço -> esse endereço guarda um valor

	i := 1
	fmt.Println("Valor inicial:", i)
	fmt.Println("Valor end de memória:", &i) // &var -> passa o endereço de memória

	a := &i
	fmt.Println("Variável a:", a)
}
