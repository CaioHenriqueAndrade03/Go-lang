package main

import "fmt"

var n int

func main() {
	fmt.Println("16.9 - Init")
	fmt.Println("Funcao main sendo executada")
	fmt.Println(n)
}

func init() {
	fmt.Println("Executando a funcao init")
	n = 10
}
