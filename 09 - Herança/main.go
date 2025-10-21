package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    string
}

type estudante struct {
	pessoa
	curso     string
	matricula int
}

func main() {
	fmt.Println("08 - Heranças")

	p := pessoa{
		nome:      "Caio",
		sobrenome: "Andrade",
		idade:     21,
		altura:    "183cm",
	}

	e := estudante{
		pessoa:    p,
		curso:     "Engenharia da computação",
		matricula: 1220108359,
	}

	fmt.Println(e)
}
