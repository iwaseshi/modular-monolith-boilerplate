FROM golang:1.22.0 AS builder
WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -tags=micro -o callanotherapi ./services/callanotherapi/main.go

FROM gcr.io/distroless/base
COPY --from=builder /workspace/callanotherapi /callanotherapi
CMD ["/callanotherapi"]
