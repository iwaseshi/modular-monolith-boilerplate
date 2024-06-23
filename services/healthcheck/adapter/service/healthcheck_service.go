package service

import (
	"context"
	pb "modular-monolith-boilerplate/pkg/adapter/rpc"
	"modular-monolith-boilerplate/pkg/di"
	"modular-monolith-boilerplate/services/healthcheck/domain"
	"modular-monolith-boilerplate/services/healthcheck/usecase"
)

func init() {
	di.RegisterBean(NewServer)
}

type HealthCheckServer struct {
	pb.UnimplementedHealthCheckServiceServer
	healthCheckUseCase usecase.HealthCheckUseCase
}

func NewServer(healthCheckUseCase usecase.HealthCheckUseCase) *HealthCheckServer {
	return &HealthCheckServer{
		healthCheckUseCase: healthCheckUseCase,
	}
}

func (s *HealthCheckServer) Ping(ctx context.Context, in *pb.PingRequest) (*pb.PingResponse, error) {
	message, err := s.healthCheckUseCase.Ping(ctx)
	if err != nil {
		return nil, err.Unwrap()
	}
	return &pb.PingResponse{Message: *message}, nil
}

func (s *HealthCheckServer) Readiness(ctx context.Context, in *pb.ReadyRequest) (*pb.ReadyResponse, error) {
	req := &domain.ReadyRequest{}
	// Bind JSON request to req (if necessary)
	res, err := s.healthCheckUseCase.Readiness(ctx, req)
	if err != nil {
		return nil, err.Unwrap()
	}
	return &pb.ReadyResponse{Message: res.Message}, nil
}
