package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {
	//HTML = hyper text markup language

	//aqui estamos jogando nossos htmls nos templates
	templates = template.Must(template.ParseGlob("22-HTML/*.html"))

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		user := usuario{"Victor", "victorsm9@hotmail.com"}

		//o terceiro parametro permiti mandar dados para o html
		templates.ExecuteTemplate(w, "home.html", user)
	})

	fmt.Println("Escutando a porta 5000")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
