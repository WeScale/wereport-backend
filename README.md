# wereport-backend

## Install

docker run --name some-cassandra -d cassandra

docker build -t wereport-backend .
docker run -p 8080:8080 --link some-cassandra -it wereport-backend

## client management

### create

curl -H "Content-Type: application/json" -d '{"Name":"SG", "Service":"RET/API"}' http://localhost:8080/clients
curl -H "Content-Type: application/json" -d '{"Name":"SG", "Service":"MKT"}' http://localhost:8080/clients

### get

curl -H "Content-Type: application/json" http://localhost:8080/clients
