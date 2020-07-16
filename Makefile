
build:
	@echo "Building Paxos Server"	
	go build -o paxos/paxos main.go

test:
	@echo "Testing GoPaxos"	
	go test -v --cover ./...

provision:
	@echo "Provisioning Paxos Cluster"	
	bash scripts/provision.sh

paxos-build:
	@echo "Building Paxos Docker Image"	
	docker build -t paxos -f Dockerfile .

paxos-run:
	@echo "Running Single Paxos Docker Container"
	docker run -p 8080:8080 -d paxos

info:
	echo "Paxos Cluster Nodes"
	docker ps | grep 'paxos'
	docker network ls | grep paxos_network

clean:
	@echo "Cleaning Paxos Cluster"
	docker ps -a | awk '$$2 ~ /paxos/ {print $$1}' | xargs -I {} docker rm -f {}
	docker network rm paxos_network