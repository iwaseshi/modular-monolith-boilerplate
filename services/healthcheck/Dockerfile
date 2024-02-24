FROM golang:1.22.0 AS builder
WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o healthcheck ./services/healthcheck/main.go

FROM gcr.io/distroless/base
COPY --from=builder /workspace/healthcheck /healthcheck
CMD ["/healthcheck"]
