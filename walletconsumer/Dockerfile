FROM golang:1.21.5

WORKDIR /consumerapp/

RUN apt-get update && apt-get install -y librdkafka-dev

COPY . .

#CMD ["tail", "-f", "/dev/null"]
CMD ["go", "run", "./cmd/walletconsumer/main.go"]