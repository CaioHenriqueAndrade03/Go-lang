package main

import "fmt"

func main() {
	fmt.Println("16.3 - Funções anonimas")
	retorno := func(texto string) string{
		return fmt.Sprintf("Recebido -> %s",texto)
	}("Passando parametro")
	fmt.Println(retorno)
}