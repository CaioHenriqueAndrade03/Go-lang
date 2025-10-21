package main

import "fmt"

func recuperarExecução() {
	if r := recover(); r != nil {
		fmt.Println("tentando recuperar")
	}
}
func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperarExecução()
	defer fmt.Println("Media calculada, resultado sera retornado")
	media := (n1 + n2) / 2
	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	panic("A media é exatamente 6")
}
func main() {
	fmt.Println("16.6 - Panic e Recover")
	fmt.Println(alunoEstaAprovado(7, 8))
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução")
}
