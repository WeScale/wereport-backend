package Data

import (
	"time"

	"github.com/gocql/gocql"
	"github.com/rickar/cal"
)

type ViewReport struct {
	ID       gocql.UUID `json:"id"`
	Contrat  Contrat    `json:"contrat"`
	ListDays []float32  `json:"list_days"`
}

type ViewReports []ViewReport

var calendar = createCalendar()

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
func ChangeDataType(clt Report, consultant Consultant) ViewReports {
	var data ViewReports

	m := make(map[string]bool)
	for _, element := range clt.Days {
		if m[element.ContratID.String()] == false { // si c'est la premiere fois qu'on a le client, le client n'existe pas
			m[element.ContratID.String()] = true
			var listDay = make([]float32, daysIn(time.Month(clt.Month), clt.Year))
			for i := 0; i < daysIn(time.Month(clt.Month), clt.Year); i++ {
				listDay[i] = 0
			}
			listDay[element.Day-1] = element.Time
			data = append(data, ViewReport{ID: element.Report, Contrat: element.Contrat, ListDays: listDay})
		} else { //si le client existe
			for _, report := range data {
				if report.Contrat.ID == element.ContratID {
					report.ListDays[element.Day-1] = element.Time
				}
			}
		}
	}
	contrats := consultant.RepoContrats()
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
