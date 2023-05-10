package main

import "fmt"

type user struct {
	nome  string
	idade uint8
}

//dessa forma vinculamos o salvar a struct de user
//podendo ser acessados as propriedades através do vinculo
func (u user) salvar() {
	fmt.Printf("Salvando user: (%s - %d) no bando de dados\n", u.nome, u.idade)
}

func (u user) maiorDeIdate() bool {
	return u.idade >= 18
}

func (u *user) fazerAniversario() {
	u.idade++
}

func main() {
	user1 := user{nome: "Victor", idade: 28}
	fmt.Println(user1)
	user1.salvar()

	res1 := user1.maiorDeIdate()
	fmt.Println(res1)

	user1.fazerAniversario()
	fmt.Println(user1)

}
