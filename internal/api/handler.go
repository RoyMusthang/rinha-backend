package api

import (
	"encoding/json"
	"net/http"

	"github.com/RoyMusthang/backend/internal/config"
	"github.com/RoyMusthang/backend/internal/db"
)

func (a *Api) createPerson(w http.ResponseWriter, r *http.Request) {
	var p db.Pessoa
	body := json.NewDecoder(r.Body)
	if err := body.Decode(&p); err != nil {
		w.WriteHeader(int(http.StatusBadRequest))
		return
	}
	defer r.Body.Close()
	if err := p.Validate(); err == false {
		w.WriteHeader(int(http.StatusUnprocessableEntity))
		return
	}
	if err := p.createPerson(a.DB); err != nil {
		respondWithError(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
}

func (a *Api) searchPerson(w http.ResponseWriter, r *http.Request) {

}

func (a *Api) getPerson(w http.ResponseWriter, r *http.Request) {

}

func (a *Api) countPerson(w http.ResponseWriter, r *http.Request) {

}
