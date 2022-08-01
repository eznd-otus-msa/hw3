build:
	docker build -f docker/Dockerfile . -t eznd/otus-msa-hw3:latest

push:
	docker push eznd/otus-msa-hw3:latest

k8s-deploy:
	helm/deploy.sh

k8s-remove:
	helm/remove.sh
