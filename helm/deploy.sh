#!/usr/bin/env sh

kubectl create ns otus-msa-hw3
helm upgrade --install -n otus-msa-hw3 otus-msa-hw3 helm/chart