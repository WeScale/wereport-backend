package Data

import (
	"log"
	"time"

	"github.com/gocql/gocql"
)

//Client ben un client quoi
type Client struct {
	ID       gocql.UUID `json:"id"`
	Name     string     `json:"name"`
	Service  string     `json:"service"`
	Creation time.Time  `json:"creation"`
}

//Clients tous les clients
type Clients []Client

func init() {
	log.Printf("Create table client")
	if err := session.Query(`CREATE TABLE IF NOT EXISTS we.client(ID UUID, Name text, Service text, Creation timestamp, PRIMARY KEY(id))`).Exec(); err != nil {
		log.Println(err)
	}

	if err := session.Query(`CREATE INDEX IF NOT EXISTS index_Name ON we.client (Name)`).Exec(); err != nil {
		log.Println(err)
	}

	var wescale Client
	wescale.Name = "WeScale"
	wescale.RepoFindClient()
	if wescale == (Client{}) {
		wescale = Client{
			Name:    "WeScale",
			Service: "internal",
		}.RepoCreateClient()
	}
}

func (t Clients) RepoClients() {

	var client Client

	iter := session.Query(`SELECT ID, Name, Service, Creation FROM client`).Iter()
	for iter.Scan(&client.ID, &client.Name, &client.Service, &client.Creation) {
		t = append(t, client)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
}

//RepoFindClient find one client
func (t Client) RepoFindClient() {

	if len(t.Name) > 0 {
		if err := session.Query(`SELECT ID, Name, Service, Creation FROM client WHERE ID = ? `,
			t.ID).Consistency(gocql.One).Scan(t.ID, t.Name, t.Service, t.Creation); err != nil {
			log.Println("not find client", err, t.ID)
			t = Client{}
		}
	} else {
		if err := session.Query(`SELECT ID, Name, Service, Creation FROM client WHERE Name = ? `,
			t.Name).Consistency(gocql.One).Scan(t.ID, t.Name, t.Service, t.Creation); err != nil {
			log.Println("not find client", err, t.Name)
			t = Client{}
		}
	}
}

//RepoCreateClient create client
func (t Client) RepoCreateClient() Client {

	t.ID = gocql.TimeUUID()
	t.Creation = time.Now()

	if err := session.Query(`INSERT INTO client (ID, Name, Service, Creation) VALUES (?, ?, ?, ?)`,
		t.ID, t.Name, t.Service, t.Creation).Exec(); err != nil {
		log.Fatal(err)
	}

	return t
}

func (t Client) RepoDestroyClient() error {

	//DELETE FROM Persons WHERE familyname='BARON';
	return nil
}
