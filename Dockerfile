# FROM ubuntu:20.04 AS dependencies
# RUN apt update
# RUN apt install -y ifstat=1.1-8.1build2
# RUN apt install -y ethtool=1:5.4-1

FROM golang:1.18.0 AS builder
WORKDIR /build
COPY go.mod go.sum ./
RUN GOOS=linux GOARCH=amd64 go mod download
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o sysmetric cmd/main.go

FROM --platform=linux/amd64 ubuntu:20.04
WORKDIR /app
RUN apt update
RUN apt install -y ifstat=1.1-8.1build2
RUN apt install -y ethtool=1:5.4-1
RUN apt install -y vnstat=2.6-1
# COPY --from=dependencies /sbin/ethtool /app/
# COPY --from=dependencies /usr/bin/ifstat /app/
COPY --from=builder /build/sysmetric .
CMD ["/app/sysmetric"]