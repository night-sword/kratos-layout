package repo

import (
	"context"
	"io"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/night-sword/kratos-kit/log"
	"github.com/night-sword/kratos-kit/middleware"
	googlegrpc "google.golang.org/grpc"

	"github.com/night-sword/kratos-layout/internal/conf"
)

func closeResource(r io.Closer, name string) {
	log.Infof("start closing the %s resources...", name)
	if r != nil {
		log.E(r.Close())
	}
	log.Infof("finish close %s resources", name)
}

func NewGrpcConnDirect(cfg *conf.Data_Service) (conn googlegrpc.ClientConnInterface, err error) {
	return grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(cfg.GetHost()),
		grpc.WithMiddleware(
			recovery.Recovery(),
			middleware.FormatError(),
		),
		grpc.WithTimeout(cfg.GetTimeout().AsDuration()),
	)
}
