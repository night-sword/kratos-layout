package data

import (
	"context"
	"crypto/tls"
	"database/sql"
	"fmt"

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

func NewData(cfg *conf.Bootstrap) (data *Data, cleanup func(), err error) {
	db := newDB(cfg.GetData())
	rds := newRedis(cfg.GetData())

	cleanup = func() {
		log.Info("closing the data resources")

		if e := db.Close(); e != nil {
			log.Error(e)
		}
		if e := rds.Close(); e != nil {
			log.Error(e)
		}
	}

	data = &Data{
		cfg:   cfg,
		db:    db,
		redis: rds,
	}

	return
}

func (inst *Data) cacheKey(key string) string {
	return fmt.Sprintf("%s:%s", inst.cfg.GetBusiness().GetName(), key)
}

func newDB(config *conf.Data) (db *sql.DB) {
	if config.GetDatabase().GetSource() == "" {
		panic("database source not config, if do not need db, pls remove this fn call")
	}

	db, err := sql.Open(config.GetDatabase().GetDriver(), config.GetDatabase().GetSource())
	if err != nil {
		panic(err)
	}
	return
}

// new etcd client
func newEtcdClient(config *conf.Data) (client *etcdv3.Client) {
	client, err := etcdv3.New(etcdv3.Config{
		Endpoints: config.GetRegistrar().GetEndpoints(),
	})
	if err != nil {
		panic(err)
	}
	return
}

// new discovery with etcd client
func newDiscovery(client *etcdv3.Client) (discovery *etcd.Registry) {
	return etcd.New(client)
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

func newRedis(config *conf.Data) (cache *redis.Client) {
	if config.GetRedis().GetAddr() == "" {
		panic("redis addr not config, if do not need redis, pls remove this fn call")
	}

	opts := &redis.Options{
		Addr:     config.GetRedis().GetAddr(),
		Password: config.GetRedis().GetPwd(),
	}

	if config.GetRedis().GetTls() {
		opts.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	return redis.NewClient(opts)
}
