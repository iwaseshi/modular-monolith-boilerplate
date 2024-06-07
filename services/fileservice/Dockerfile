FROM golang:1.22.2 AS builder
WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o fileservice ./services/fileservice/main.go

# glibcのVerに合わせたイメージを使用する https://repology.org/project/glibc/versions
FROM ubuntu:24.04 
# 最新のセキュリティパッチを受け取るためにバージョンの明示的な指定は行わない。
# hadolint ignore=DL3008
RUN apt-get update && apt-get install -y \
    --no-install-recommends\
    ca-certificates \
    && update-ca-certificates \
    && rm -rf /var/lib/apt/lists/*
COPY --from=builder /workspace/fileservice /app/fileservice
COPY config.yml /app/config.yml 
CMD ["/app/fileservice"]

