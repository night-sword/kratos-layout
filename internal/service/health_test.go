package service

import (
	"context"
	"testing"

	v1 "github.com/night-sword/kratos-layout/api/service/v1"
)

func TestHealth_HealthCheck(t *testing.T) {
	h := NewHealth()
	reply, err := h.HealthCheck(context.Background(), &v1.HealthRequest{})
	if err != nil {
		t.Fatalf("HealthCheck() error = %v", err)
	}
	if reply.Status != v1.Status_UP {
		t.Errorf("HealthCheck() status = %v, want %v", reply.Status, v1.Status_UP)
	}
}
