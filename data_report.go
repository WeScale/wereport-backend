package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

//Client ben un client quoi
type Report struct {
	ID        gocql.UUID `json:"id"`
	ReportDay gocql.UUID `json:"client"`
	Year      int        `json:"year"`
	Month     int        `json:"month"`
	Creation  time.Time  `json:"creation"`
}

//Clients tous les clients
type Reports []Report

//REATE TABLE IF NOT EXISTS we.report(id UUID, reportday UUID, month int

func init() {

	log.Printf("Create table report")
	if err := session.Query(`CREATE TABLE IF NOT EXISTS we.report(ID UUID, Reportday UUID, Month int, PRIMARY KEY(id))`).Exec(); err != nil {
		log.Println(err)
	}
}

func RepoReports() Reports {

	var unique Report
	var list Reports

	iter := session.Query(`SELECT ID, Reportday, Month FROM Report`).Iter()
	for iter.Scan(&unique.ID, &unique.ReportDay, &unique.Month) {
		list = append(list, unique)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return list
}

//RepoFindReport find one client
func RepoFindReport(id gocql.UUID) Report {

	var unique Report
	if err := session.Query(`SELECT ID, Reportday, Month FROM Report WHERE id = ? `,
		id).Consistency(gocql.One).Scan(&unique.ID, &unique.ReportDay, &unique.Month); err != nil {
		log.Println(err)
		return Report{}
	}

	// return empty Todo if not found
	return unique
}

//RepoCreateReport create client
func RepoCreateReport(unique Report) Report {

	unique.ID = gocql.TimeUUID()
	if err := session.Query(`INSERT INTO Report (ID, Reportday, Month) VALUES (?, ?, ?)`,
		&unique.ID, &unique.ReportDay, &unique.Month).Exec(); err != nil {
		log.Fatal(err)
	}

	return unique
}

func RepoDestroyReport(id gocql.UUID) error {

	//Todo

	return fmt.Errorf("Could not find Client with id of %d to delete", id)
}
