package main

import (
	"fmt"
)

func main() {

	ResultadoSoma := Soma(10, 20)
	fmt.Println(ResultadoSoma)

	nick1, nick2 := PrintaNickCompleto("Brkoelho", "Xaibot")
	fmt.Println(nick1, nick2)

	Name := PrintaNome("Bruno")
	fmt.Println(Name)

	Subz := Sub(20, 30)
	fmt.Println(Subz)
}

// Funções Criadas

// Funções com início minúsculo são privadas, com maiúsculo são públicas

func PrintaNickCompleto(nome, sobrenome string) (string, string) {
	return nome, sobrenome
}

func PrintaNome(nome string) string {
	return nome
}

func Soma(x int, y int) int {
	return x + y
}

func Sub(x int, y int) int {
	return x - y
}
