package main

import (
	"fmt"
)

func main() {

	// Exemplo 1
	idade := map[string]int{}
	idade["Bruno"] = 34
	idade["Ana"] = 26
	fmt.Println(idade)
	fmt.Println(idade["Bruno"])
	fmt.Println(idade["Ana"])

	// Exemplo 2
	anoNasc := map[string]int{
		"Bruno": 1990,
		"Ana":   1999,
	}
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["Bruno"])
	fmt.Println(anoNasc["Ana"])

	// Desse jeito aqui add na var anoNasc mais um dado
	anoNasc["golangDoZero"] = 2024
	fmt.Println(anoNasc["golangDoZero"])
	fmt.Println(anoNasc)
}

// 2 - Maps: Heterogêneos
// pode misturar tipos
// estrutura chave - valor
// [key] = value
// chave tem um tipo, e o valor pode ter outro ou o mesmo
// map[k]v -> k = chave, v = valor

// Exemplo
// map[string]int
// { "steph": 28, "bento": 4}
// map[string]string
// { "steph": "cardoso", "bento": "cardoso" }
