# Dockerizing


source ../golang/setup.source.me

file maze

docker build -t tomek:latest .

docker run tomek:latest 3 3
#TODO1

./build.sh
docker build -t tomek:latest .
docker run tomek:latest 3 3

docker ps -a | head
docker commit 8e9f6b65afa2 tmp_img
docker run --entrypoint=sh -it tmp_img

alias rerun='docker run --entrypoint=sh -it tmp_img'

docker run tomek:latest hello
