package main

import (
	"fmt"
)

func main() {

	// Variáveis
	// Var + Nome da variavel + Tipo
	var nome string
	nome = "Bruno"
	fmt.Println(nome)

	nome = "Ana"
	fmt.Println(nome)

	// Agora com variável int
	var idade int
	idade = 4
	fmt.Println(idade)

	// Outra maneira de escrever
	var a = "Vitória"
	fmt.Println(a)

	// Aqui mostra que dá pra por mais de um por vez
	var b, c int = 100, 1000
	fmt.Println(b)
	fmt.Println(c)

	// Também serve pra booleano
	var d = true
	fmt.Println(d)

	// Esse daqui vai ser mais usado
	m := "Maça"
	fmt.Println(m)

	// Esse modo é mais fácil
	x := 1
	fmt.Println(x)

	k := 1.2
	fmt.Println(k)

	// Constantes
	const Pi = 3.1415
	fmt.Println(Pi)
}
