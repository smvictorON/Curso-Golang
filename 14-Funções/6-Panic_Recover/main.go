package main

import "fmt"

//nesse exemplo usamos o panic caso a nota nao possa ser 6
//com o panic ele chama todos os defer e depois termina
//sem prosseguir a apliação se nao tiver um recover
//diferente do erro que o programa segue
func alunoAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é 6!")
}

func recuperarExecucao() {
	//mesmo que seja executada sem a função entrar em panic
	//o resultado do recover vai ser nil e não vai executar o if
	if r := recover(); r != nil {
		fmt.Println("Recuperada!")
	}
}

func main() {
	fmt.Println(alunoAprovado(5, 5))
	fmt.Println("Terminou!")
}
