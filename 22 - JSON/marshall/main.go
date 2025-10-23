package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

//json.Marshal()->Map ou struct para json
//json.Marshal()->Json para map ou struct

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	//------------------------cachorro usando struct-----------------------------
	c := cachorro{"Rex", "Dalmata", 3}
	fmt.Println(c) //Criado a instancia da struct

	cachorroEmJSON, erro := json.Marshal(c) //Criado o json baseado na instancia
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorroEmJSON) //esse json esta em forma de slice de bytes([]bytes)

	fmt.Println(bytes.NewBuffer(cachorroEmJSON)) //conversor de bytes para linguagem
	//------------------------cachorro usando map------------------------------
	c2 := map[string]string{
		"nome": "Toby",
		"raca": "Poodle",
	}
	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}
	fmt.Println(cachorro2EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))

	
}
