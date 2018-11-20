# Pushing & pulling images

docker login
docker tag maze-gen:0.0.1 tsmsap/mazegen:0.0.1
docker push tsmsap/mazegen:0.0.1

#/generated comes from step2/Dockerfile
docker run -v "$PWD"/images:/generated tsmsap/mazegen:0.0.1 "i am kuba"
docker rmi tsmsap/mazegen:0.0.1
docker run -v "$PWD"/images:/generated tsmsap/mazegen:0.0.1 "kuba tomek"

