# Build the docker image
docker-build:
	docker build -f ./Dockerfile -t iaso/iaso:v1.2.1 .