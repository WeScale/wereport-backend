package Data

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

//Client ben un mois quoi
type Report struct {
	ID         gocql.UUID `json:"id"`
	Consultant gocql.UUID `json:"consultant"`
	Year       int        `json:"year"`
	Month      int        `json:"month"`
	Creation   time.Time  `json:"creation"`
	Days       ReportDays `json:"days"`
}

type Reports []Report

//REATE TABLE IF NOT EXISTS we.report(id UUID, reportday UUID, month int
func init() {

	log.Printf("Create table report")
	if err := session.Query(`CREATE TABLE IF NOT EXISTS we.report(ID UUID, Consultant UUID, Year int, Month int, PRIMARY KEY(ID))`).Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query(`CREATE INDEX IF NOT EXISTS index_Report ON we.report (Year)`).Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query(`CREATE INDEX IF NOT EXISTS index_Report ON we.report (Month)`).Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query(`CREATE INDEX IF NOT EXISTS index_Report ON we.report (Consultant)`).Exec(); err != nil {
		log.Println(err)
	}
}

func (list Reports) RepoReports(year int, month int) {

	var unique Report

	iter := session.Query(`SELECT ID, Consultant, Year, Month FROM Report WHERE Year = ? AND Month = ? ALLOW FILTERING`, year, month).Iter()
	for iter.Scan(&unique.ID, &unique.Consultant, &unique.Year, &unique.Month) {
		unique.Days.RepoReportDays(unique.ID)
		list = append(list, unique)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}

//RepoFindReport find one client
func (unique Report) RepoFindReport(year int, month int, consultant Consultant) {

	if err := session.Query(`SELECT ID, Consultant, Year, Month FROM Report WHERE Year = ? AND Month = ? AND Consultant = ? ALLOW FILTERING`,
		year, month, consultant.ID).Consistency(gocql.One).Scan(&unique.ID, &unique.Consultant, &unique.Year, &unique.Month); err != nil {
		unique.CreateEmptyReport(year, month, consultant)
	}
	unique.Days.RepoReportDays(unique.ID)
}

//RepoCreateReport create client
func (unique Report) RepoCreateReport() {

	unique.ID = gocql.TimeUUID()
	if err := session.Query(`INSERT INTO Report (ID, Consultant, Year, Month) VALUES (?, ?, ?, ?)`,
		&unique.ID, &unique.Consultant, &unique.Year, &unique.Month).Exec(); err != nil {
		log.Fatal(err)
		unique = Report{}
	}
}

func (unique Report) RepoDestroyReport() error {

	//Todo

	return fmt.Errorf("Could not find Client with id of %d to delete", unique.ID)
}

func (unique Report) CreateEmptyReport(year int, month int, consultant Consultant) {

	unique.Consultant = consultant.ID
	unique.Creation = time.Now()
	unique.Month = month
	unique.Year = year

	unique.RepoCreateReport()

	var contrats Contrats
	contrats = consultant.RepoContrats()

	for i, element := range contrats {
		var retDay ReportDay

		retDay.Day = (i + 1)
		retDay.Report = unique.ID
		retDay.Time = 0
		retDay.ContratID = element.ID
		retDay.Contrat = element
		retDay.RepoCreateReportDay()
	}
}
