package main

import (
	"database/sql"
	"net/http"
	"text/template"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func conectaComBanco() *sql.DB {
	connStr := "user=app-web-go password=app-web-go dbname=root host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		panic(err.Error())
	}

	return db
}

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
