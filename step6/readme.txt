# Prototyping the solution: Worker
# Worker listens on http and generates images. After generating it makes an entry in Redis.

1) Run redis in docker
docker run --name maze-redis --rm -p 127.0.0.1:6379:6379 -d redis
docker run -it --link maze-redis:redis --rm redis redis-cli -h redis -p 6379
#Redis TTL: TTL <key>

2) Worker: Explain + Diagram
   Worker assumes "success" and puts an entry into redis.
./build_worker.sh
file maze-worker
env WORKER_DIR="./worker" MAZEGEN_CMD="./test.sh" REDIS_ADDR="127.0.0.1:6379" ./maze-worker
curl -i "http://localhost:8081/generate?token=123&text=abc"
docker run -it --link maze-redis:redis --rm redis redis-cli -h redis -p 6379

3) Docker image:

#TODO1
./build_worker.sh
file maze-worker
docker build -t tsmsap/mazegen-worker:0.0.1 .
docker run -v "$PWD"/images:/workers -e REDIS_ADDR=host.docker.internal:6379 -d -p 127.0.0.1:8081:8081 --rm --name worker tsmsap/mazegen-worker:0.0.1
docker ps
docker logs -f worker
curl -i "http://localhost:8081/generate?token=123&text=abc"
docker inspect -f "{{ .Mounts }}" worker
