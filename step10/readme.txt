#mount host dir into minikube (separate terminal!)
minikube mount /data/mazepv:/mnt/mazepv


kubectl apply -f redis-deployment.yaml
kubect get pod
kubectl get services
kubectl create -f worker-deployment.yaml
kubectl logs maze-worker-deployment-7d8ff78c44-fgx5z
kubectl create -f controller-deployment.yaml
kubectl get pods
kubectl logs maze-controller-deployment-74cf9f85f8-g8r2j
kubectl get svc

minikube service list

curl "http://192.168.99.100:31752/maze?text=a"
curl "http://192.168.99.100:31752/token/90fb833e-2c37-4d25-a6f3-2d43abc1d4ca"

./fetch.sh "kuba+i+tomek" | pbcopy

#see shared folder :)

