#!/bin/bash
set -e

VERSION=0.1.1
IMAGE=bluescreen10/chester

# create build dir
mkdir -p build

# generate config.yml from template
cp config.template.yml build/config.yml

# build engine amd64
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X 'main.version=$VERSION'" -o build/chester ./cmd/

# build the Docker image
docker buildx create --use --name multiarch-builder || docker buildx use multiarch-builder
docker buildx build --platform linux/amd64 -t $IMAGE:amd64-latest --push .

# build engine arm64
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w -X 'main.version=$VERSION'" -o build/chester ./cmd/

# build the Docker image
docker buildx create --use --name multiarch-builder || docker buildx use multiarch-builder
docker buildx build --platform linux/arm64 -t $IMAGE:arm64-latest --push .

docker buildx imagetools create \
  -t $IMAGE:latest -t $IMAGE:$VERSION \
  $IMAGE:amd64-latest \
  $IMAGE:arm64-latest

# run
#docker run -it --env-file .env bluescreen10/chester:latest /bin/bash