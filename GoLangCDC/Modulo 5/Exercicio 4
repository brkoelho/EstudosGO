package kata

// Função Para Contar Positivos e Somar Negativos...
// Função de números de uma lista inteira resultando em lista de inteiros
func CountPositivesSumNegatives(numbers []int) []int {
	// Duas variáveis suportes para contar e somar
	var QtdPositivos, SumNegatives int
	// For+Range é usado para percorrer toda lista
	// O underline é para não usar o i e chegar todos índices da lista numbers
	for _, number := range numbers {
		if number > 0 {
			QtdPositivos = QtdPositivos + 1
		}
		if number < 0 {
			SumNegatives = SumNegatives + number
		}
	}
	return []int{QtdPositivos, SumNegatives}
}

// Percorrer a lista (For Range)
// Se n > 0 -> contar qtos -> fazer variável para contar
// Se n < 0 -> iremos somar os negativos
// Retornar no formato de lista: [QtdPositivos, SomaNegativos]


//// Fiz esse novo modo tbm sem erros :D

// package kata

// func CountPositivesSumNegatives(numbers []int) []int {
//   var x, sn, sp int
//   for _, x = range numbers {
//     if x>0 {
//       sp = sp + 1
//     }
//     if x<0 {
//       sn = sn + x
//     }
//   }
//   return []int{sp, sn}
// }
