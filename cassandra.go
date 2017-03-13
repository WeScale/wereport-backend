package main

import (
	"log"
	"os"

	"github.com/gocql/gocql"
)

//NewCassandra Give us some seed data
func NewCassandra() (*gocql.ClusterConfig, *gocql.Session) {
	cluster := gocql.NewCluster(os.Getenv("CASSANDRA_HOSTS"))
	cluster.Consistency = gocql.Quorum

	log.Printf("Connexion to " + cluster.Hosts[0])

	session, err := cluster.CreateSession()

	log.Printf("Connexion to " + cluster.Hosts[0] + "... done")

	if err != nil {
		log.Fatal(err)
	}

	if err := session.Query(`CREATE KEYSPACE IF NOT EXISTS we WITH REPLICATION = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 }`).Exec(); err != nil {
		log.Fatal(err)
	}

	cluster.Keyspace = "we"
	sessionfinal, errfinal := cluster.CreateSession()

	if errfinal != nil {
		log.Fatal(errfinal)
	}

	return cluster, sessionfinal
}

var cluster, session = NewCassandra()
