#!/bin/bash

# set project name
PROJECT_NAME="YourProjectName"
CONTAINER_NAME=$PROJECT_NAME+"-container"
IMAGE_NAME=$PROJECT_NAME+":latest"

# create logs directory
if [ ! -d "logs" ]; then
  echo "Creating directory: logs"
  mkdir logs
fi

# check container is exists and stop
if [ "$(docker ps -aq -f name=$CONTAINER_NAME)" ]; then
    echo "Stopping and removing container: $CONTAINER_NAME"
    docker stop $CONTAINER_NAME
    docker rm $CONTAINER_NAME
else
    echo "No container named $CONTAINER_NAME found."
fi

# check image is exists and remove
if [ "$(docker images -q $IMAGE_NAME)" ]; then
    echo "Removing image: $IMAGE_NAME"
    docker rmi $IMAGE_NAME
else
    echo "No image named $IMAGE_NAME found."
fi

# build new image
echo "Building new image: $IMAGE_NAME"
docker build -t $IMAGE_NAME .


# start container
echo "Running container: $CONTAINER_NAME"
docker run --name $CONTAINER_NAME -d -v ${PWD}/logs:/logs $IMAGE_NAME

