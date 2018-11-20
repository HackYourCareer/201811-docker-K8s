Generating and converting, composing docker images

change something in base image to repeat entire Alpine install procedure

docker build -t maze-gen:0.0.1 .
#for the first time you can see installing alpine packages, show that next run is shorter!
docker build -t maze-gen:0.0.1 .

docker run maze-gen:0.0.1

docker ps -a | head
docker commit 8e9f6b65afa2 temporary_image
docker run --entrypoint=sh -it temporary_image
