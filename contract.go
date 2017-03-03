package main

import (
	"fmt"
	"log"

	"time"

	"github.com/gocql/gocql"
)

//Contrat ben un Contrat quoi
type Contrat struct {
	ID         gocql.UUID `json:"id"`
	Consultant gocql.UUID `json:"consultant"`
	Tjm        float32    `json:"tjm"`
	Bdc        string     `json:"bdc"`
	Debut      time.Time  `json:"debut"`
	Fin        time.Time  `json:"fin"`
}

//Contrats tous les Contrat
type Contrats []Contrat

// CREATE TABLE IF NOT EXISTS
// we.contract(ID UUID,
// Consultant UUID,
// Tjm int,
// Bdc text,
// Debut timestamp,
// Fin timestamp, PRIMARY KEY(id))

func RepoContrat(cluster *gocql.ClusterConfig) Contrats {

	var unique Contrat
	var list Contrats

	session, _ := cluster.CreateSession()
	iter := session.Query(`SELECT ID, Consultant, Tjm, Bdc, Debut, Fin FROM contrat`).Iter()
	for iter.Scan(&unique.ID, &unique.Consultant, &unique.Tjm, &unique.Bdc, &unique.Debut, &unique.Fin) {
		list = append(list, unique)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return list
}

//RepoFindContrat find one client
func RepoFindContrat(cluster *gocql.ClusterConfig, id gocql.UUID) Contrat {

	var unique Contrat
	session, _ := cluster.CreateSession()
	if err := session.Query(`SELECT ID, Consultant, Tjm, Bdc, Debut, Fin FROM contrat WHERE id = ? `,
		id).Consistency(gocql.One).Scan(&unique.ID, &unique.Consultant, &unique.Tjm, &unique.Bdc, &unique.Debut, &unique.Fin); err != nil {
		log.Fatal(err)
		return Contrat{}
	}

	// return empty Todo if not found
	return unique
}

//RepoCreateContrat create client
func RepoCreateContrat(cluster *gocql.ClusterConfig, t Contrat) Contrat {

	session, err := cluster.CreateSession()

	if err != nil {
		log.Fatal(err)
	}

	if err := session.Query(`INSERT INTO contrat (ID, Consultant, Tjm, Bdc, Debut, Fin) VALUES (?, ?, ?, ?, ?, ?)`,
		gocql.TimeUUID(), t.Consultant, t.Tjm, t.Bdc, t.Debut, t.Fin).Exec(); err != nil {
		log.Fatal(err)
	}

	return t
}

func RepoDestroyContrat(cluster *gocql.ClusterConfig, id gocql.UUID) error {

	//Todo

	return fmt.Errorf("Could not find Client with id of %d to delete", id)
}
