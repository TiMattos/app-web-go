package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/TiMattos/app-web-go/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//o FormValue obtem o valor através do nome da propriedade que foi inserir no form html, no caso new.html
		//lá, o nome da propriedade que vai receber o nome do produto foi definido como "nome"
		//o mesmo acontece para todos os outros controles

		//<input type="text" name="nome" class="form-control">
		nome := r.FormValue("nome")
		//<input type="text" name="descricao" class="form-control">
		descricao := r.FormValue("descricao")
		// <input type="number" name="preco" class="form-control" step="0.01">
		//mesmo no html estando como number, o request chega para o back como string, neste caso é necessário converver
		//para o formato da propriedade, no caso abaixo estou fazendo direto na mesma linha, porem é para teste
		//o correto seria criar uma variavel, converter o preco e depois verificar se não houve erro na conversão, para, ai sim
		//jogar o valor da variavel nova para a varivel preco
		preco, _ := strconv.ParseFloat(r.FormValue("preco"), 64)
		//<input type="number" name="quantidade" class="form-control">
		//como a quantidade tambem chega como string e deve ser convertido apra int, farei agora da forma correta
		quantidade := r.FormValue("quantidade")

		quantidadeConvertida, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Houve um erro ao converter a quantidade")
		}

		models.CriarNovoProduto(nome, descricao, preco, quantidadeConvertida)
		http.Redirect(w, r, "/", 301)

	}
}
