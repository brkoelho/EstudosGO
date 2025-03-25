package main

import (
	"fmt"
)

func main() {

	// Array - tamanho fixo
	var array [2]string
	array[0] = "Eai"
	array[1] = "Irmão"
	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println(array[0], array[1])
	fmt.Println(array)

	// Aqui já é com declaração curta
	numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numPrimos)
	fmt.Println(numPrimos[0:4])
	fmt.Println(numPrimos[0:6])
	fmt.Println(numPrimos[2:3])

	// Slices - criados do zero não usa essa abaixo
	//var slice []string
	// Slice não é finito pode aumentar qdo quiser
	slice := make([]string, 6)
	slice[0] = "Cheguei"
	slice[1] = "Brasil"
	fmt.Println(slice[0], slice[1])
	fmt.Println(slice[0])
	fmt.Println(slice[1])
	fmt.Println(slice[2])
	slice[2] = "Baronil"
	fmt.Println(slice[2])
	fmt.Println(slice)

	// Len Slice é quantos slices tem...
	fmt.Println(len(slice))
	fmt.Println(slice[2])
	fmt.Println(slice[3])
	fmt.Println(slice[4])
	fmt.Println(slice[5])

	// Slice Inicial
	numPares := []int{2, 4, 6, 8, 10, 12}
	fmt.Println(numPares)

	// Appendar é adicionar só pode do mesmo tipo
	numPares = append(numPares, 20, 22, 40, 60)
	fmt.Println(numPares)
}

// // LISTAS

// // 1 - Arrays e Slices: Homogêneos
// // todos os elementos tem o mesmo tipo
// // [1, 2, 3, 4, 5, 6] - []int
// // ["steph", "bento", "golang"] - []string

// // 2 - Maps: Heterogêneos
// // pode misturar tipos
// // estrutura chave - valor
// // [key] = value
// // chave tem um tipo, e o valor pode ter outro
// // map[string]int
// // { "steph": 28, "bento": 4}
// // map[string]string
// // { "steph": "cardoso", "bento": "cardoso" }

// // Array

// // Tamanho fixo, de zero ou mais elementos do mesmo tipo
// // Acessamos os valores com índice: a[0], a[1]...
// // Função embutida len() retorna o tamanho do array
// // Por conta do tamanho fixo, não é tão usado. Só em casos específicos

// // Slice

// // Tipo o array, mas sem tamanho fixo
// // Acessamos os valores com índice: a[0], a[1]...
// // Função embutida len() retorna o tamanho do slice
// // Função append() usada para adicionar valores nesse slice
