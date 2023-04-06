package main

import (
	"net/http"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul top", Preco: 29.90, Quantidade: 4},
		{"Notebook", "razoavel", 1229.90, 10},
		{"Fone", "Horrivel", 129.90, 2},
	}
	temp.ExecuteTemplate(w, "Index", produtos)
}
