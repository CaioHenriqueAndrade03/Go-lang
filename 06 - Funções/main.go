package main

import "fmt"

func main() {
	fmt.Println("06 - Funções")
	soma := somar(8, 10)
	fmt.Println(soma)

	f()
	g("oiee")
	resultado := g("Texto da func G")
	fmt.Println(resultado)
	resultadoSoma, resultadoSubtracao, resultadoMultiplicacao, resultadoDivisao := calculosMatematicos(20, 10)
	fmt.Println(resultadoDivisao)
	fmt.Println(resultadoMultiplicacao)
	fmt.Println(resultadoDivisao)
	fmt.Println(resultadoSoma)
	fmt.Println(resultadoSubtracao)
}

func somar(n1 int, n2 int) int {
	return n1 + n2
}

var f = func() {
	fmt.Println("funçao f")
}

var g = func(txt string) string {
	fmt.Println(txt)
	return "esse é o retorno"
}

func calculosMatematicos(n1, n2 int) (int, int, int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	multiplicacao := n1 * n2
	divisao := n1 / n2
	return soma, subtracao, multiplicacao, divisao
}
