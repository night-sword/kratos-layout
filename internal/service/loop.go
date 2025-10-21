package service

import (
	"github.com/night-sword/kratos-layout/internal/conf"
)

type LoopService struct {
	cfg *conf.Bootstrap
}

func NewLoopService(cfg *conf.Bootstrap) (inst *LoopService) {
	return &LoopService{
		cfg: cfg,
	}
}
