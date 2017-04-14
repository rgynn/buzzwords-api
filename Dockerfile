FROM golang:1.8.0-alpine

ENV PORT "3000"
EXPOSE 3000

ADD . /go/src/github.com/hajhatten/buzzwords-api

RUN apk update && \
    apk upgrade && \
    apk add git && \
    rm -rf /var/cache/apk/*


RUN go install github.com/hajhatten/buzzwords-api

CMD ["buzzwords-api"]