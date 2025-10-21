package service

import (
	"context"

	v1 "github.com/night-sword/kratos-layout/api/service/v1"
)

type Health struct {
	v1.UnimplementedHealthServer
}

func NewHealth() *Health {
	return &Health{}
}

func (s *Health) HealthCheck(_ context.Context, _ *v1.HealthRequest) (*v1.HealthReply, error) {
	return &v1.HealthReply{
		Status: v1.Status_UP,
	}, nil
}
