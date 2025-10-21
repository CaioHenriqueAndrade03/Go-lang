package main

import "fmt"

func main() {
	fmt.Println("07 - Operadores")

	//artmeticos
	println("----------------------------")
	soma := 1 + 2
	subtracao := 2 - 1
	multiplicacao := 2 * 3
	divisao := 2.0 / 3
	restoDiv := 10 % 2
	fmt.Println(soma, subtracao, multiplicacao, divisao, restoDiv)

	//atribuiçao
	println("----------------------------")
	var variavel1 string = "variavel com tipo explicito (var variavel1 string)"
	variavel2 := "variavel declarada abreviação (variavel3 :=)"
	fmt.Println(variavel1, variavel2)

	//relacionais
	println("----------------------------")
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 < 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 != 2)

	//logicos
	println("----------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)

	//unarios
	println("----------------------------")
	numero := 10
	fmt.Println(numero)
	numero++
	fmt.Println(numero)

	//ternario
	println("----------------------------")

}
