FROM golang:1.22.1 AS builder
WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o healthcheck ./services/healthcheck/main.go

# glibcのVerに合わせたイメージを使用する https://repology.org/project/glibc/versions
FROM ubuntu:22.04  
COPY --from=builder /workspace/healthcheck /app/healthcheck
COPY config.yml /app/config.yml 
CMD ["/app/healthcheck"]

