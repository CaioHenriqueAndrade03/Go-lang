package main

import (
	"fmt"
	"sync"
	"time"
)

//Concorrencia != Paralelismo

//Concorrencia, sao processos que estao sendo alternados no processador
//Paralelismo, são processos rodando em nucleos diferentes do processador simultaneamente

func main() {
	fmt.Println("20 - Go routines")
	var waitGroup sync.WaitGroup
	waitGroup.Add(2) // cria um contador que diz que o codigo tera que espera 2 go routines acabarem
	go func() {
		escrever("Olá mundo")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Programando em go!!")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // fala para ele esperar ate que o contador chegue a 0
}

func escrever(texto string) {

	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
