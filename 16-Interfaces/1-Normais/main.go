package main

import (
	"fmt"
	"math"
)

//aqui implementamos uma interface que deve ter assinatura dos métodos
type forma interface {
	area() float64
}

type retangulo struct {
	altura  float64
	largura float64
}

type circulo struct {
	raio float64
}

//se fossemos fazer uma função para calcular a area de duas structs
//diferentes, como poderiamos colocar os parametros? teriamos que
//fazer duas funçoes uma que recebe circulo e outra retangulo
//por isso definimos uma interface em comum, porém o go permite
//implementemos uma interface sem dizer, ele só vai checar se existe
//um método com aquela assinatura pra struct que foi passada e executa
func escreverArea(f forma) {
	fmt.Printf("A forma da area é %0.2f\n", f.area())
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

//pow é usado para potencias, primeiro parametro é o numero e o segundo a potencia
func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {
	ret := retangulo{altura: 10, largura: 15}
	circ := circulo{raio: 10}

	escreverArea(ret)
	escreverArea(circ)
}
