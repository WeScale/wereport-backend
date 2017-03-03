package main

import (
	"log"
	"os"

	"github.com/gocql/gocql"
)

//NewCassandra Give us some seed data
func NewCassandra() *gocql.ClusterConfig {
	cluster := gocql.NewCluster(os.Getenv("CASSANDRA_HOSTS"))
	cluster.Consistency = gocql.Quorum
	cluster.Keyspace = "we"

	log.Printf("Connexion to " + cluster.Hosts[0])

	session, err := cluster.CreateSession()

	log.Printf("Connexion to " + cluster.Hosts[0] + "... done")

	if err != nil {
		log.Fatal(err)
	}

	// if err := session.Query(`CREATE KEYSPACE IF NOT EXISTS we WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 }`).Exec(); err != nil {
	// 	log.Fatal(err)
	// }

	log.Printf("Create table client")
	if err := session.Query(`CREATE TABLE IF NOT EXISTS we.client(ID UUID, Name text, Service text, Creation timestamp, PRIMARY KEY(id))`).Exec(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Create table contract")
	if err := session.Query(`CREATE TABLE IF NOT EXISTS we.contract(ID UUID, Consultant UUID, Tjm int, Bdc text, Debut timestamp, Fin timestamp, PRIMARY KEY(id))`).Exec(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Create table consultant")
	if err := session.Query(`CREATE TABLE IF NOT EXISTS we.consultant(ID UUID, FirstName text, LastName text, PRIMARY KEY(id))`).Exec(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Create table facture")
	if err := session.Query(`CREATE TABLE IF NOT EXISTS we.facture(id UUID, contract UUID, client UUID, days float, PRIMARY KEY(id))`).Exec(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Create table reportday")
	if err := session.Query(`CREATE TABLE IF NOT EXISTS we.reportday(id UUID, client UUID, month int, day int, time float, PRIMARY KEY(id))`).Exec(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Create table report")
	if err := session.Query(`CREATE TABLE IF NOT EXISTS we.report(id UUID, reportday UUID, month int, PRIMARY KEY(id))`).Exec(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Create table report...done")

	// if err := session.Query(`CREATE INDEX ON wereport.client(name)`).Exec(); err != nil {
	// 	log.Fatal(err)
	// }

	return cluster
}

var cluster = NewCassandra()
