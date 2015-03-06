# FHIR Docker

MITRE's [Intervention Engine FHIR server](https://github.com/intervention-engine/fhir) in a Docker container

## Installation

Intervention Engine requires a MongoDB server. Included is are Dockerfiles for a
MongoDB data volume container and the MongoDB database container.

Create and run the Data Volume Container and the MongoDB container, then run the FHIR server container. Note the use of named containers.


	docker build -f Dockerfile.datavolume -t local/mongo_datavol .
	docker build -f Dockerfile.mongodb -t local/mongodb .
	docker build -f Dockerfile.fhir -t local/fhir .

	docker run -d --name mongo_datavol local/mongo_datavol
	docker run -d -p 27017:27017 --volumes-from mongo_datavol --name mongodb local/mongodb
	docker run -d --link mongodb:mongodb -p 3001:3001 --name fhir local/fhir
