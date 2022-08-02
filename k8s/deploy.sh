#!/usr/bin/env sh

kubectl create ns prometheus-operator
helm upgrade --install -n prometheus-operator prometheus-operator prometheus-community/kube-prometheus-stack -f k8s/prometheus-operator-values.yaml


kubectl create ns otus-msa-hw3
helm upgrade --install -n otus-msa-hw3 otus-msa-hw3 k8s/chart