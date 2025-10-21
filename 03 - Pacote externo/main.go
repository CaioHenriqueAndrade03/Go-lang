package main

import (
	"fmt"
	"pacote-externos/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("03 - pacotes externos")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("chrj@gmail.com")
	fmt.Println(erro)
}
