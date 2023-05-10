//TESTE UNITARIO
package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes ", "Rodovia"},
		{"Estrada dos Negantes", "Estrada"},
		{"RUA ABC", "Rua"},
		{"AVENIDA Paulista", "Avenida"},
		{"RODOVIA dos Imigrantes ", "Rodovia"},
		{"ESTRADA dos Negantes", "Estrada"},
		{"rua ABC", "Rua"},
		{"avenida Paulista", "Avenida"},
		{"rodovia dos Imigrantes ", "Rodovia"},
		// {"estrada dos Negantes", "Estrada"},
		// {"Praça das Rosas", "Tipo Inválido"},
		// {"", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereço(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("Tipo recebido(%s) é diferente do esperado(%s)",
				retornoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

}
