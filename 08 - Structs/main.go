package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}
type endereco struct {
	rua    string
	cep    int
	cidade string
}

func main() {
	fmt.Println("08 - Structs")
	//desse jeito, meio que mesclamos o primeiro e o segundo modo
	enderecoExemplo := endereco{rua: "alfredo ceschiatti", cep: 22775045, cidade: "Rj"}

	u := usuario{nome: "Caio", idade: 21, endereco: enderecoExemplo}
	fmt.Println(u)

	gabi := usuario{nome: "Gabi", idade: 21, endereco: enderecoExemplo}
	fmt.Println(gabi)
}
