package main

import (
	"errors"
	"fmt"
)

func main() {
	//INT
	//int8, int16, int32, int64, int,
	//*** int sozinho usar a arquitetura do computador como base
	//o tipo int no go suporta numeros negativos, e o tipo uint u de unsigned nao aceita negativos
	var number8 int8 = -100
	var number16 int16 = -10000
	var number32 int32 = -1000000000
	var number64 int64 = -1000000000000000000
	var number int = -1000000000000000000
	numb := -1000000000000000000

	var unumber8 uint8 = 100
	var unumber16 uint16 = 10000
	var unumber32 uint32 = 1000000000
	var unumber64 uint64 = 1000000000000000000
	var unumber uint = 1000000000000000000

	fmt.Println("ints")
	fmt.Println(number8)
	fmt.Println(number16)
	fmt.Println(number32)
	fmt.Println(number64)
	fmt.Println(number)
	fmt.Println(numb)

	fmt.Println("uints")
	fmt.Println(unumber8)
	fmt.Println(unumber16)
	fmt.Println(unumber32)
	fmt.Println(unumber64)
	fmt.Println(unumber)

	//alias
	//podemos usar alias byte para uint8
	//podemos usar alias rune para int32
	var number_byte byte = 123
	var number_rune rune = 1234

	fmt.Println(number_byte)
	fmt.Println(number_rune)

	//FLOAT
	//float 32, float 64
	var flt32 float32 = 123456789012345678901234567890123456789.01
	var flt64 float64 = 123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789.01

	fmt.Println(flt32)
	fmt.Println(flt64)

	//STRING
	//go nao tem char, porém quando colocamos um caracter com aspas simples ele retorna o numero
	//do caracter referente a tabela asc
	var txt1 string = "Texto1"
	txt2 := "Texto2"
	char := 'V'

	fmt.Println(txt1)
	fmt.Println(txt2)
	fmt.Println(char)

	//BOOLEAN
	var bolltrue bool = true
	bollfalse := false

	fmt.Println(bolltrue)
	fmt.Println(bollfalse)

	//ERROR
	//no go o erro é um tipo de dados, precisamos usar o pacote nativo errors para definir um erro
	var err error = errors.New("Erro Interno")

	fmt.Println(err)

	//SEM VALOR
	var sem_texto string
	var sem_numero int
	var sem_bool bool
	var sem_erro error

	fmt.Println(sem_texto)
	fmt.Println(sem_numero)
	fmt.Println(sem_bool)
	fmt.Println(sem_erro)
}
