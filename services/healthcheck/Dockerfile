FROM golang:1.22.0 AS builder
WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o healthcheck ./services/healthcheck/main.go

# glibcのVerに合わせたイメージを使用する https://repology.org/project/glibc/versions
FROM ubuntu:22.04  
COPY --from=builder /workspace/healthcheck /healthcheck
CMD ["/healthcheck"]
