FROM lichessbotdevs/lichess-bot

ARG BINARY

WORKDIR /lichess-bot

COPY ./build/config.yml .
COPY ./build/${BINARY} ./chester
RUN chmod +x ./chester

CMD ["python", "lichess-bot.py"]