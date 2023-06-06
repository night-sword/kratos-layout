package service

import (
	"context"

	pb "github.com/night-sword/kratos-layout/api/service/v1"
)

type HealthService struct {
	pb.UnimplementedHealthServer
}

func NewHealthService() *HealthService {
	return &HealthService{}
}

func (s *HealthService) HealthCheck(ctx context.Context, req *pb.HealthRequest) (*pb.HealthReply, error) {
	return &pb.HealthReply{
		Status: pb.Status_UP,
	}, nil
}
