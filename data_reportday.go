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
	ID       gocql.UUID `json:"id"`
	Contrat  gocql.UUID `json:"client"`
	Year     int        `json:"year"`
	Month    int        `json:"month"`
	Day      int        `json:"day"`
	Time     float32    `json:"time"`
	Creation time.Time  `json:"creation"`
}

//Clients tous les clients
type ReportDays []ReportDay

func init() {
	log.Printf("Create table ReportDay")
	if err := session.Query(`CREATE TABLE IF NOT EXISTS we.ReportDay(ID UUID, Contrat UUID, Month int, Day int, Time float, PRIMARY KEY(id))`).Exec(); err != nil {
		log.Println(err)
	}

}

func RepoReportDays() ReportDays {

	var unique ReportDay
	var list ReportDays

	iter := session.Query(`SELECT ID, Contrat, Month, Day, Time FROM ReportDay`).Iter()
	for iter.Scan(&unique.ID, &unique.Contrat, &unique.Month, &unique.Day, &unique.Time) {
		list = append(list, unique)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return list
}

//RepoFindReportDay find one client
func RepoFindReportDay(id gocql.UUID) ReportDay {

	var unique ReportDay
	if err := session.Query(`SELECT ID, Contrat, Month, Day, Time FROM ReportDay WHERE id = ? `,
		id).Consistency(gocql.One).Scan(&unique.ID, &unique.Contrat, &unique.Month, &unique.Day, &unique.Time); err != nil {
		log.Println(err)
		return ReportDay{}
	}

	// return empty Todo if not found
	return unique
}

//RepoCreateReportDay create client
func RepoCreateReportDay(unique ReportDay) ReportDay {

	unique.ID = gocql.TimeUUID()
	if err := session.Query(`INSERT INTO ReportDay (ID, Contrat, Month, Day, Time) VALUES (?, ?, ?, ?, ?)`,
		&unique.ID, &unique.Contrat, &unique.Month, &unique.Day, &unique.Time).Exec(); err != nil {
		log.Fatal(err)
	}

	return unique
}

func RepoDestroyReportDay(id gocql.UUID) error {

	//Todo
	return fmt.Errorf("Could not find Client with id of %d to delete", id)
}
