# wereport-backend

## Install
### Prerequisite
This project need cassandra DB to be launch
```
docker run --name some-cassandra -d cassandra
```

### Build and launch localy
```
docker build -t wereport-backend .
docker run \
    -p 8080:8080 \
    -p 8081:8081 \
    --link some-cassandra \
    --env CASSANDRA_HOSTS=172.17.0.2 \
    -it wereport-backend
```

### Install in GKE
See http://github.com/wescale/wereport-iac/

## Documentation

