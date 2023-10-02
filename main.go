package main

import (
	"net/http"

	"github.com/RoyMusthang/backend/internal/api"
)

const (
	host     = "localhost"
	user     = "seuusuario"
	password = "suasenha"
	dbname   = "seudb"
)

func main() {
	api := api.Api{}
	api.Initialize(host, user, password, dbname)

	http.ListenAndServe(":8080", api.Router)
}
