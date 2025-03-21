package data

import (
	"context"
	"crypto/tls"
	"database/sql"
	"fmt"
	"io"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/night-sword/kratos-kit/log"
	"github.com/redis/go-redis/v9"
	etcdv3 "go.etcd.io/etcd/client/v3"
	googlegrpc "google.golang.org/grpc"

	"github.com/night-sword/kratos-layout/internal/conf"
)

type Data struct {
	cfg *conf.Bootstrap

	db    *sql.DB
	redis *redis.Client
}

func NewData(cfg *conf.Bootstrap) (data *Data, cleanup func()) {
	db := newDB(cfg.GetData().GetDatabase())
	_redis := newRedis(cfg.GetData().GetRedis())

	closer := []io.Closer{
		db, _redis, _redis,
	}

	cleanup = func() {
		log.Info("closing the data resources...")
		for i := range closer {
			log.E(closer[i].Close())
		}
		log.Info("finish close data resources")
	}

	data = &Data{
		cfg:   cfg,
		db:    db,
		redis: _redis,
	}
	return
}

func (inst *Data) cacheKey(key string) string {
	return fmt.Sprintf("%s:%s", inst.cfg.GetBusiness().GetName(), key)
}

func newDB(cfg *conf.Data_Database) (db *sql.DB) {
	db, err := sql.Open(cfg.GetDriver(), cfg.GetSource())
	if err != nil {
		panic(err)
	}
	return
}

// new discovery with etcd client
func newDiscovery(cfg *conf.Data_Registrar) (discovery *etcd.Registry, client *etcdv3.Client, err error) {
	if len(cfg.GetEndpoints()) == 0 {
		return
	}

	ec := etcdv3.Config{Endpoints: cfg.GetEndpoints()}
	if client, err = etcdv3.New(ec); err != nil {
		return
	}

	discovery = etcd.New(client)
	return
}

func newGrpcConn(serviceCfg *conf.Data_Service, discovery *etcd.Registry) (conn googlegrpc.ClientConnInterface) {
	endpoint := "discovery:///" + serviceCfg.GetName()
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(discovery),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
		grpc.WithTimeout(serviceCfg.GetTimeout().AsDuration()),
	)
	if err != nil {
		panic(err)
	}
	return
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
