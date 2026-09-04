FROM lichessbotdevs/lichess-bot

# TARGETARCH is supplied automatically by buildx, once per platform, so a
# single multi-platform build can pick the matching binary for each image in
# the manifest. Its values -- amd64, arm64 -- are the same strings GOARCH
# uses, which is what makes the binaries line up by name.
ARG TARGETARCH

WORKDIR /lichess-bot

COPY ./config.yml .

# --chmod avoids a RUN step, so an image for a foreign architecture is
# assembled by copying layers instead of emulating a shell under QEMU.
COPY --chmod=0755 ./build/chester-${TARGETARCH} ./chester

CMD ["python", "lichess-bot.py"]
