package main

// se o canal nao tiver buffer e estiver na mesma função ele vai dar problema, por que se o envio e o recebimento de dados estão no mesmo escopo,
// eles nao se reconhecem e acabam travados no envio dos dados causando um deadlock
import "fmt"

//canal sem buffer VS canal com buffer
// nesse caso a baixo, estou dizendo que o canal pode receber até 2 dados antes de encontrar o deadlock,
func main() {
	canal := make(chan string, 2) // aqui esta sendo definido a capacidade do buffer do canal, que nesse caso foi definido um buffer de 2

	canal <- "Olá mundo"
	canal <- "Me chamo Caio"

	mensagem := <-canal
	mensagem2 := <-canal 

	//O canal gera uma fila de execução, por isso o primeiro print sempre vai buscar a primeira menssagem que ele recebeu
	fmt.Println(mensagem)
	fmt.Println(mensagem2)

}
