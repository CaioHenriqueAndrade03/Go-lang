package main

/*Para rodarmos o teste de uma função, vamos usar a função interna do GO, Testing
e para isso temos algumas regras:

1 - o teste de uma func, nao fica no mesmo arquivo que ela

2 - os arquivos de teste precisam ter um nome especifico: {nome do arquivo}_test.go

3 - a funcao de teste, sempre vai seguir esse molde aqui:
func Test{Nome da func que esta sendo testada com letra inicial maiuscula}(t *testing.T) {}

4 - Comando para roda os testes no terminal: go test, lembresse de estar no path correto -> esse comando faz apenas o basico que é retornar se passou ou não

5 -  go test ./... ele entra em todos os pacotes desde a raiz e roda os testes, caso apareça cached,
é por que a funcao nao foi alteradadesde a ultima vez que rodou

6 - go test -v , -v é o verbose e o codigo vai mostrar exatamente qual função que ele está rodando naquele momento

7 - go test --cover mostra como esta a cobertura de testes daquele arquivo

8 - go test --coverprofile {nome do arquivo que recebe o relatorio}.txt esse comando vai mostrar quais linhas estão sendo cobertas por testes

9 - go tool cover --func={ {nome do arquivo do relatorio}.txt, ele vai ler o relatorio e dizer quais funçoes e quantos % esta cobertos

10 - go tool cover --html={ {nome do arquivo do relatorio}.txt, ele vai mostrar quais linhas estao sendo testas ou não
de cada pacote(cada linha verdes, são linhas que foram executados enquanto o teste rodava, e vermelho sao linhas que nao rodaram e cinza são linhas nao trackeadas)

11 - os pacotes do arquivo _test, podem se chamar inves de enderecos ser enderecos_test, 
mas isso afeta que as funções do pacote endereco, agora precisam ser importadas e para evitarmos isso, podemos usar o alias
que é feito assim '. "introducao-testes/enderecos" ' pois ai dizemos ao codigo que esse era o principal pacote usado 
permitindo a instancia de funções diretamente

obs: os testes sao a unica coisa que pode ter 2 pacotes
diferentes dentro do mesmo arquivo
*/

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEnderecos("Rua alfredo ceschiatti")

	fmt.Println(tipoEndereco)
}
