package main

import (
	"encoding/json"
	"fmt"
	"html"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index, %q", html.EscapeString(r.URL.Path))
}

func GetClients(w http.ResponseWriter, r *http.Request) {
	clients = RepoClients(cluster)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(clients); err != nil {
		panic(err)
	}
}

func GetOneClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientID, err := gocql.ParseUUID(vars["clientid"])
	if err != nil {
		panic(err)
	}
	var clt Client
	clt = RepoFindClient(cluster, clientID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(clt); err != nil {
		panic(err)
	}
}

func ClientCreate(w http.ResponseWriter, r *http.Request) {
	var client Client
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &client); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	t := RepoCreateClient(cluster, client)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}