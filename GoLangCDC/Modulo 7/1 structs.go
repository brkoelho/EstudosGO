package main

import (
	"fmt"
)

type Pessoa struct {
	Nome      string
	Idade     int
	Profissao string
}

// Esse foi modo inicial //
// func ListaNomeEIdade(nome string, idade int) string {
// 	return fmt.Sprintf("%s tem %d anos", nome, idade)
// }

// func main() {

// 	pessoa := Pessoa{
// 		Nome:      "Bruno",
// 		Idade:     34,
// 		Profissao: "aprendev",
// 	}

// 	fmt.Println(ListaNomeEIdade(pessoa.Nome, pessoa.Idade))
// }

func (p Pessoa) ListaNomeEIdade() string { // (p Pessoa) receiver
	return fmt.Sprintf("%s tem %d anos!", p.Nome, p.Idade)
}

func main() {

	pessoa := Pessoa{
		Nome:      "Bruno",
		Idade:     34,
		Profissao: "aprendev",
	}

	fmt.Println(pessoa.ListaNomeEIdade())
}
