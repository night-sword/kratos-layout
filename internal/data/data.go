package data

import (
	"context"
	"crypto/tls"
	"database/sql"
	"io"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/night-sword/kratos-kit/errors"
	"github.com/night-sword/kratos-kit/log"
	"github.com/night-sword/kratos-kit/middleware"
	"github.com/redis/go-redis/v9"
	etcdv3 "go.etcd.io/etcd/client/v3"
	googlegrpc "google.golang.org/grpc"

	"github.com/night-sword/kratos-layout/internal/conf"
)

type Data struct {
	cfg *conf.Bootstrap

	db        *sql.DB
	redis     *redis.Client
	discovery registry.Discovery
}

func NewData(cfg *conf.Bootstrap) (data *Data, cleanup func(), err error) {
	db, err := newDB(cfg.GetData().GetDatabase())
	if err != nil {
		return
	}
	_redis := newRedis(cfg.GetData().GetRedis())
	discovery, _etcd, err := newDiscovery(cfg.GetData().GetRegistrar())
	if err != nil {
		return
	}

	closer := []io.Closer{
		db, _redis, _redis, _etcd,
	}

	cleanup = func() {
		log.Info("closing the data resources...")
		for i := range closer {
			if closer[i] != nil {
				log.E(closer[i].Close())
			}
		}
		log.Info("finish close data resources")
	}

	data = &Data{
		cfg:       cfg,
		db:        db,
		redis:     _redis,
		discovery: discovery,
	}
	return
}

func (inst *Data) GetDiscovery() (discovery registry.Discovery, err error) {
	if inst.discovery == nil {
		err = errors.InternalServer(errors.RsnInternal, "discovery not init")
		return
	}

	discovery = inst.discovery
	return
}

func newDB(cfg *conf.Data_Database) (db *sql.DB, err error) {
	return sql.Open(cfg.GetDriver(), cfg.GetSource())
}

// new discovery with etcd client
func newDiscovery(cfg *conf.Data_Registrar) (discovery registry.Discovery, client io.Closer, err error) {
	if len(cfg.GetEndpoints()) == 0 {
		return
	}

	_etcd, err := etcdv3.New(etcdv3.Config{Endpoints: cfg.GetEndpoints()})
	if err != nil {
		return
	}
	client = _etcd

	discovery = etcd.New(_etcd)
	return
}

func newGrpcConn(serviceCfg *conf.Data_Service, discovery registry.Discovery) (conn googlegrpc.ClientConnInterface, err error) {
	endpoint := "discovery:///" + serviceCfg.GetName()
	return grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(discovery),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
		grpc.WithTimeout(serviceCfg.GetTimeout().AsDuration()),
	)
}

func newGrpcConnDirect(cfg *conf.Data_Service) (conn googlegrpc.ClientConnInterface, err error) {
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

func newRedis(cfg *conf.Data_Redis) (client *redis.Client) {
	opts := &redis.Options{
		Addr:     cfg.GetAddr(),
		Password: cfg.GetPwd(),
	}

	if cfg.GetTls() {
		opts.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	return redis.NewClient(opts)
}
