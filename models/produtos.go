package models

import "github.com/TiMattos/app-web-go/db"

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
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
func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaComBanco()
	insereNoBancoDeDados, err := db.Prepare("Insert into produtos(nome, descricao, preco, quantidade) values ($1,$2,$3,$4)")

	if err != nil {
		panic(err.Error())
	}
	insereNoBancoDeDados.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}
