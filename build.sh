#!/bin/bash
set -e

VERSION=0.4.2
IMAGE=bluescreen10/chester

# create build dir
mkdir -p build

# 1. Regenerate the PGO profile
#
# cmd/default.pgo is picked up automatically by go build. It has to cover both
# move generation and search, or the compiler optimizes only half the engine.
# Profiles key on function names, not machine code, so one generated here
# applies to every architecture built below.
#
# The generator itself is built without PGO (there is no default.pgo in its own
# directory), which keeps each profile a measurement of the plain build rather
# than of the previous profile's decisions.
#
# Set SKIP_PGO=1 to reuse the committed profile, which is what you want on a
# loaded or shared machine: a profile collected under contention misattributes
# time and is worse than a slightly stale one.
if [ "${SKIP_PGO:-0}" = "1" ]; then
  echo "--- Skipping PGO regeneration (SKIP_PGO=1) ---"
else
  echo "--- Regenerating PGO profile ---"
  go run ./internal/cmd/pgo
fi

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
