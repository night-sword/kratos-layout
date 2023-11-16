package data

import (
	"context"
	"database/sql"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/night-sword/kratos-kit/log"
	etcdv3 "go.etcd.io/etcd/client/v3"
	googlegrpc "google.golang.org/grpc"

	"github.com/night-sword/kratos-layout/internal/conf"
	"github.com/night-sword/kratos-layout/internal/dao"
)

type Data struct {
	config *conf.Data
	db     *sql.DB
}

func NewData(cfg *conf.Data) (data *Data, cleanup func(), err error) {
	db := newDB(cfg)

	cleanup = func() {
		log.Info("closing the data resources")

		if e := db.Close(); e != nil {
			log.Error(e)
		}
	}

	data = &Data{
		config: cfg,
		db:     db,
	}
	return
}

func newDB(cfg *conf.Data) (db *sql.DB) {
	db, err := sql.Open(cfg.GetDatabase().GetDriver(), cfg.GetDatabase().GetSource())
	if err != nil {
		panic(err)
	}

	return
}

func newDao(db *sql.DB) (querys *dao.Queries) {
	return dao.New(db)
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
