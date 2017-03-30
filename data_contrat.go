package main

import (
	"fmt"
	"log"

	"time"

	"github.com/gocql/gocql"
)

//Contrat ben un Contrat quoi
type Contrat struct {
	ID           gocql.UUID `json:"id"`
	Name         string     `json:"name"`
	Tjm          float32    `json:"tjm"`
	Bdc          string     `json:"bdc"`
	Debut        time.Time  `json:"debut"`
	Fin          time.Time  `json:"fin"`
	ClientID     gocql.UUID `json:"client_id"`
	ConsultantID gocql.UUID `json:"consultant_id"`
	Client       Client     `json:"client"`
	Consultant   Consultant `json:"consultant"`
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
	if err := session.Query(`CREATE TABLE IF NOT EXISTS we.contrat(ID UUID, Name text, Consultant UUID, Client UUID, Tjm float, Bdc text, Debut timestamp, Fin timestamp, PRIMARY KEY(id))`).Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query(`CREATE INDEX IF NOT EXISTS index_Consultant ON we.contrat (Consultant)`).Exec(); err != nil {
		log.Println(err)
	}
}

func RepoContrats() Contrats {

	var unique Contrat
	var list Contrats

	iter := session.Query(`SELECT ID, Name, Consultant, Client, Tjm, Bdc, Debut, Fin FROM contrat`).Iter()
	for iter.Scan(&unique.ID, &unique.Name, &unique.ConsultantID, &unique.ClientID, &unique.Tjm, &unique.Bdc, &unique.Debut, &unique.Fin) {
		unique.Client = RepoFindClient(unique.ClientID)
		unique.Consultant = RepoFindConsultantByID(unique.ConsultantID)
		list = append(list, unique)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return list
}

func RepoContratsOneConsultant(id gocql.UUID) Contrats {

	var unique Contrat
	var list Contrats

	iter := session.Query(`SELECT ID, Name, Consultant, Client, Tjm, Bdc, Debut, Fin FROM contrat WHERE Consultant = ?`, id).Iter()
	for iter.Scan(&unique.ID, &unique.Name, &unique.ConsultantID, &unique.ClientID, &unique.Tjm, &unique.Bdc, &unique.Debut, &unique.Fin) {
		unique.Client = RepoFindClient(unique.ClientID)
		unique.Consultant = RepoFindConsultantByID(unique.ConsultantID)
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
	if err := session.Query(`SELECT ID, Name, Consultant, Client, Tjm, Bdc, Debut, Fin FROM contrat WHERE id = ?`,
		id).Consistency(gocql.One).Scan(&unique.ID, &unique.Name, &unique.ConsultantID, &unique.ClientID, &unique.Tjm, &unique.Bdc, &unique.Debut, &unique.Fin); err != nil {
		return Contrat{}
	}

	unique.Client = RepoFindClient(unique.ClientID)
	unique.Consultant = RepoFindConsultantByID(unique.ConsultantID)

	// return empty Todo if not found
	return unique
}

//RepoCreateContrat create client
func RepoCreateContrat(unique Contrat) Contrat {

	client := RepoFindClient(unique.ClientID)
	if (Client{}) == client {
		return Contrat{}
	}
	unique.Client = client

	consultant := RepoFindConsultantByID(unique.ConsultantID)
	if (Consultant{}) == consultant {
		return Contrat{}
	}
	unique.Consultant = consultant

	unique.ID = gocql.TimeUUID()
	if err := session.Query(`INSERT INTO Contrat (ID, Name, Consultant, Client, Tjm, Bdc, Debut, Fin) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`,
		&unique.ID, &unique.Name, &unique.ConsultantID, &unique.ClientID, &unique.Tjm, &unique.Bdc, &unique.Debut, &unique.Fin).Exec(); err != nil {
		log.Println(err)
	}

	return unique
}

func RepoDestroyContrat(id gocql.UUID) error {

	//Todo

	return fmt.Errorf("Could not find Client with id of %d to delete", id)
}
