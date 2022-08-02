#!/usr/bin/env sh

kubectl create ns prometheus-operator
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install -n prometheus-operator prometheus-operator prometheus-community/kube-prometheus-stack -f k8s/values/prometheus-operator-values.yaml

kubectl create ns ingress-nginx
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx/
helm repo update
helm install -n ingress-nginx ingress-nginx ingress-nginx/ingress-nginx -f k8s/values/ingress-nginx-values.yaml

kubectl create ns postgres-exporter
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
helm install -n postgres-exporter postgres-exporter prometheus-community/prometheus-postgres-exporter -f k8s/values/postgres-exporter.yaml

kubectl create ns otus-msa-hw3