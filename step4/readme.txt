# Interaction with container: Mounting volumes

#first, show that this dir is empty!
#/generated comes from step2/Dockerfile

docker run -v "$PWD"/images:/generated maze-gen:0.0.1

docker run -v "$PWD"/images:/generated maze-gen:0.0.1 "hello tomek"

# Explain, the we can run images like tools, but some images are just meant to be "base" for others and don't have defined entrypoints.
