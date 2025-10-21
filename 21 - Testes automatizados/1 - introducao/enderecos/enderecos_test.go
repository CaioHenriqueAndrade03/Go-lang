// Teste unitario
package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenariosDeTeste := []cenarioDeTeste{
		{"Avenida Paulista", "Avenida"},
		{"Rua Jorge Rudge", "Rua"},
		{"Rodovia São Luiz", "Rodovia"},
		{"Estrada do Urubu", "Estrada"},
		{" ", "Tipo invalido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEnderecos(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo de endereco recebido e o esperado estao diferentes, eu esperava %s e recebi %s", tipoDeEnderecoRecebido, cenario.retornoEsperado)
		}
	}

}

func TestQualquer(t *testing.T) {
	t.Parallel()

	if 1 > 2 {
		t.Errorf("Teste quebrou")
	}
}
