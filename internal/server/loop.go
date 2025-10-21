package server

import (
	"context"

	"github.com/night-sword/kratos-layout/internal/conf"
	"github.com/night-sword/kratos-layout/internal/service"
)

type LoopServer struct {
	cfg    *conf.Bootstrap
	cancel context.CancelFunc
	loop   *service.LoopService
}

func NewLoopServer(cfg *conf.Bootstrap, loop *service.LoopService) *LoopServer {
	return &LoopServer{
		cfg:    cfg,
		cancel: nil,
		loop:   loop,
	}
}

func (inst *LoopServer) Start(ctx context.Context) (err error) {
	_, cancel := context.WithCancel(ctx)
	inst.cancel = cancel
	return
}

func (inst *LoopServer) Stop(_ context.Context) (err error) {
	inst.cancel()
	return
}
