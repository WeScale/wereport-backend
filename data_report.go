package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
	"github.com/rickar/cal"
)

//Client ben un mois quoi
type Report struct {
	ID         gocql.UUID  `json:"id"`
	Consultant gocql.UUID  `json:"consultant"`
	Year       int         `json:"year"`
	Month      int         `json:"month"`
	Creation   time.Time   `json:"creation"`
	Days       []ReportDay `json:"days"`
}

type ViewReport struct {
	ID       gocql.UUID `json:"id"`
	Contrat  Contrat    `json:"contrat"`
	ListDays []float32  `json:"list_days"`
}

type ViewReports []ViewReport
type Reports []Report

var calendar = createCalendar()

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

func createCalendar() *cal.Calendar {
	calendar := cal.NewCalendar()

	// add holidays for the business
	calendar.AddHoliday(cal.ECB_NewYearsDay)
	calendar.AddHoliday(cal.ECB_EasterMonday)
	calendar.AddHoliday(cal.DE_TagderArbeit)
	calendar.AddHoliday(cal.NewHoliday(time.May, 8))
	//jeudi de l'acension :/
	calendar.AddHoliday(cal.DE_Himmelfahrt)
	//lundi de pentecote :/
	calendar.AddHoliday(cal.DE_Pfingstmontag)

	calendar.AddHoliday(cal.NewHoliday(time.July, 14))
	calendar.AddHoliday(cal.NewHoliday(time.August, 15))
	calendar.AddHoliday(cal.NewHoliday(time.November, 1))
	calendar.AddHoliday(cal.NewHoliday(time.November, 11))
	calendar.AddHoliday(cal.US_Christmas)

	return calendar
}

func daysIn(m time.Month, year int) int {
	// This is equivalent to time.daysIn(m, year).
	return time.Date(year, m+1, 0, 0, 0, 0, 0, time.UTC).Day()
}

//ChangeDataType toto
func ChangeDataType(clt Report, consultantid gocql.UUID) ViewReports {
	var data ViewReports

	m := make(map[string]bool)
	for _, element := range clt.Days {
		if m[element.Contrat.String()] == false { // si c'est la premiere fois qu'on a le client, le client n'existe pas
			m[element.Contrat.String()] = true
			var listDay = make([]float32, daysIn(time.Month(clt.Month), clt.Year))
			for i := 0; i < daysIn(time.Month(clt.Month), clt.Year); i++ {
				listDay[i] = 0
			}
			listDay[element.Day-1] = element.Time
			data = append(data, ViewReport{ID: element.Report, Contrat: element.ContratData, ListDays: listDay})
		} else { //si le client existe
			for _, report := range data {
				if report.Contrat.ID == element.Contrat {
					report.ListDays[element.Day-1] = element.Time
				}
			}
		}
	}
	contrats := RepoContratsOneConsultant(consultantid)
	for _, contrat := range contrats {
		var test bool
		test = false
		for _, report := range data {
			if report.Contrat.ID == contrat.ID {
				test = true
			}
		}
		if test == false {
			var listDay = make([]float32, daysIn(time.Month(clt.Month), clt.Year))
			for i := 0; i < daysIn(time.Month(clt.Month), clt.Year); i++ {
				listDay[i] = 0
			}
			data = append(data, ViewReport{ID: clt.ID, Contrat: contrat, ListDays: listDay})
		}
	}
	//add weekend
	var listDay = make([]float32, daysIn(time.Month(clt.Month), clt.Year))
	for i := 0; i < daysIn(time.Month(clt.Month), clt.Year); i++ {
		location, _ := time.LoadLocation("Europe/Paris")
		testdate := time.Date(clt.Year, time.Month(clt.Month), (i + 1), 0, 0, 0, 0, location)
		if calendar.IsWorkday(testdate) {
			listDay[i] = 0
		} else {
			listDay[i] = 1
		}
	}
	data = append(data, ViewReport{ID: gocql.TimeUUID(), Contrat: Contrat{Name: "nonworkday"}, ListDays: listDay})

	return data

}

func RepoReports(year int, month int) Reports {

	var unique Report
	var list Reports

	iter := session.Query(`SELECT ID, Consultant, Year, Month FROM Report WHERE Year = ? AND Month = ? ALLOW FILTERING`, year, month).Iter()
	for iter.Scan(&unique.ID, &unique.Consultant, &unique.Year, &unique.Month) {
		unique.Days = RepoReportDays(unique.ID)
		list = append(list, unique)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return list
}

//RepoFindReport find one client
func RepoFindReport(year int, month int, consultantid gocql.UUID) Report {

	var unique Report

	if err := session.Query(`SELECT ID, Consultant, Year, Month FROM Report WHERE Year = ? AND Month = ? AND Consultant = ? ALLOW FILTERING`,
		year, month, consultantid).Consistency(gocql.One).Scan(&unique.ID, &unique.Consultant, &unique.Year, &unique.Month); err != nil {
		unique = CreateEmptyReport(year, month, consultantid)
	}
	unique.Days = RepoReportDays(unique.ID)
	return unique
}

//RepoCreateReport create client
func RepoCreateReport(unique Report) Report {

	unique.ID = gocql.TimeUUID()
	if err := session.Query(`INSERT INTO Report (ID, Consultant, Year, Month) VALUES (?, ?, ?, ?)`,
		&unique.ID, &unique.Consultant, &unique.Year, &unique.Month).Exec(); err != nil {
		log.Fatal(err)
	}

	return unique
}

func RepoDestroyReport(id gocql.UUID) error {

	//Todo

	return fmt.Errorf("Could not find Client with id of %d to delete", id)
}

func CreateEmptyReport(year int, month int, consultantid gocql.UUID) Report {
	var ret Report

	ret.Consultant = consultantid
	ret.Creation = time.Now()
	ret.Month = month
	ret.Year = year

	ret = RepoCreateReport(ret)

	var contrats Contrats
	contrats = RepoContratsOneConsultant(consultantid)

	for i, element := range contrats {
		var retDay ReportDay

		retDay.Day = (i + 1)
		retDay.Report = ret.ID
		retDay.Time = 0
		retDay.Contrat = element.ID
		retDay.ContratData = element
		RepoCreateReportDay(retDay)
	}

	return ret
}
