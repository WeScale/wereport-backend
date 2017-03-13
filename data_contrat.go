package main

import (
	"fmt"
	"log"

	"time"

	"github.com/gocql/gocql"
)

//Contrat ben un Contrat quoi
type Contrat struct {
	ID             gocql.UUID `json:"id"`
	Client         gocql.UUID `json:"client"`
	Consultant     gocql.UUID `json:"consultant"`
	Tjm            float32    `json:"tjm"`
	Bdc            string     `json:"bdc"`
	Debut          time.Time  `json:"debut"`
	Fin            time.Time  `json:"fin"`
	ClientData     Client     `json:"client_data"`
	ConsultantData Consultant `json:"consultant_data"`
}

//Contrats tous les Contrat
type Contrats []Contrat

// CREATE TABLE IF NOT EXISTS
// we.contract(ID UUID,
// Contrat UUID,
// Tjm int,
// Bdc text,
// Debut timestamp,
// Fin timestamp, PRIMARY KEY(id))

func init() {

	log.Printf("Create table contrat")
	if err := session.Query(`CREATE TABLE IF NOT EXISTS we.contrat(ID UUID, Consultant UUID, Client UUID, Tjm float, Bdc text, Debut timestamp, Fin timestamp, PRIMARY KEY(id))`).Exec(); err != nil {
		log.Println(err)
	}

}

func RepoContrats() Contrats {

	var unique Contrat
	var list Contrats

	iter := session.Query(`SELECT ID, Consultant, Client, Tjm, Bdc, Debut, Fin FROM contrat`).Iter()
	for iter.Scan(&unique.ID, &unique.Consultant, &unique.Client, &unique.Tjm, &unique.Bdc, &unique.Debut, &unique.Fin) {
		list = append(list, unique)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return list
}

//RepoFindContrat find one client
func RepoFindContrat(id gocql.UUID) Contrat {

	var unique Contrat
	if err := session.Query(`SELECT ID, Consultant, Client, Tjm, Bdc, Debut, Fin FROM contrat WHERE id = ? `,
		id).Consistency(gocql.One).Scan(&unique.ID, &unique.Consultant, &unique.Client, &unique.Tjm, &unique.Bdc, &unique.Debut, &unique.Fin); err != nil {
		log.Println(err)
		return Contrat{}
	}

	// return empty Todo if not found
	return unique
}

//RepoCreateContrat create client
func RepoCreateContrat(unique Contrat) Contrat {

	client := RepoFindClient(unique.Client)
	if (Client{}) == client {
		return Contrat{}
	}
	unique.ClientData = client

	consultant := RepoFindConsultant(unique.Consultant)
	if (Consultant{}) == consultant {
		return Contrat{}
	}
	unique.ConsultantData = consultant

	unique.ID = gocql.TimeUUID()
	if err := session.Query(`INSERT INTO Contrat (ID, Consultant, Client, Tjm, Bdc, Debut, Fin) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		&unique.ID, &unique.Consultant, &unique.Client, &unique.Tjm, &unique.Bdc, &unique.Debut, &unique.Fin).Exec(); err != nil {
		log.Println(err)
	}

	return unique
}

func RepoDestroyContrat(id gocql.UUID) error {

	//Todo

	return fmt.Errorf("Could not find Client with id of %d to delete", id)
}
