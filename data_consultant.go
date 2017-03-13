package main

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

//Consultant ben un consultant quoi
type Consultant struct {
	ID        gocql.UUID `json:"id"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
}

//Consultants tous les consultant
type Consultants []Consultant

// CREATE TABLE IF NOT EXISTS
// we.consultant
// (id UUID,
// FirstName text,
// LastName text,
// PRIMARY KEY(id))

func init() {

	log.Printf("Create table consultant")
	if err := session.Query(`CREATE TABLE IF NOT EXISTS we.consultant(ID UUID, FirstName text, LastName text, PRIMARY KEY(id))`).Exec(); err != nil {
		log.Println(err)
	}
}

func RepoConsultants() Consultants {

	var unique Consultant
	var list Consultants

	iter := session.Query(`SELECT ID, FirstName, LastName FROM consultant`).Iter()
	for iter.Scan(&unique.ID, &unique.FirstName, &unique.LastName) {
		list = append(list, unique)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return list
}

//RepoFindConsultant find one client
func RepoFindConsultant(id gocql.UUID) Consultant {

	var unique Consultant
	if err := session.Query(`SELECT ID, FirstName, LastName FROM consultant WHERE ID = ? `,
		id).Consistency(gocql.One).Scan(&unique.ID, &unique.FirstName, &unique.LastName); err != nil {
		log.Println(err)
		return Consultant{}
	}

	// return empty Todo if not found
	return unique
}

//RepoCreateConsultant create client
func RepoCreateConsultant(unique Consultant) Consultant {

	unique.ID = gocql.TimeUUID()
	if err := session.Query(`INSERT INTO consultant (ID, FirstName, LastName) VALUES (?, ?, ?)`,
		unique.ID, unique.FirstName, unique.LastName).Exec(); err != nil {
		log.Fatal(err)
	}

	return unique
}

func RepoDestroyConsultant(id gocql.UUID) error {

	//Todo

	return fmt.Errorf("Could not find Client with id of %d to delete", id)
}
