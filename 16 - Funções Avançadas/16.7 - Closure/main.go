package main

import "fmt"

func Clousure() func() {
	texto := "Dentro da funcao closure"

	funcao := func() {
		fmt.Println(texto)
	}
	return funcao
}

func main() {
	fmt.Println("16.7 - Clousure")
	texto :="Dentro da funcao main"
	fmt.Println(texto)

	funcaoNova := Clousure()
	funcaoNova()

}
