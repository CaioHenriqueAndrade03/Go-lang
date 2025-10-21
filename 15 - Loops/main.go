package main

import (
	"fmt"
)

func main() {
	fmt.Println("15 - Loops")

	//i := 0
	// for i < 10 {
	// 	time.Sleep(time.Second)
	// 	i++
	// 	fmt.Println("Passei pelo i pela:", i, "vez")
	// }
	// fmt.Println("--------------------------------------------")
	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Passei pelo j pela:", j, "vez")
	// 	time.Sleep(time.Second)
	// }
	fmt.Println("--------------------------------------------")
	nomes := []string{"Caio", "João", "Gabi"}

	for indice, nome := range nomes {
		fmt.Println(nome, indice)
	}

	for indice, letra := range "Palavra" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"Nome":      "Caio",
		"Sobrenome": "Andrade",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {

		fmt.Println("Oi")
	}
}
