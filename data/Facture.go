package Data

import (
	"time"
)

//Facture ben un client quoi
type Facture struct {
	Contrat  Contrat   `json:"contrat"`
	Days     float32   `json:"days"`
	Cost     float32   `json:"cost"`
	Creation time.Time `json:"creation"`
	Bdc      string    `json:"bdc"`
}

//Factures tous les factures
type Factures []Facture

func init() {

	// log.Printf("Create table facture")
	// if err := session.Query(`CREATE TABLE IF NOT EXISTS we.facture(ID UUID, Contrat UUID, Client UUID, Days float, PRIMARY KEY(id))`).Exec(); err != nil {
	// 	log.Println(err)
	// }
}

//RepoCreateFacture create client
func RepoFindFactures(year int, month int) Factures {

	//get all reports for one month
	var reports Reports
	reports.RepoReports(year, month)

	mFacture := make(map[string]Facture)
	//for all report
	for _, report := range reports {
		for _, day := range report.Days {
			//si le contrat n'est pas dans la liste
			if mFacture[day.Contrat.Bdc] == (Facture{}) {
				mFacture[day.Contrat.Bdc] = Facture{
					Contrat:  day.Contrat,
					Days:     day.Time,
					Cost:     (day.Time * day.Contrat.Tjm),
					Creation: time.Now(),
					Bdc:      day.Contrat.Bdc,
				}

			} else { //si le contrat est dans la liste
				mFacture[day.Contrat.Bdc] = Facture{
					Contrat:  day.Contrat,
					Days:     (mFacture[day.Contrat.Bdc].Days + (day.Time)),
					Cost:     (mFacture[day.Contrat.Bdc].Cost + (day.Time * day.Contrat.Tjm)),
					Creation: time.Now(),
					Bdc:      day.Contrat.Bdc,
				}
			}
		}
	}

	var factures Factures
	for _, facture := range mFacture {
		factures = append(factures, facture)
	}

	return factures
}
