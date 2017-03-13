#!/bin/bash

docker run --name some-cassandra -d cassandra

docker run -p 8080:8080 --link some-cassandra -it test-go

curl -H "Content-Type: application/json" -d '{"Name":"SG", "Service":"RET/API"}' http://localhost:8080/clients
curl -H "Content-Type: application/json" -d '{"Name":"SG", "Service":"MKT"}' http://localhost:8080/clients

curl -H "Content-Type: application/json" -d '{"FirstName":"Sébastien", "LastName":"Lavayssière"}' http://localhost:8080/consultants
curl -H "Content-Type: application/json" -d '{"FirstName":"Cédric", "LastName":"Hauber"}' http://localhost:8080/consultants

curl -H "Content-Type: application/json" http://localhost:8080/clients
curl -H "Content-Type: application/json" http://localhost:8080/consultants
curl -H "Content-Type: application/json" http://localhost:8080/contrats

curl -H "Content-Type: application/json" -d '{"Consultant":"", "Client":"", "Tjm":100.1, "Bdc":"1234"}' http://localhost:8080/contrats
curl -H "Content-Type: application/json" -d '{"Consultant":"", "Client":"", "Tjm":200.2, "Bdc":"1234"}' http://localhost:8080/contrats

# curl -H "Content-Type: application/json" -d '{"Contrat":"", "Year":"2017", "Month":"03", "Day":"01", "Time":"1"}' http://localhost:8080/reportdays

# curl -H "Content-Type: application/json" http://localhost:8080/factures/client/-clientid-
# curl -H "Content-Type: application/json" http://localhost:8080/reports/consultant/-consultantid-

