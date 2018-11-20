# Prototyping the solution: Controller
# Controller listens on http.
# /maze flow
1) A Request /maze from an user
2) Controller generates token
3) Controller dispatches the call to workers
4) Controller responds with 200 OK and the payload: {"tokenUrl": "http://..../token/...."}

# /token flow
1) A Request /token from an user
2) Controller checks in Redis using token
3) Controller responds with 404 if not found
4) Controller responds with 200 OK and the payload: {"image": "http://../maze.original.png", "printVersion": "http://../maze.enhanced.png"}

1) Run standalone
./build_controller.sh
file maze-controller

env WORKER_DIR="./worker" WORKER_HOST="localhost:8081" REDIS_ADDR="127.0.0.1:6379" ./maze-controller
or with worker running in background (docker run)
env WORKER_DIR="../step6/images" WORKER_HOST="localhost:8081" REDIS_ADDR="127.0.0.1:6379" ./maze-controller

curl -i "http://localhost:8080/maze?text=a"
curl -i "http://localhost:8080/maze?height=4&width=4"
open ../step6/images/123/*.png


2) Docker image:
./build_controller.sh
file maze-controller
docker build -t tsmsap/mazegen-controller:0.0.1 .
docker run -v "$PWD"/../step6/images:/workers -d -p 127.0.0.1:8080:8080 --rm --name controller tsmsap/mazegen-controller:0.0.1
FULL version: docker run -v "$PWD"/../step6/images:/workers -e WORKER_HOST=host.docker.internal:8081 -e REDIS_ADDR=host.docker.internal:6379 -d -p 127.0.0.1:8080:8080 --rm --name controller tsmsap/mazegen-controller:0.0.1
docker ps
docker logs -f controller
curl -i "http://localhost:8080/maze?text=hack"
docker inspect -f "{{ .Mounts }}" worker
