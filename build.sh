#!/bin/bash
set -e

VERSION=0.3.0
IMAGE=bluescreen10/chester

# create build dir
mkdir -p build

# generate config.yml from template
cp config.template.yml build/config.yml

# 2. Setup Docker Buildx once
docker buildx create --use --name multiarch-builder 2>/dev/null || docker buildx use multiarch-builder

# 3. Build & Push AMD64
echo "--- Building AMD64 ---"
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w -X 'main.version=$VERSION'" -o build/chester-amd64 ./cmd/
docker buildx build --platform linux/amd64 \
  --build-arg BINARY=chester-amd64 \
  -t $IMAGE:amd64-latest --push .

# 4. Build & Push ARM64
echo "--- Building ARM64 ---"
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w -X 'main.version=$VERSION'" -o build/chester-arm64 ./cmd/
docker buildx build --platform linux/arm64 \
  --build-arg BINARY=chester-arm64 \
  -t $IMAGE:arm64-latest --push .

# 5. Create Multi-arch Manifest
echo "--- Creating Manifests ---"
docker buildx imagetools create \
  -t $IMAGE:latest -t $IMAGE:$VERSION \
  $IMAGE:amd64-latest \
  $IMAGE:arm64-latest

echo "Successfully pushed $IMAGE:$VERSION"