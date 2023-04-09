package main

import (
	"net/http"

	"github.com/TiMattos/app-web-go/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
