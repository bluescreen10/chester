FROM lichessbotdevs/lichess-bot

WORKDIR /lichess-bot

COPY ./build/ .

CMD ["python", "lichess-bot.py"]