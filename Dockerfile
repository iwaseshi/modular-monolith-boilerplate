FROM golang:1.22.1 AS builder
WORKDIR /workspace
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o monoapp

# glibcのVerに合わせたイメージを使用する https://repology.org/project/glibc/versions
FROM ubuntu:22.04  
COPY --from=builder /workspace/monoapp /monoapp
EXPOSE 8080
CMD ["/monoapp"]
