FROM golang:1.22.1 AS builder
WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -tags=micro -o /workspace/intermediary ./services/intermediary/main.go

# glibcのVerに合わせたイメージを使用する https://repology.org/project/glibc/versions
FROM ubuntu:22.04  
# 最新のセキュリティパッチを受け取るためにバージョンの明示的な指定は行わない。
# hadolint ignore=DL3008
RUN apt-get update && apt-get install -y \
    --no-install-recommends\
    ca-certificates \
    && update-ca-certificates \
    && rm -rf /var/lib/apt/lists/*
COPY --from=builder /workspace/intermediary /intermediary
ENV INTERNAL_API_BASE_URL="https://healthcheck-rftndbrsdq-an.a.run.app"
CMD ["/intermediary"]
