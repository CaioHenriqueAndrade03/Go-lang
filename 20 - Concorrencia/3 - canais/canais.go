package main

// quando escrevemos canal <- {variavel} dizemos que o canal esta recebendo um valor
// quando escrevemos <- canal, dizemos que o canal esta esperando um valor
// mensagem := <-canal   a variavel mensagem, esta recebendo o valor que chegou no canal

// em go, um canal pode estar aberto ou fechado, caso ele esteja aberto significa que ainda enviará dados, caso esteja fechado significa que ele ja acabou
import (
	"fmt"
	"time"
)

func escrever(texto string, canal chan string) {

	for i := 0; i < 5; i++ {
		canal <- texto //a variavel texto ta sendo enviada para o canal
		time.Sleep(time.Second)
	}
	close(canal) //fecha o canal para nao enviar nem receber mensagem
}

func main() {
	canal := make(chan string) // <- Como criamos o canal, seguimos o modelo make(chan {tipo do canal}) no caso, esse canal so envia e recebe strings
	go escrever("Olá mundo!", canal)
	fmt.Println("Depois da funcao escrever")

	for mesage := range canal {
		fmt.Println(mesage)
	} // nesse caso, como o canal é o range, nao precisamos da sintax da setinha por que quando ele fechar ele para

	// for {
	// 	mensagem, aberto := <-canal //aqui eu digo que o canal esta esperando um valor
	// 	if !aberto {
	// 		fmt.Println("fechei o canal")
	// 		break
	// 	}
	// 	fmt.Println(mensagem)
	// }
}
