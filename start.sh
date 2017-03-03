#!/bin/bash

# docker run --name some-cassandra -d cassandra
# docker exec -it some-cassandra bash
# cqlsh
# create keyspace example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
# create table example.client(id UUID, name text, completed boolean, due timestamp, PRIMARY KEY(id));
# create index on example.client(name);

docker build -t test-go .
docker run -p 8080:8080 --link some-cassandra -it test-go

curl -H "Content-Type: application/json" -d '{"Name":"SG", "Service":"RET/API"}' http://localhost:8080/clients
curl -H "Content-Type: application/json" -d '{"Name":"SG", "Service":"MKT"}' http://localhost:8080/clients

curl -H "Content-Type: application/json" http://localhost:8080/clients
curl -H "Content-Type: application/json" http://localhost:8080/clients/

