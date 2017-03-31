FROM golang:1.8.0-alpine

ENV PORT "3000"
EXPOSE 3000

COPY . /go/src/github.com/hajhatten/buzzwords-api
COPY buzzwords.json /go/buzzwords.json

RUN go install github.com/hajhatten/buzzwords-api

CMD ["buzzwords-api"]