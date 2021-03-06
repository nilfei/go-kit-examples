package gokit_foundation

import (
	"context"
	"google.golang.org/grpc/health/grpc_health_v1"
)

/*
grpc的健康检查接口，提供给consul调用
*/
type HealthCheckServer struct{}

func NewHealthCheckSrv() *HealthCheckServer {
	return &HealthCheckServer{}
}

func (s *HealthCheckServer) Check(_ context.Context, _ *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	//log.Println("health Checking...")
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (s *HealthCheckServer) Watch(*grpc_health_v1.HealthCheckRequest, grpc_health_v1.Health_WatchServer) error {
	return nil
}
