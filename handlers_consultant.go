package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func GetConsultants(w http.ResponseWriter, r *http.Request) {
	var Consultants Consultants
	Consultants = RepoConsultants(cluster)
	w.Header().Set("Content-Typea", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(Consultants); err != nil {
		panic(err)
	}
}

func GetOneConsultant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ConsultantID, err := gocql.ParseUUID(vars["consultantid"])
	if err != nil {
		panic(err)
	}
	var clt Consultant
	clt = RepoFindConsultant(cluster, ConsultantID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(clt); err != nil {
		panic(err)
	}
}

func ConsultantCreate(w http.ResponseWriter, r *http.Request) {
	var clt Consultant
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &clt); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateConsultant(cluster, clt)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
