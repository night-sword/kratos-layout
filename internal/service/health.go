package service

import (
	"context"

	pb "github.com/night-sword/kratos-layout/api/service/v1"
)

type Health struct {
	pb.UnimplementedHealthServer
}

func NewHealth() *Health {
	return &Health{}
}

func (s *Health) HealthCheck(_ context.Context, _ *pb.HealthRequest) (*pb.HealthReply, error) {
	return &pb.HealthReply{
		Status: pb.Status_UP,
	}, nil
}
