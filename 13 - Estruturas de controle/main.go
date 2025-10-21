package main

import "fmt"

func main() {
	fmt.Println("13 - Estruturas de controle")

	numero := 30

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual que 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Numero é maior que 0")
	} else {
		fmt.Println("Numero nao é maior que 0")
	}

	//fmt.Println(outroNumero)

	numeroNovo := 20

	if numeroNovo > 20 {
		fmt.Println("numero maior que 20")
	} else if numeroNovo < 20 {
		fmt.Println("numero menor que 20")
	} else {
		fmt.Println("numero igual a 20")
	}

}
