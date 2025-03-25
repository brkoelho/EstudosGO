package main

import (
	"fmt"
)

// Loops / Laços de repetição / Repetir tarefas
// A cláusula do for em Go é composta de três partes separadas por ponto e vírgula (;): inicialização, condição e incremento.

func main() {

	////// EXEMPLO 1 //////
	// sum := 0
	// // Nesse caso o i tá sendo o índice
	// // i++ -> soma 1 no índice (vai pro próximo)
	// // i-- -> subtraindo 1 (volta pro anterior)
	// // Ele irá repetir o FOR enquanto o i for menor que 10
	// for i := 1; i < 10; i++ {
	// 	fmt.Println("Sum:", sum)
	// 	// Printa o valor que está a soma no momento
	// 	fmt.Println("Indice:", i)
	// 	// Printa o índice que está no momento
	// 	sum += i
	// 	// <Bem comum de usar pra ir somando os valores>
	// 	// Ou seja:
	// 	// é a mesma coisa que: sum = sum + i
	// 	// sum -= i -> sum = sum - i
	// }

	////// EXEMPLO 2 //////
	// Loop infinito -> não tem condição pós 'for'
	// for {
	// 	fmt.Println("Golang do zero")
	// 	// time.sleep é uma função para colocar delay na execução da tarefa ou continuação do loop
	// 	time.Sleep(2 * time.Second)
	// }

	////// TESTE 1 //////
	////// Esse é um loop infinito //////
	// i := 0
	// for {
	// 	fmt.Println("GoLangTeste", i)
	// 	time.Sleep(1 * time.Second)
	// 	i++
	// }

	////// EXEMPLO 3 //////
	//For com Range
	frutas := []string{"laranja", "maça", "banana", "uva", "kiwi"}
	// Esse underline é o índice
	for _, fruta := range frutas {
		fmt.Println("Fruta:", fruta)
	}
}
