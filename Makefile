
build:
	@echo "Building Paxos Server"	
	go build -o paxos/paxos main.go

provision:
	@echo "Provisioning Paxos Cluster"	
	bash scripts/provision.sh

paxos-build:
	@echo "Building Paxos Docker Image"	
	docker build -t paxos -f Dockerfile .

paxos-run:
	@echo "Running Single Paxos Docker Container"
	docker run -p 8080:8080 -d paxos

clean:
	@echo "Cleaning Paxos Cluster"
	docker ps -a | awk '$$2 ~ /paxos/ {print $$1}' | xargs -I {} docker rm -f {}