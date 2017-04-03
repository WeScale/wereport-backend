package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func init() {
	routes = append(routes,
		Route{
			"ConsultantIndex",
			"GET",
			"/consultants",
			GetConsultants,
		},
		Route{
			"ConsultantShow",
			"GET",
			"/consultants/{id}",
			GetOneConsultant,
		},
		Route{
			"ConsultantCreate",
			"POST",
			"/consultants",
			ConsultantCreate,
		})
}

func GetConsultants(w http.ResponseWriter, r *http.Request) {
	var consultants Consultants
	consultants = RepoConsultants()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(consultants); err != nil {
		panic(err)
	}
}

func GetOneConsultant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	consultantID, err := gocql.ParseUUID(vars["id"])
	if err != nil {
		panic(err)
	}
	var clt Consultant
	clt = RepoFindConsultantByID(consultantID)
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

	t := RepoCreateConsultant(clt)
	ConsultantWebSocketSend(t)
	consultantdata := MarshalHateoas(t)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(consultantdata); err != nil {
		panic(err)
	}
}
