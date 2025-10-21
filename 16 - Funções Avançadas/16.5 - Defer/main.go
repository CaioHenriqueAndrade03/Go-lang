package main

import "fmt"

func f1() {
	fmt.Println("Executando a funcao 1")
}

func f2() {
	fmt.Println("Executando a funcao 2")
}
func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media calculada, resultado sera retornado")
	media := (n1 + n2) / 2
	if media >= 6 {
		return true
	}
	return false
}
func main() {
	fmt.Println("16.5 - Defer")
	fmt.Println(alunoEstaAprovado(7, 8))
}
