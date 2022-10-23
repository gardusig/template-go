FROM golang:1.19-alpine AS base
WORKDIR /app
COPY go.mod ./
COPY *.go ./
RUN go mod tidy

FROM base AS build
RUN go build

FROM base AS unit_tests
RUN go test
