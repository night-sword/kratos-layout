package data

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hibiken/asynq"
	clientv3 "go.etcd.io/etcd/client/v3"
	googlegrpc "google.golang.org/grpc"

	"github.com/night-sword/kratos-layout/internal/conf"
	"github.com/night-sword/kratos-layout/internal/dao"
)

type Data struct {
	config *conf.Data
	db     *sql.DB

	queueClient    *asynq.Client
	queueInspector *asynq.Inspector
}

func NewData(config *conf.Data) (data *Data, cleanup func(), err error) {
	cleanup = func() { log.Info("closing the data resources") }

	// TODO init instance
	data = &Data{}

	return
}

func newDB(conf *conf.Data) (db *sql.DB) {
	db, err := sql.Open(conf.Database.GetDriver(), conf.Database.GetSource())
	if err != nil {
		panic(err)
	}

	return
}

func newDao(db *sql.DB) (querys *dao.Queries) {
	return dao.New(db)
}

func newDemoGrpcClient(config *conf.Data_DemoGrpc) (client googlegrpc.ClientConnInterface) {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(config.GetAddr()),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(conn)

	client = &googlegrpc.ClientConn{}
	return
}

func newDemoGrpcClientWithDiscovery(config *conf.Data) (client googlegrpc.ClientConnInterface) {
	endpoint := "discovery:///" + config.GetDemoGrpc().GetName()

	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(endpoint),
		grpc.WithDiscovery(discovery(config)),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}

	fmt.Println(conn)

	client = &googlegrpc.ClientConn{}
	return
}

// new discovery with etcd client
func discovery(config *conf.Data) (discovery *etcd.Registry) {
	// new etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: config.GetRegistrar().GetEndpoints(),
	})
	if err != nil {
		panic(err)
	}

	// new discovery with etcd client
	discovery = etcd.New(client)
	return
}

func newQueueClient(conf *conf.Data) (client *asynq.Client) {
	redisCfg := conf.GetRedis()
	redisOpt := asynq.RedisClientOpt{
		Addr:     redisCfg.GetAddr(),
		Password: redisCfg.GetPwd(),
	}

	client = asynq.NewClient(redisOpt)
	return
}

func newQueueInspector(conf *conf.Data) (inspector *asynq.Inspector) {
	redisCfg := conf.GetRedis()
	redisOpt := asynq.RedisClientOpt{
		Addr:     redisCfg.GetAddr(),
		Password: redisCfg.GetPwd(),
	}

	inspector = asynq.NewInspector(redisOpt)
	return
}
