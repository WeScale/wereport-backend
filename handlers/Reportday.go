package Handlers

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"log"

	"github.com/WeScale/wereport-backend/data"
	"github.com/gocql/gocql"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func GetReportDays(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var reportDays Data.ReportDays
	reportdayID, err := gocql.ParseUUID(vars["id"])
	if err != nil {
		panic(err)
	}

	reportDays.RepoReportDays(reportdayID)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(reportDays); err != nil {
		panic(err)
	}
}

func ReportDayCreate(w http.ResponseWriter, r *http.Request) {
	var clt Data.ReportDay
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &clt); err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	var consultant Data.Consultant
	consultant = context.Get(r, UserData).(Data.Consultant)
	clt.Owner = consultant.ID

	var statusHTTP int
	clt.RepoCreateReportDay()
	if clt == (Data.ReportDay{}) {
		statusHTTP = http.StatusConflict
	} else {
		statusHTTP = http.StatusCreated
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusHTTP)
	if err := json.NewEncoder(w).Encode(clt); err != nil {
		panic(err)
	}
}
