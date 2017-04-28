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
	wescale = RepoFindClientByName("WeScale")
	if wescale == (Client{}) {
		wescale = RepoCreateClient(Client{
			Name:    "WeScale",
			Service: "internal",
		})
	}
}

func RepoClients() Clients {

	var client Client
	var list Clients

	iter := session.Query(`SELECT ID, Name, Service, Creation FROM client`).Iter()
	for iter.Scan(&client.ID, &client.Name, &client.Service, &client.Creation) {
		list = append(list, client)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}

	return list
}

//RepoFindClient find one client
func RepoFindClient(id gocql.UUID) Client {

	var client Client
	if err := session.Query(`SELECT ID, Name, Service, Creation FROM client WHERE ID = ? `,
		id).Consistency(gocql.One).Scan(&client.ID, &client.Name, &client.Service, &client.Creation); err != nil {
		log.Println("not find client", err, id)
		return Client{}
	}

	// return empty Todo if not found
	return client
}

//RepoFindClient find one client
func RepoFindClientByName(name string) Client {

	var client Client
	if err := session.Query(`SELECT ID, Name, Service, Creation FROM client WHERE Name = ? `,
		name).Consistency(gocql.One).Scan(&client.ID, &client.Name, &client.Service, &client.Creation); err != nil {
		return Client{}
	}

	// return empty Todo if not found
	return client
}

//RepoCreateClient create client
func RepoCreateClient(t Client) Client {

	t.ID = gocql.TimeUUID()
	t.Creation = time.Now()

	if err := session.Query(`INSERT INTO client (ID, Name, Service, Creation) VALUES (?, ?, ?, ?)`,
		t.ID, t.Name, t.Service, t.Creation).Exec(); err != nil {
		log.Fatal(err)
	}

	return t
}

func RepoDestroyClient(id gocql.UUID) error {

	//DELETE FROM Persons WHERE familyname='BARON';
	return nil
}
