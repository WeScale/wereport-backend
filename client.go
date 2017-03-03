package main

import (
	"fmt"
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

// CREATE TABLE IF NOT EXISTS
// we.client(
// 	ID UUID,
// 	Name text,
// 	Service text,
// 	Creation timestamp, PRIMARY KEY(id))

func RepoClients(cluster *gocql.ClusterConfig) Clients {

	var client Client
	var list Clients

	session, _ := cluster.CreateSession()
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
func RepoFindClient(cluster *gocql.ClusterConfig, id gocql.UUID) Client {

	var client Client
	session, _ := cluster.CreateSession()
	if err := session.Query(`SELECT ID, Name, Service, Creation FROM client WHERE id = ? `, id).Consistency(gocql.One).Scan(&client.ID, &client.Name, &client.Service, &client.Creation); err != nil {
		log.Fatal(err)
	}

	// return empty Todo if not found
	return client
}

//RepoCreateClient create client
func RepoCreateClient(cluster *gocql.ClusterConfig, t Client) Client {

	session, err := cluster.CreateSession()

	if err != nil {
		log.Fatal(err)
	}

	if err := session.Query(`INSERT INTO client (ID, Name, Service, Creation) VALUES (?, ?, ?, ?)`,
		gocql.TimeUUID(), t.Name, t.Service, time.Now()).Exec(); err != nil {
		log.Fatal(err)
	}

	return t
}

func RepoDestroyClient(cluster *gocql.ClusterConfig, id gocql.UUID) error {

	//Todo

	return fmt.Errorf("Could not find Client with id of %d to delete", id)
}
