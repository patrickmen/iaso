# Build the docker image
docker-build:
	docker build -f ./Dockerfile -t patrickmen/iaso/iaso:v1.2.0 .