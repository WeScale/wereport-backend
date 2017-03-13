package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func GetReports(w http.ResponseWriter, r *http.Request) {
	var Reports Reports
	Reports = RepoReports()
	w.Header().Set("Content-Typea", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(Reports); err != nil {
		panic(err)
	}
}

func GetOneReport(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reportID, err := gocql.ParseUUID(vars["id"])
	if err != nil {
		panic(err)
	}
	var clt Report
	clt = RepoFindReport(reportID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(clt); err != nil {
		panic(err)
	}
}

func ReportCreate(w http.ResponseWriter, r *http.Request) {
	var clt Report
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

	t := RepoCreateReport(clt)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
