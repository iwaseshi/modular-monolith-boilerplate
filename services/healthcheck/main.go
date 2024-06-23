package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	pb "modular-monolith-boilerplate/pkg/adapter/rpc"
	"modular-monolith-boilerplate/pkg/di"
	"modular-monolith-boilerplate/services/healthcheck/adapter/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

//nolint:unused
func main() {
	// controller.RegisterRouting()
	// _ = restapi.Run(restapi.DefaultPort)

	port := 8080
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	if err := di.GetContainer().Invoke(func(server *service.HealthCheckServer) {
		pb.RegisterHealthCheckServiceServer(s, server)
	}); err != nil {
		log.Fatalf("failed to invoke container: %v", err)
	}
	reflection.Register(s)

	go func() {
		log.Printf("start gRPC server port: %v", port)
		if err := s.Serve(listener); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// 4.Ctrl+Cが入力されたらGraceful shutdownされるようにする
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("stopping gRPC server...")
	s.GracefulStop()
}
