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
	Email     string     `json:"email"`
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
	if err := session.Query(`CREATE TABLE IF NOT EXISTS we.consultant(ID UUID, FirstName text, LastName text, Email text, PRIMARY KEY(id))`).Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query(`CREATE INDEX IF NOT EXISTS index_Email ON we.consultant (Email)`).Exec(); err != nil {
		log.Println(err)
	}
}

func RepoConsultants() Consultants {

	var unique Consultant
	var list Consultants

	iter := session.Query(`SELECT ID, FirstName, LastName, Email FROM consultant`).Iter()
	for iter.Scan(&unique.ID, &unique.FirstName, &unique.LastName, &unique.Email) {
		list = append(list, unique)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	return list
}

//RepoFindConsultantByID find one client
func RepoFindConsultantByID(id gocql.UUID) Consultant {

	var unique Consultant
	if err := session.Query(`SELECT ID, FirstName, LastName, Email FROM consultant WHERE ID = ? `,
		id).Consistency(gocql.One).Scan(&unique.ID, &unique.FirstName, &unique.LastName, &unique.Email); err != nil {
		log.Println(err)
		return Consultant{}
	}

	// return empty Todo if not found
	return unique
}

//RepoFindConsultantByEmail find one client
func RepoFindConsultantByEmail(email string) Consultant {

	var unique Consultant
	if err := session.Query(`SELECT ID, FirstName, LastName, Email FROM consultant WHERE Email = ? `,
		email).Consistency(gocql.One).Scan(&unique.ID, &unique.FirstName, &unique.LastName, &unique.Email); err != nil {
		log.Println(err)
		return Consultant{}
	}

	// return empty Todo if not found
	return unique
}

//RepoCreateConsultant create client
func RepoCreateConsultant(unique Consultant) Consultant {

	unique.ID = gocql.TimeUUID()
	if err := session.Query(`INSERT INTO consultant (ID, FirstName, LastName, Email) VALUES (?, ?, ?, ?)`,
		unique.ID, unique.FirstName, unique.LastName, unique.Email).Exec(); err != nil {
		log.Fatal(err)
	}

	return unique
}

func RepoDestroyConsultant(id gocql.UUID) error {

	//Todo

	return fmt.Errorf("Could not find Client with id of %d to delete", id)
}
