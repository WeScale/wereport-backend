package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"log"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func GetContrats(w http.ResponseWriter, r *http.Request) {
	var Contrats Contrats
	Contrats = RepoContrats()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(Contrats); err != nil {
		panic(err)
	}
}

func GetOneContrat(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	contratID, err := gocql.ParseUUID(vars["id"])
	if err != nil {
		panic(err)
	}
	var clt Contrat
	clt = RepoFindContrat(contratID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(clt); err != nil {
		panic(err)
	}
}

func ContratCreate(w http.ResponseWriter, r *http.Request) {
	var clt Contrat
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &clt); err != nil {
		log.Println("error mapping")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	} else {
		var status int
		contrat := RepoCreateContrat(clt)

		if contrat == (Contrat{}) {
			status = http.StatusBadRequest
		} else {
			status = http.StatusCreated
		}

		log.Println(contrat)
		contratdata := MarshalHateoas(contrat)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(status)
		if err := json.NewEncoder(w).Encode(contratdata); err != nil {
			panic(err)
		}
	}
}
