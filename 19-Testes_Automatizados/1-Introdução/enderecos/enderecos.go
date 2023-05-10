package enderecos

import "strings"

func TipoDeEndereço(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	minuscula := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(minuscula, " ")[0]

	valido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavra {
			valido = true
		}
	}

	if valido {
		return strings.Title(primeiraPalavra)
	}

	return "Tipo Inválido"
}
