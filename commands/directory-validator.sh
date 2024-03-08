#!/usr/bin/env bash

DOCKER_IMAGE_INTERNAL_DIRECTORY_VALIDATOR="internal-directory-validator:latest"

if [ "$(docker inspect --type=image $DOCKER_IMAGE_INTERNAL_DIRECTORY_VALIDATOR 2> /dev/null)" == "[]" ]; then
    docker build -t $DOCKER_IMAGE_INTERNAL_DIRECTORY_VALIDATOR -f $(pwd)/commands/Dockerfile.directory-validator .
fi

docker run --rm -w=/go/app -v=$(pwd):/go/app $DOCKER_IMAGE_INTERNAL_DIRECTORY_VALIDATOR directory-validator .
