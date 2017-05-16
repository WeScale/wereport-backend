package Handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/WeScale/wereport-backend/data"
	"github.com/WeScale/wereport-backend/websockets"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func GetClients(w http.ResponseWriter, r *http.Request) {
	var clients Data.Clients
	clients.RepoClients()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(clients); err != nil {
		panic(err)
	}
}

func GetOneClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	clientID, err := gocql.ParseUUID(vars["id"])
	if err != nil {
		panic(err)
	}
	var clt Data.Client
	clt.ID = clientID
	clt.RepoFindClient()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(clt); err != nil {
		panic(err)
	}
}

func ClientCreate(w http.ResponseWriter, r *http.Request) {
	var client Data.Client
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

	client.RepoCreateClient()
	Websockets.ClientWebSocketSend(client)
	clientdata := MarshalHateoas(client)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(clientdata); err != nil {
		panic(err)
	}
}
