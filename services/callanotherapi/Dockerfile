FROM golang:1.22.0 AS builder
WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -tags=micro -o /workspace/callanotherapi ./services/callanotherapi/main.go

# glibcのVerに合わせたイメージを使用する https://repology.org/project/glibc/versions
FROM ubuntu:22.04  
RUN apt-get update && apt-get install -y \
    ca-certificates \
    && update-ca-certificates
COPY --from=builder /workspace/callanotherapi /callanotherapi
ENV INTERNAL_API_BASE_URL="https://healthcheck-rftndbrsdq-an.a.run.app"
CMD ["/callanotherapi"]
