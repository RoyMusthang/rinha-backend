package api

import "github.com/gorilla/mux"

type Api struct {
	Router *mux.Router
}

func (a *Api) Initialize(host, user, password, dbname string) {
	a.Router = mux.NewRouter()
	a.initRoutes()
}

func (a *Api) initRoutes() {
	a.Router.HandleFunc("/pessoas", a.createPerson).Methods("POST")
	a.Router.HandleFunc("/pessoas", a.searchPerson).Methods("GET")
	a.Router.HandleFunc("/pessoas/{id}", a.getPerson).Methods("GET")
	a.Router.HandleFunc("/contagem-pessoas", a.countPerson).Methods("GET")
}
