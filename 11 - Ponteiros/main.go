package main

import "fmt"

func main() {
	fmt.Println("11 - ponteiros")
	fmt.Println("============Passagem por copia===============")
	variavel1 := 10
	variavel2 := variavel1

	variavel1++

	fmt.Println(variavel1, variavel2)
	fmt.Println("============Passagem por ponteiro===============")
	//Ponteiro é uma referencia de memória

	var variavel3 int
	var ponteiro *int

	variavel3 = 11
	ponteiro = &variavel3
	variavel3++
	fmt.Println(variavel3, *ponteiro)//desreferenciação
}
