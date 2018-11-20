#!/usr/bin/env bash

kubectl apply -f redis-deployment.yaml
kubectl apply -f worker-deployment.yaml
kubectl apply -f controller-deployment.yaml

