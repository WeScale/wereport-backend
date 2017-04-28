package Handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"log"

	"github.com/WeScale/wereport-backend/data"
	"github.com/WeScale/wereport-backend/websockets"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func GetContrats(w http.ResponseWriter, r *http.Request) {
	var contrats Data.Contrats
	contrats.RepoContrats()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(contrats); err != nil {
		panic(err)
	}
}

func GetOneContrat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contratID, err := gocql.ParseUUID(vars["id"])
	if err != nil {
		panic(err)
	}
	var clt Data.Contrat
	clt.ID = contratID
	clt.RepoFindContrat()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(clt); err != nil {
		panic(err)
	}
}

func GetContratsConsultant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	consultantID, err := gocql.ParseUUID(vars["id"])
	if err != nil {
		panic(err)
	}
	var contrats Data.Contrats
	var consultant Data.Consultant
	consultant.ID = consultantID
	consultant.RepoFindConsultant()
	contrats = consultant.RepoContrats()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(contrats); err != nil {
		panic(err)
	}
}

func ContratCreate(w http.ResponseWriter, r *http.Request) {
	var clt Data.Contrat
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &clt); err != nil {
		log.Println("error mapping", err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	} else {
		var status int
		contrat := clt.RepoCreateContrat()

		if contrat == (Data.Contrat{}) {
			status = http.StatusBadRequest
		} else {
			status = http.StatusCreated
			Websockets.ContratWebSocketSend(contrat)
		}

		contratdata := MarshalHateoas(contrat)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(status)
		if err := json.NewEncoder(w).Encode(contratdata); err != nil {
			panic(err)
		}
	}
}

//e89d05eb-10ab-11e7-b858-0242ac110003
