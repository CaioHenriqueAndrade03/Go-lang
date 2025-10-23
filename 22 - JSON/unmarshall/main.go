package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"nome"` //caso a tag aqui e no json estejam diferentes,
	// ele perde a referencia e nao consegue linkar os dados

	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	//-----------------------JSON -> STRUCT------------------
	cachorroEmJSON := `{"nome":"Rex","raca":"Dalmata","idade":3}`
	var c cachorro //criamos uma instancia vazia de cachorro

	//passamos o unmarshall convertendo o tipo
	//enviando o cachorroEmJSON como parametro e
	// pedindo para ele altera o endereço de memoria
	if erro := json.Unmarshal([]byte(cachorroEmJSON), &c); erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(c)
	//--------------------JSOn -> MAP--------------------------
	//Importante lembrar que o map tem dificuldades e converter tipos alem do
	//retorno esperado
	cachorro2EmJSON := `{"nome":"Toby","raca":"Poodle"}`
	c2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(cachorro2EmJSON), &c2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(c2)
}
