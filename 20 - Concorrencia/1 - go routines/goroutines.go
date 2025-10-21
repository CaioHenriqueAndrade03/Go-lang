package main

import (
	"fmt"
	"time"
)

//Concorrencia != Paralelismo

//Concorrencia, sao processos que estao sendo alternados no processador
//Paralelismo, são processos rodando em nucleos diferentes do processador simultaneamente

func main() {
	fmt.Println("20 - Go routines")
	go escrever("Olá mundo ") //GO routine -> metodo que torna essa chamada paralela
	escrever("Programando em GO!!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
