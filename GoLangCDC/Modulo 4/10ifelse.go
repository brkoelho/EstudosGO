package main

import (
	"fmt"
)

// IF (se) | ELSE (senão)

func main() {

	// Exemplo if (condição) { <ação> }
	N := 2
	if N == 1 {
		fmt.Println("Valor é igual a 1")
	} else {
		fmt.Println("Valor não é igual a 1")
	}

	// outro exemplo com if, else if e else
	X := 2
	if X == 1 {
		fmt.Println("Valor é igual a 1")
	} else if X == 2 {
		fmt.Println("Valor é igual a 2")
	} else {
		fmt.Println("Valor é diferente de 1 e 2")
	}

	// Aqui é aquele exemplo do resto da divisão
	numero := 8
	if numero%2 == 0 {
		fmt.Printf("%d é par\n", numero) //Esse %d remete ao número
	} else {
		fmt.Printf("%d é impar\n", numero) //Esse %d remete ao número
	}

	Nick := "Brk"
	if Nick == "Brk" {
		fmt.Println("Nick show!")
	}

	// Operações básicas...
	// fmt.Println(1 + 1)
	// fmt.Println(2 - 1)
	// fmt.Println(2 * 2)
	// fmt.Println(10 / 2)
	// Resto da divisão
	// fmt.Println(10 % 2)
}
