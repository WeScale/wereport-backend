package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

//Client ben un client quoi
type Facture struct {
	ID       gocql.UUID `json:"id"`
	Contract gocql.UUID `json:"contract"`
	Client   gocql.UUID `json:"client"`
	Days     int        `json:"days"`
	Creation time.Time  `json:"creation"`
}

//Clients tous les clients
type Factures []Facture

//facture(ID UUID, Contrat UUID, client UUID, Days float, Creation

func init() {

	log.Printf("Create table facture")
	if err := session.Query(`CREATE TABLE IF NOT EXISTS we.facture(ID UUID, Contract UUID, Client UUID, Days float, PRIMARY KEY(id))`).Exec(); err != nil {
		log.Println(err)
	}
}

func RepoFactures() Factures {

	var unique Facture
	var list Factures

	iter := session.Query(`SELECT ID, Contract, Client, Days, Creation FROM facture`).Iter()
	for iter.Scan(&unique.ID, &unique.Contract, &unique.Client, &unique.Days, &unique.Creation) {
		list = append(list, unique)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return list
}

//RepoFindFacture find one client
func RepoFindFacture(id gocql.UUID) Facture {

	var unique Facture
	if err := session.Query(`SELECT ID, Contract, Client, Days, Creation FROM facture WHERE id = ? `,
		id).Consistency(gocql.One).Scan(&unique.ID, &unique.Contract, &unique.Client, &unique.Days, &unique.Creation); err != nil {
		log.Println(err)
		return Facture{}
	}

	// return empty Todo if not found
	return unique
}

//RepoCreateFacture create client
func RepoCreateFacture(unique Facture) Facture {

	unique.ID = gocql.TimeUUID()
	if err := session.Query(`INSERT INTO Facture (ID, Contract, Client, Days, Creation) VALUES (?, ?, ?, ?, ?)`,
		&unique.ID, &unique.Contract, &unique.Client, &unique.Days, &unique.Creation).Exec(); err != nil {
		log.Fatal(err)
	}

	return unique
}

func RepoDestroyFacture(id gocql.UUID) error {

	//Todo

	return fmt.Errorf("Could not find Client with id of %d to delete", id)
}
