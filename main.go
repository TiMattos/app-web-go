package main

import (
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaComBanco()
	retornoConsulta, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}
	produtos := []Produto{}

	for retornoConsulta.Next() {
		produto := Produto{}
		err = retornoConsulta.Scan(&produto.Id, &produto.Nome, &produto.Descricao, &produto.Preco, &produto.Quantidade)

		if err != nil {
			panic(err.Error())
		}

		produtos = append(produtos, produto)

	}

	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}
