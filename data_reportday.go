//CREATE TABLE IF NOT EXISTS we.ReportDayday(id UUID, client UUID, month int, day int, time float

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

//Client ben un client quoi
type ReportDay struct {
	ID          gocql.UUID `json:"id"`
	Contrat     gocql.UUID `json:"contrat"`
	Report      gocql.UUID `json:"report"`
	Owner       gocql.UUID `json:"day_owner"`
	Day         int        `json:"day"`
	Time        float32    `json:"time"`
	Creation    time.Time  `json:"creation"`
	ContratData Contrat    `json:"contrat_data"`
}

//Clients tous les clients
type ReportDays []ReportDay

func init() {
	log.Printf("Create table ReportDay")
	if err := session.Query(`CREATE TABLE IF NOT EXISTS we.ReportDay(ID UUID, Contrat UUID, Report UUID, Owner UUID, Day int, Time float, Creation timestamp, PRIMARY KEY(ID))`).Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query(`CREATE INDEX IF NOT EXISTS index_Report ON we.ReportDay (Report)`).Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query(`CREATE INDEX IF NOT EXISTS index_Report ON we.ReportDay (Day)`).Exec(); err != nil {
		log.Println(err)
	}

}

func RepoReportDays(report gocql.UUID) ReportDays {

	var unique ReportDay
	var list ReportDays

	iter := session.Query(`SELECT ID, Contrat, Report, Owner, Day, Time, Creation FROM ReportDay WHERE Report = ?  ALLOW FILTERING`, report).Iter()
	for iter.Scan(&unique.ID, &unique.Contrat, &unique.Report, &unique.Owner, &unique.Day, &unique.Time, &unique.Creation) {
		unique.ContratData = RepoFindContrat(unique.Contrat)
		list = append(list, unique)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	return list
}

//RepoFindReportDay
func RepoFindReportDay(report gocql.UUID, day int) ReportDays {

	var unique ReportDay
	var list ReportDays

	iter := session.Query(`SELECT ID, Contrat, Report, Owner, Day, Time, Creation FROM ReportDay WHERE Report = ? AND Day = ? ALLOW FILTERING`, report, day).Iter()
	for iter.Scan(&unique.ID, &unique.Contrat, &unique.Report, &unique.Owner, &unique.Day, &unique.Time, &unique.Creation) {
		unique.ContratData = RepoFindContrat(unique.Contrat)
		list = append(list, unique)
	}
	if err := iter.Close(); err != nil {
		log.Println(err)
	}

	return list
}

//RepoFindOneReportDay
func RepoFindOneReportDay(report gocql.UUID, day int, contrat gocql.UUID) ReportDay {
	var unique ReportDay

	if err := session.Query(`SELECT ID, Contrat, Report, Owner, Day, Time, Creation FROM ReportDay WHERE Report = ? AND Day = ? AND Contrat = ? ALLOW FILTERING`,
		report, day, contrat).Consistency(gocql.One).Scan(&unique.ID, &unique.Contrat, &unique.Report, &unique.Owner, &unique.Day, &unique.Time, &unique.Creation); err != nil {
		return ReportDay{}
	}

	unique.ContratData = RepoFindContrat(unique.Contrat)

	return unique
}

//RepoCreateReportDay create client
func RepoCreateReportDay(unique ReportDay) ReportDay {

	var search ReportDays
	search = RepoFindReportDay(unique.Report, unique.Day)

	var totalTime float32
	if unique.Time != 0 {
		totalTime = 0
		for _, element := range search {
			totalTime = totalTime + element.Time
		}
	}

	if totalTime < 1 {
		repDay := RepoFindOneReportDay(unique.Report, unique.Day, unique.Contrat)
		if repDay == (ReportDay{}) {
			unique.ID = gocql.TimeUUID()
			if err := session.Query(`INSERT INTO ReportDay (ID, Contrat, Report, Owner, Day, Time, Creation) VALUES (?, ?, ?, ?, ?, ?, ?)`,
				&unique.ID, &unique.Contrat, &unique.Report, &unique.Owner, &unique.Day, &unique.Time, &unique.Creation).Exec(); err != nil {
				log.Fatal(err)
			}
		} else {
			if err := session.Query(`UPDATE ReportDay SET Time = ? WHERE ID = ?`,
				&unique.Time, &repDay.ID).Exec(); err != nil {
				log.Fatal(err)
			}
		}
	} else {
		unique = ReportDay{}
	}

	return unique
}

func RepoDestroyReportDay(id gocql.UUID) error {

	//Todo
	return fmt.Errorf("Could not find Client with id of %d to delete", id)
}
