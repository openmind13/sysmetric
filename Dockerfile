FROM golang:1.18.0-alpine3.15 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN GOOS=linux GOARCH=amd64 go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o sysmetric cmd/main.go

FROM --platform=linux/amd64 alpine:3.15
WORKDIR /app
RUN ifconfig
RUN apk add ethtool
COPY --from=builder /build/sysmetric .
CMD ["/app/sysmetric"]