//CREATE TABLE IF NOT EXISTS we.ReportDayday(id UUID, client UUID, month int, day int, time float

package Data

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

//Client ben un client quoi
type ReportDay struct {
	ID        gocql.UUID `json:"id"`
	ContratID gocql.UUID `json:"contrat"`
	Report    gocql.UUID `json:"report"`
	Owner     gocql.UUID `json:"day_owner"`
	Day       int        `json:"day"`
	Time      float32    `json:"time"`
	Creation  time.Time  `json:"creation"`
	Contrat   Contrat    `json:"contrat_data"`
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

func (list ReportDays) RepoReportDays(report gocql.UUID) {

	var unique ReportDay

	iter := session.Query(`SELECT ID, Contrat, Report, Owner, Day, Time, Creation FROM ReportDay WHERE Report = ?  ALLOW FILTERING`, report).Iter()
	for iter.Scan(&unique.ID, &unique.ContratID, &unique.Report, &unique.Owner, &unique.Day, &unique.Time, &unique.Creation) {
		unique.Contrat.ID = unique.ContratID
		unique.Contrat.RepoFindContrat()
		list = append(list, unique)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}

//RepoFindReportDay
func (list ReportDays) RepoFindReportDay(report gocql.UUID, day int) {

	var unique ReportDay

	iter := session.Query(`SELECT ID, Contrat, Report, Owner, Day, Time, Creation FROM ReportDay WHERE Report = ? AND Day = ? ALLOW FILTERING`, report, day).Iter()
	for iter.Scan(&unique.ID, &unique.ContratID, &unique.Report, &unique.Owner, &unique.Day, &unique.Time, &unique.Creation) {
		unique.Contrat.ID = unique.ContratID
		unique.Contrat.RepoFindContrat()
		list = append(list, unique)
	}
	if err := iter.Close(); err != nil {
		log.Println(err)
	}
}

//RepoFindOneReportDay
func (unique ReportDay) RepoFindOneReportDay(report gocql.UUID, day int, contrat gocql.UUID) {

	if err := session.Query(`SELECT ID, Contrat, Report, Owner, Day, Time, Creation FROM ReportDay WHERE Report = ? AND Day = ? AND Contrat = ? ALLOW FILTERING`,
		report, day, contrat).Consistency(gocql.One).Scan(&unique.ID, &unique.ContratID, &unique.Report, &unique.Owner, &unique.Day, &unique.Time, &unique.Creation); err != nil {
		unique = ReportDay{}
	}

	unique.Contrat.ID = unique.ContratID
	unique.Contrat.RepoFindContrat()
}

//RepoCreateReportDay create client
func (unique ReportDay) RepoCreateReportDay() {

	var search ReportDays
	search.RepoFindReportDay(unique.Report, unique.Day)

	var totalTime float32
	if unique.Time != 0 {
		totalTime = 0
		for _, element := range search {
			totalTime = totalTime + element.Time
		}
	}

	if totalTime < 1 {
		var repDay ReportDay
		repDay.RepoFindOneReportDay(unique.Report, unique.Day, unique.ContratID)
		if repDay == (ReportDay{}) {
			unique.ID = gocql.TimeUUID()
			if err := session.Query(`INSERT INTO ReportDay (ID, Contrat, Report, Owner, Day, Time, Creation) VALUES (?, ?, ?, ?, ?, ?, ?)`,
				&unique.ID, &unique.ContratID, &unique.Report, &unique.Owner, &unique.Day, &unique.Time, &unique.Creation).Exec(); err != nil {
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
}

func (unique ReportDay) RepoDestroyReportDay() error {

	//Todo
	return fmt.Errorf("Could not find Client with id of %d to delete", unique.ID)
}
