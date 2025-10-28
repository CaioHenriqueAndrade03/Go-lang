package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}

var templates *template.Template

type usuario struct {
	Nome      string
	Idade     uint
	Profissao string
}

func main() {
	u := usuario{
		"Caio",
		22,
		"Programador",
	}
	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/ping", ping)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templates.ExecuteTemplate(w, "home.html", u)
	})

	fmt.Println("Escutando na porta 5000")

	log.Fatal(http.ListenAndServe(":5000", nil))
}
