FROM golang:1.8.0-alpine

ENV PORT "3000"
EXPOSE 3000

WORKDIR /go/src/github.com/hajhatten/buzzwords-api
ADD . /go/src/github.com/hajhatten/buzzwords-api

RUN apk update && \
    apk upgrade && \
    apk add git

RUN go get github.com/prometheus/client_golang/prometheus
RUN go get github.com/hajhatten/buzzwords

RUN go install github.com/hajhatten/buzzwords-api

CMD ["buzzwords-api"]