#Test if it works with direct IP connection

#get IP address of one of the workers
kubectl get pods -o wide

#update controller deployment with IP address of one worker
172.17.0.7

#fetch IP address of the controller
#ssh into minikube and curl the controller
curl "http://<controller-IP>:8080/maze?text=hack"
# look for token in response, check shared folder for data :)

