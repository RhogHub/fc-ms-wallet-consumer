FROM golang:1.21.5

WORKDIR /app/

RUN apt-get update && apt-get install -y librdkafka-dev

COPY . .

CMD ["go", "run", "./cmd/walletcore/main.go"]