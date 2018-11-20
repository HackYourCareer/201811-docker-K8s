minikube start
#mount host dir into minikube (separate terminal!)
minikube mount /data/mazepv:/mnt/mazepv

#deploy workers

#show that shared dir is mounted to worker as expected!
kubectl exec -it maze-worker-deployment-649b7b58c-2ckww /bin/sh
