FROM golang:1.19-alpine AS base_image
WORKDIR /example
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./

FROM base_image AS run_tests
RUN go test .
