FROM golang:1.19.1-alpine3.16

ENV MIGRATE_VERSION 4.15.2

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

RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.2

COPY go.sum go.mod .

RUN go mod download
