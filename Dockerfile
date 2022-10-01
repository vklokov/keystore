FROM golang:1.19.1-alpine3.16

RUN apk update && apk upgrade

RUN apk add --no-cache \
    bash \
    curl \
    build-base \
    tzdata \
    postgresql14-contrib \
    postgresql14-dev \
    postgresql14-client

WORKDIR /app

COPY go.sum go.mod .

RUN go mod download
