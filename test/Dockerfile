FROM golang:1.19-alpine AS base_image
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY integration.go ./
RUN go get github.com/gardusig/template-go@${TEMPLATE_PACKAGE_VERSION}

FROM base_image AS integration_tests
RUN go test .
