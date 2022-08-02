build:
	docker build -f docker/Dockerfile . -t eznd/otus-msa-hw3:latest

push:
	docker push eznd/otus-msa-hw3:latest

k8s-pre-reqs:
	k8s/pre-reqs.sh

k8s-deploy:
	k8s/deploy.sh

k8s-remove:
	k8s/remove.sh

loadtest:
	ab -c 5 -n 1000000000 http://arch.homework/api/users/1