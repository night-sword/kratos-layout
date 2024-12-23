package server

import (
	km "github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/night-sword/kratos-kit/log"
	"github.com/night-sword/kratos-kit/middleware"

	v1 "github.com/night-sword/kratos-layout/api/service/v1"
	"github.com/night-sword/kratos-layout/internal/conf"
	"github.com/night-sword/kratos-layout/internal/service"
)

func NewHTTPServer(cfg *conf.Bootstrap, health *service.Health) *http.Server {
	c := cfg.GetServer()
	ms := []km.Middleware{
		recovery.Recovery(),
		middleware.LogServer(log.GetLogger()),
		validate.Validator(),
	}

	ak, sk := cfg.GetBusiness().GetToken().GetAk(), cfg.GetBusiness().GetToken().GetSk()
	if ak != "" && sk != "" {
		ms = append(ms, middleware.HeaderSK(ak, sk))
	}

	var opts = []http.ServerOption{
		http.Middleware(ms...),
	}

	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterHealthHTTPServer(srv, health)

	return srv
}
