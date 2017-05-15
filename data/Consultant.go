package Data

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

func (list Consultants) RepoConsultants() {

	var unique Consultant

	iter := session.Query(`SELECT ID, FirstName, LastName, Email, Profil FROM consultant`).Iter()
	for iter.Scan(&unique.ID, &unique.FirstName, &unique.LastName, &unique.Email, &unique.Profil) {
		list = append(list, unique)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}

//RepoFindConsultantByID find one consultant
func (unique Consultant) RepoFindConsultant() {
	if len(unique.Email) > 0 {
		if err := session.Query(`SELECT ID, FirstName, LastName, Email, Profil FROM consultant WHERE Email = ? `,
			unique.Email).Consistency(gocql.One).Scan(&unique.ID, &unique.FirstName, &unique.LastName, &unique.Email, &unique.Profil); err != nil {
			log.Println("cannot find consultant", err, unique.Email)
			unique = Consultant{}
		}
	} else {
		if err := session.Query(`SELECT ID, FirstName, LastName, Email, Profil FROM consultant WHERE ID = ? `,
			unique.ID).Consistency(gocql.One).Scan(&unique.ID, &unique.FirstName, &unique.LastName, &unique.Email, &unique.Profil); err != nil {
			log.Println("cannot find consultant", err, unique.ID)
			unique = Consultant{}
		}
	}
}

//RepoCreateConsultant create client
func (unique Consultant) RepoCreateConsultant() {

	unique.ID = gocql.TimeUUID()

	switch unique.Email {
	case "slemesle@wescale.fr":
		unique.Profil = DIRECTION
	case "celine.rochay@wescale.fr":
		unique.Profil = DIRECTION
	case "sebastien.lavayssiere@wescale.fr":
		unique.Profil = ADMINISTRATOR
	case "aurelien.maury@wescale.fr":
		unique.Profil = MANAGER
	default:
		unique.Profil = CONSULTANT
	}

	log.Println("Profil add: ", unique.Profil)

	if err := session.Query(`INSERT INTO consultant (ID, FirstName, LastName, Email, Profil) VALUES (?, ?, ?, ?, ?)`,
		unique.ID, unique.FirstName, unique.LastName, unique.Email, unique.Profil).Exec(); err != nil {
		log.Fatal(err)
	}

	var wescale Client
	wescale.Name = "WeScale"
	wescale.RepoFindClient()

	Contrat{
		Name:         "RTT",
		Tjm:          0,
		Bdc:          "NA",
		Debut:        time.Now(),
		ClientID:     wescale.ID,
		ConsultantID: unique.ID,
	}.RepoCreateContrat()

	Contrat{
		Name:         "Congé Payé",
		Tjm:          0,
		Bdc:          "NA",
		Debut:        time.Now(),
		ClientID:     wescale.ID,
		ConsultantID: unique.ID,
	}.RepoCreateContrat()

	Contrat{
		Name:         "Congé Maladie",
		Tjm:          0,
		Bdc:          "NA",
		Debut:        time.Now(),
		ClientID:     wescale.ID,
		ConsultantID: unique.ID,
	}.RepoCreateContrat()

	Contrat{
		Name:         "Absence",
		Tjm:          0,
		Bdc:          "NA",
		Debut:        time.Now(),
		ClientID:     wescale.ID,
		ConsultantID: unique.ID,
	}.RepoCreateContrat()

	Contrat{
		Name:         "WeShare & Conf",
		Tjm:          0,
		Bdc:          "NA",
		Debut:        time.Now(),
		ClientID:     wescale.ID,
		ConsultantID: unique.ID,
	}.RepoCreateContrat()

	Contrat{
		Name:         "Formation",
		Tjm:          0,
		Bdc:          "NA",
		Debut:        time.Now(),
		ClientID:     wescale.ID,
		ConsultantID: unique.ID,
	}.RepoCreateContrat()

	Contrat{
		Name:         "Intercontrat",
		Tjm:          0,
		Bdc:          "NA",
		Debut:        time.Now(),
		ClientID:     wescale.ID,
		ConsultantID: unique.ID,
	}.RepoCreateContrat()
}

func (unique Consultant) RepoDestroyConsultant() error {

	//Todo

	return fmt.Errorf("Could not find Client with id of %d to delete", unique.ID)
}
