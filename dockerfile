FROM golang:latest

WORKDIR /go/app

RUN api-get update && api-get install -y librdkafka-dev

CMD [ "tail", "-f", "/dev/null" ]