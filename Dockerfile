FROM golang:1.22.2 AS builder
WORKDIR /workspace
COPY . .
RUN go mod tidy && \
    go mod download && \
    go build -o monoapp

# glibcのVerに合わせたイメージを使用する https://repology.org/project/glibc/versions
FROM ubuntu:24.04 
ENV MODE="mono" 
WORKDIR /app
COPY --from=builder /workspace/monoapp /app/monoapp
COPY config.yml /app/config.yml 
EXPOSE 8080
CMD ["/app/monoapp"]
