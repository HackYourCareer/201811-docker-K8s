#!/usr/bin/env bash

kubectl apply -f worker-deployment.yaml
kubectl apply -f controller-deployment.yaml

