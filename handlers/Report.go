package Handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/WeScale/wereport-backend/data"
	"github.com/WeScale/wereport-backend/websockets"
	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func GetReports(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year, err := strconv.Atoi(vars["year"])
	month, err := strconv.Atoi(vars["month"])
	if err != nil {
		panic(err)
	}
	var clt Data.Reports
	clt = Data.RepoReports(year, month)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(clt); err != nil {
		panic(err)
	}
}

func GetReportsOneConsultant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	year, err := strconv.Atoi(vars["year"])
	month, err := strconv.Atoi(vars["month"])
	consultantid, err := gocql.ParseUUID(vars["id"])
	if err != nil {
		panic(err)
	}
	var clt Data.Report
	clt = Data.RepoFindReport(year, month, consultantid)

	var cltData Data.ViewReports
	cltData = Data.ChangeDataType(clt, consultantid)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(cltData); err != nil {
		panic(err)
	}
}

func ReportCreate(w http.ResponseWriter, r *http.Request) {
	var clt Data.Report
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

	t := Data.RepoCreateReport(clt)
	Websockets.ReportWebSocketSend(clt)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}
