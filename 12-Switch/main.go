package main

import "fmt"

func diaDaSemana(num int) string {
	switch num {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sabado"
	default:
		return "Numero inválido"
	}
}

func diaDaSemana2(num int) string {
	switch {
	case num == 1:
		return "Domingo"
	case num == 2:
		return "Segunda"
	case num == 3:
		return "Terça"
	case num == 4:
		return "Quarta"
	case num == 5:
		return "Quinta"
	case num == 6:
		return "Sexta"
	case num == 7:
		return "Sabado"
	default:
		return "Numero inválido"
	}
}

func diaDaSemana3(num int) string {
	var dia string

	switch {
	case num == 1:
		dia = "Domingo"
	case num == 2:
		dia = "Segunda"
	case num == 3:
		dia = "Terça"
		//essa flag joga o resultado para o proximo caso
		fallthrough
	case num == 4:
		dia = "Quarta"
	case num == 5:
		dia = "Quinta"
	case num == 6:
		dia = "Sexta"
	case num == 7:
		dia = "Sabado"
	default:
		dia = "Numero inválido"
	}

	return dia
}

func main() {
	res := diaDaSemana(1)
	fmt.Println(res)

	res = diaDaSemana(5)
	fmt.Println(res)

	res = diaDaSemana(9)
	fmt.Println(res)

	res = diaDaSemana2(2)
	fmt.Println(res)

	res = diaDaSemana2(6)
	fmt.Println(res)

	res = diaDaSemana2(8)
	fmt.Println(res)

	res = diaDaSemana3(3)
	fmt.Println(res)

	res = diaDaSemana3(7)
	fmt.Println(res)

	res = diaDaSemana3(0)
	fmt.Println(res)
}
