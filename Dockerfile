FROM golang:1.18

WORKDIR /go/src

RUN apt-get update && apt-get install build-essential librdkafka-dev -y

CMD ["tail", "-f", "/dev/null"]