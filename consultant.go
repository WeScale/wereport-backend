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

func RepoConsultant(cluster *gocql.ClusterConfig) Consultants {

	var unique Consultant
	var list Consultants

	session, _ := cluster.CreateSession()
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
func RepoFindConsultant(cluster *gocql.ClusterConfig, id gocql.UUID) Consultant {

	var unique Consultant
	session, _ := cluster.CreateSession()
	if err := session.Query(`SELECT ID, FirstName, LastName FROM consultant WHERE id = ? `,
		id).Consistency(gocql.One).Scan(&unique.ID, &unique.FirstName, &unique.LastName); err != nil {
		log.Fatal(err)
		return nil
	}

	// return empty Todo if not found
	return unique
}

//RepoCreateConsultant create client
func RepoCreateConsultant(cluster *gocql.ClusterConfig, t Consultant) Consultant {

	session, err := cluster.CreateSession()

	if err != nil {
		log.Fatal(err)
	}

	if err := session.Query(`INSERT INTO consultant (ID, FirstName, LastName) VALUES (?, ?, ?)`,
		gocql.TimeUUID(), t.FirstName, t.LastName).Exec(); err != nil {
		log.Fatal(err)
	}

	return t
}

func RepoDestroyConsultant(cluster *gocql.ClusterConfig, id gocql.UUID) error {

	//Todo

	return fmt.Errorf("Could not find Client with id of %d to delete", id)
}
