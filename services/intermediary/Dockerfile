FROM golang:1.22.1 AS builder
WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy
COPY . .
RUN go build -tags=micro -o /workspace/intermediary ./services/intermediary/main.go

# glibcのVerに合わせたイメージを使用する https://repology.org/project/glibc/versions
FROM ubuntu:22.04  
ENV MODE="micro"
# 最新のセキュリティパッチを受け取るためにバージョンの明示的な指定は行わない。
# hadolint ignore=DL3008
RUN apt-get update && apt-get install -y \
    --no-install-recommends\
    ca-certificates \
    && update-ca-certificates \
    && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /workspace/intermediary /app/intermediary
COPY config.yml /app/config.yml 
CMD ["/app/intermediary"]
