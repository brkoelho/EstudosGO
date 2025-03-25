package main

import (
	"fmt"
)

func zeroValue(i int) {
	i = 0
	fmt.Println("End memória do i dentro da func:", &i)
}

func zeroPointer(i *int) {
	*i = 0
}

func main() {
	// Memória -> essa memória tem um endereço -> esse endereço guarda um valor

	i := 1
	fmt.Println("Valor inicial:", i)
	fmt.Println("Valor end de memória:", &i) // &var -> passa o endereço de memória

	zeroValue(i)
	fmt.Println("zeroval:", i)

	zeroPointer(&i)
	fmt.Println("zeroptr:", i)
	fmt.Println("pointer:", &i)

	// fmt.Println("Variável a:", a)
	// fmt.Println("Variável *a:", *a) // qdo uso o *, estamos falando do endereço! e não do valor
	// fmt.Println("Variável &a:", &a) // o i tem seu endereço e o a tem seu endereço, mesmo que um esteja recebendo no momento o valor do outro
}
