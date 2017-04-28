package Handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"log"

	"github.com/WeScale/wereport-backend/data"
	"github.com/gorilla/mux"
)

// func GetFactures(w http.ResponseWriter, r *http.Request) {
// 	var Factures Factures
// 	Factures = RepoFactures()
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)

// 	if err := json.NewEncoder(w).Encode(Factures); err != nil {
// 		panic(err)
// 	}
// }

func GetOneMonthFactures(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	month, err := strconv.Atoi(vars["month"])
	year, err := strconv.Atoi(vars["year"])
	if err != nil {
		panic(err)
	}
	log.Printf("facture for %d/%d", year, month)
	clt := Data.RepoFindFactures(year, month)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(clt); err != nil {
		panic(err)
	}
}
