package main

import (
	"fmt"
	"log"

	"time"

	"github.com/gocql/gocql"
)

//Consultant ben un consultant quoi
type Consultant struct {
	ID        gocql.UUID `json:"id"`
	FirstName string     `json:"firstname"`
	LastName  string     `json:"lastname"`
	Email     string     `json:"email"`
	Profil    Profil     `json:"profile"`
}

type Profil int

const (
	CONSULTANT Profil = 1 + iota
	MANAGER
	DIRECTION
	ADMINISTRATOR
)

var profiles = [...]string{
	"CONSULTANT",
	"MANAGER",
	"DIRECTION",
	"ADMINISTRATOR",
}

// String() function will return the english name
// that we want out constant Profil be recognized as
func (profile Profil) String() string {
	return profiles[profile-1]
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
	if err := session.Query(`CREATE TABLE IF NOT EXISTS we.consultant(ID UUID, FirstName text, LastName text, Email text, Profil int, PRIMARY KEY(id))`).Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query(`CREATE INDEX IF NOT EXISTS index_Email ON we.consultant (Email)`).Exec(); err != nil {
		log.Println(err)
	}
}

func RepoConsultants() Consultants {

	var unique Consultant
	var list Consultants

	iter := session.Query(`SELECT ID, FirstName, LastName, Email, Profil FROM consultant`).Iter()
	for iter.Scan(&unique.ID, &unique.FirstName, &unique.LastName, &unique.Email, &unique.Profil) {
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
	if err := session.Query(`SELECT ID, FirstName, LastName, Email, Profil FROM consultant WHERE ID = ? `,
		id).Consistency(gocql.One).Scan(&unique.ID, &unique.FirstName, &unique.LastName, &unique.Email, &unique.Profil); err != nil {
		log.Println("find consultant", err, id)

		return Consultant{}
	}

	// return empty Todo if not found
	return unique
}

//RepoFindConsultantByEmail find one client
func RepoFindConsultantByEmail(email string) Consultant {

	var unique Consultant
	if err := session.Query(`SELECT ID, FirstName, LastName, Email, Profil FROM consultant WHERE Email = ? `,
		email).Consistency(gocql.One).Scan(&unique.ID, &unique.FirstName, &unique.LastName, &unique.Email, &unique.Profil); err != nil {
		log.Println(err)
		return Consultant{}
	}

	// return empty Todo if not found
	return unique
}

//RepoCreateConsultant create client
func RepoCreateConsultant(unique Consultant) Consultant {

	unique.ID = gocql.TimeUUID()

	switch unique.Email {
	case "celine.rochay@wescale.fr":
		unique.Profil = DIRECTION
	case "sebastien.lavayssiere@wescale.fr":
		unique.Profil = ADMINISTRATOR
	case "aurelien.maury@wescale.fr":
		unique.Profil = MANAGER
	default:
		unique.Profil = CONSULTANT
	}

	if err := session.Query(`INSERT INTO consultant (ID, FirstName, LastName, Email, Profil) VALUES (?, ?, ?, ?, ?)`,
		unique.ID, unique.FirstName, unique.LastName, unique.Email, unique.Profil).Exec(); err != nil {
		log.Fatal(err)
	}

	var wescale Client
	wescale = RepoFindClientByName("WeScale")

	RepoCreateContrat(Contrat{
		Name:         "RTT",
		Tjm:          0,
		Bdc:          "NA",
		Debut:        time.Now(),
		ClientID:     wescale.ID,
		ConsultantID: unique.ID,
	})

	RepoCreateContrat(Contrat{
		Name:         "CP",
		Tjm:          0,
		Bdc:          "NA",
		Debut:        time.Now(),
		ClientID:     wescale.ID,
		ConsultantID: unique.ID,
	})

	RepoCreateContrat(Contrat{
		Name:         "Absence",
		Tjm:          0,
		Bdc:          "NA",
		Debut:        time.Now(),
		ClientID:     wescale.ID,
		ConsultantID: unique.ID,
	})

	return unique
}

func RepoDestroyConsultant(id gocql.UUID) error {

	//Todo

	return fmt.Errorf("Could not find Client with id of %d to delete", id)
}
