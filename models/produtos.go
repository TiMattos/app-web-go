package models

import "github.com/TiMattos/app-web-go/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func buscaTodosOsProdutos() []Produto {
	db := db.ConectaComBanco()
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
		defer db.Close()

	}
	return produtos
}
