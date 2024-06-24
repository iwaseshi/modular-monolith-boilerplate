package grpc

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	server = grpc.NewServer()
)

const (
	DefaultPort = 50051
)

func GetServer() *grpc.Server {
	return server
}

func Run(port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
	env := os.Getenv("ENVIRONMENT")
	if env != "Prod" {
		reflection.Register(server)
	}

	go func() {
		log.Printf("start gRPC server port: %v", port)
		if err := server.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	server.GracefulStop()
}
