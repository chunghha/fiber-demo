ARG GO_VERSION=1.17

FROM golang:${GO_VERSION}-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o fiber-demo .

RUN go get -u github.com/swaggo/swag/cmd/swag
RUN swag init

FROM alpine:latest

RUN addgroup -S apigroup
RUN adduser -S -D -h /api apiuser apigroup

WORKDIR /api

COPY --from=builder /build/.env .
COPY --from=builder /build/docs .
COPY --from=builder /build/fiber-demo .

RUN chown -R apiuser:apigroup /api
USER apiuser

EXPOSE 3000

ENTRYPOINT ["./fiber-demo"]
