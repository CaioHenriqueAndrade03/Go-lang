package main

import "fmt"

func main() {
	fmt.Println("04 - variaveis")

	var variavel1 string = "variavel com tipo explicito (var variavel1 string)"
	var variavel2 = "variavel com tipo implicito (var variavel2) "
	variavel3 := "variavel declarada abreviação (variavel3 :=)"
	var (
		variavel4 string = "declaraçao simultanea"
		variavel5 string = "declaraçao simultanea"
		variavel6 string = "declaraçao simultanea"
	)
	variavel7, variavel8 := "Variavel7", "Variavel8"

	fmt.Println("------------variaveis------------")
	fmt.Println(variavel1)
	fmt.Println(variavel2)
	fmt.Println(variavel3)
	fmt.Println(variavel4)
	fmt.Println(variavel5)
	fmt.Println(variavel6)
	fmt.Println(variavel7)
	fmt.Println(variavel8)
	fmt.Println("-------------constantes-----------")
	const constante1 string = "Constante1"
	const constante2 = "Constate2"
	fmt.Println(constante1)
	fmt.Println(constante2)
	fmt.Println("----------pre inversão--------------")
	variavel12 := "Oi eu sou o Caio"
	variavel11 := "Oi eu sou a Gabi"
	fmt.Println(variavel11)
	fmt.Println(variavel12)
	fmt.Println("----------pos inversão--------------")
	variavel11, variavel12 = variavel12, variavel11
	fmt.Println(variavel11)
	fmt.Println(variavel12)
}
