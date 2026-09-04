#!/bin/bash
set -e

# VERSION is the release being built. The release workflow passes the git tag;
# a local run falls back to the most recent tag, and to "dev" outside a
# checkout. A leading "v" is stripped, so the tag v1.2.3 produces image tag
# 1.2.3.
VERSION="${VERSION:-$(git describe --tags --always 2>/dev/null || echo dev)}"
VERSION="${VERSION#v}"

# IMAGE is the repository to publish to. Override it to publish elsewhere; the
# release workflow points it at ghcr.io.
IMAGE="${IMAGE:-bluescreen10/chester}"

echo "--- Releasing $IMAGE:$VERSION ---"

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

# 2. Setup Docker Buildx once
docker buildx create --use --name multiarch-builder 2>/dev/null || docker buildx use multiarch-builder

# 3. Cross-compile both architectures
#
# Go does the cross-compiling, so the Dockerfile never has to run anything and
# neither build needs emulation.
echo "--- Building binaries ---"
LDFLAGS="-s -w -X 'main.version=$VERSION'"
GOOS=linux GOARCH=amd64 go build -ldflags="$LDFLAGS" -o build/chester-amd64 ./cmd/
GOOS=linux GOARCH=arm64 go build -ldflags="$LDFLAGS" -o build/chester-arm64 ./cmd/

# 4. Build and push both images as one manifest
#
# A single multi-platform build writes the manifest list directly, so there is
# no need to push per-architecture tags and stitch them together afterwards.
# The Dockerfile selects the binary by TARGETARCH.
echo "--- Building and pushing $IMAGE ---"
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  -t "$IMAGE:$VERSION" \
  -t "$IMAGE:latest" \
  --push .

echo "Successfully pushed $IMAGE:$VERSION and $IMAGE:latest"
