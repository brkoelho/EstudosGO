package main

import (
	"fmt"
)

// Structs

// Forma de criar sua própria estrutura de dados
// Personalizar de acordo com a sua necessidade
// Podemos usar vários tipos diferentes
// Formato=> type <nome da nossa estrutura> struct { <campos> }

type Pessoa struct {
	Nome  string
	Idade int
}

// Para fazer struct de struct na última parte
type Profissao struct {
	Pessoa
	Tipo string
}

func main() {
	fmt.Println(Pessoa{"Bruno", 34})
	fmt.Println(Pessoa{Nome: "Ana", Idade: 26})
	// Se nada colocar para idade, ele vai colocar um Zero-Value auto
	fmt.Println(Pessoa{Nome: "Vitória"})
	// Nada foi colocado para string Nome, logo ele coloca Zero-Value no print
	fmt.Println(Pessoa{Idade: 40})

	// Outra maneira de escrever
	p1 := Pessoa{Nome: "Chico", Idade: 45}
	fmt.Println(p1.Nome)
	fmt.Println(p1.Idade)

	p1.Idade = 48
	fmt.Println(p1.Idade)

	p2 := Pessoa{Nome: "Saike", Idade: 27}

	// Aqui é uma criação de Slice de pessoas
	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	fmt.Println(pessoas)

	// Structs + Map
	// alunos := map[string][]Pessoa{}
	// alunos["programação"] = pessoas
	// fmt.Println(alunos)

	// Outro modo - Exemplo
	var alunos = map[string][]Pessoa{
		"Programação": {{Nome: "Steph", Idade: 28}, {Nome: "Bento", Idade: 4}},
		"Engenharia":  {{Nome: "Steph2", Idade: 28}, {Nome: "Bento2", Idade: 4}},
	}
	fmt.Println(alunos)

	// Struct herdando campos de outra struct - tipo uma cadeia um puxando os dados do outro
	prof := Profissao{p2, "dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Nome)
	fmt.Println(prof.Pessoa.Idade)
	fmt.Println(prof.Tipo)
}
