package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numeros := range numeros {
		fmt.Println(texto, numeros)
	}
}
func main() {
	fmt.Println("16.2 - Funções variaticas")
	total := soma(1, 2, 4, 5, 6)
	escrever("Olá", 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(total)

}
