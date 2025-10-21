package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("05 - tipos de dados")
	//int8   suporta numeros ate 8 bits
	//int16 suporta numeros ate 16 bits
	//int32 suporta numeros ate 32 bits
	//int64 suporta numeros ate 64 bits

	var numero16bits int16 = 1000
	var numero64bits int = -100
	var unumero64bits uint = 100
	var numero rune = 100

	var PrimeiroNumero float32 = 123.456
	var SegundoNumero float64 = 789.101112

	fmt.Println(numero16bits)
	fmt.Println(numero64bits)
	fmt.Println(unumero64bits)
	fmt.Println(numero)
	fmt.Println(PrimeiroNumero)
	fmt.Println(SegundoNumero)

	var variavel1 bool
	fmt.Println(variavel1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}
