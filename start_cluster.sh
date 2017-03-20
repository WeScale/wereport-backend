#!/bin/bash

cd ~/Code/weReport && docker build -t wereport-backend .

# eval $(docker-machine env -u)
# docker-machine create \
#       --engine-env 'DOCKER_OPTS="-H unix:///var/run/docker.sock"' \
#       --driver virtualbox \
#       registry-vm

# eval $(docker-machine env registry-vm)
# docker run -d -p 5000:5000 --name registry registry:2
# eval $(docker-machine env -u)

# ip_registry=$(docker-machine ip registry-vm)
# docker tag cassandra $ip_registry:5000/cassandra
# docker tag wereport-backend $(docker-machine ip registry-vm):5000/wereport-backend

# docker push $ip_registry:5000/cassandra
# docker push $(docker-machine ip registry-vm):5000/wereport-backend

minikube start --insecure-registry=$(docker-machine ip registry-vm):5000
minikube dashboard
minikube addons enable heapster

kubectl create -f kubernetes_files/cassandra.yaml

echo "Minikube ip: $(minikube ip)"

while ! minikube addons open heapster; do
    sleep 1
done

