FROM golang:1.22.0 AS builder
WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o monoapp

FROM gcr.io/distroless/base
COPY --from=builder /workspace/monoapp /monoapp
CMD ["/monoapp"]
