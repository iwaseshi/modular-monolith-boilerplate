package service

import (
	"context"
	pb "modular-monolith-boilerplate/pkg/adapter/rpc"
	"modular-monolith-boilerplate/services/healthcheck/domain"
	"modular-monolith-boilerplate/services/healthcheck/usecase"
)

type server struct {
	pb.UnimplementedHealthCheckServiceServer
	healthCheckUseCase usecase.HealthCheckUseCase
}

func NewServer(healthCheckUseCase usecase.HealthCheckUseCase) *server {
	return &server{
		healthCheckUseCase: healthCheckUseCase,
	}
}

func (s *server) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	message, err := s.healthCheckUseCase.Ping(ctx)
	if err != nil {
		return nil, err.Unwrap()
	}
	return &pb.PingResponse{Message: *message}, nil
}

func (s *server) Readiness(ctx context.Context, in *pb.ReadyRequest) (*pb.ReadyResponse, error) {
	req := &domain.ReadyRequest{}
	// Bind JSON request to req (if necessary)
	res, err := s.healthCheckUseCase.Readiness(ctx, req)
	if err != nil {
		return nil, err.Unwrap()
	}
	return &pb.ReadyResponse{Message: res.Message}, nil
}
