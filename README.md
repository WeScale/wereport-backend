# wereport-backend

## Install
```
docker run --name some-cassandra -d cassandra
```
then connect to cassandra and add keyspace
```
docker exec -it some-cassandra bash
cqlsh
create keyspace we with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
```

```
docker build -t wereport-backend .
docker run -p 8080:8080 \
    --link some-cassandra \
    --env CASSANDRA_HOSTS=172.17.0.2 \
    -it wereport-backend
```
## client management

### create
```
curl -H "Content-Type: application/json" -d '{"Name":"SG", "Service":"RET/API"}' http://localhost:8080/clients
curl -H "Content-Type: application/json" -d '{"Name":"SG", "Service":"MKT"}' http://localhost:8080/clients
```
### get
```
curl -H "Content-Type: application/json" http://localhost:8080/clients
```

## consultant management

### create
```
curl -H "Content-Type: application/json" -d '{"FirstName":"Sébastien", "LastName":"Lavayssière"}' http://localhost:8080/consultants
curl -H "Content-Type: application/json" -d '{"FirstName":"Cédric", "LastName":"Hauber"}' http://localhost:8080/consultants
```
### get
```
curl -H "Content-Type: application/json" http://localhost:8080/consultants
```