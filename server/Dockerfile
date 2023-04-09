# syntax=docker/dockerfile:1

FROM golang:1.20.2-alpine3.16 as builder

WORKDIR /usr/local/app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o bin/binnacle-server

RUN GRPC_HEALTH_PROBE_VERSION=v0.4.17 && \
    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
    chmod +x /bin/grpc_health_probe

FROM golang:1.20.2-alpine3.16

WORKDIR /usr/local/app

COPY --from=builder /usr/local/app/bin/binnacle-server ./binnacle-server
COPY --from=builder /bin/grpc_health_probe ./grpc_health_probe