FROM golang:1.17.1-alpine as builder

WORKDIR /build
COPY ../../go.mod ../../go.sum ./
RUN go mod download
COPY ../../  ./

ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build -ldflags '-s -w' ./cmd/sample-api

FROM alpine
COPY --from=builder /build/sample-api /opt/app/
ENTRYPOINT ["/opt/app/sample-api"]