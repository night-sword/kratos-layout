package data

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	googlegrpc "google.golang.org/grpc"

	"github.com/night-sword/kratos-layout/internal/conf"
)

type Data struct {
	config *conf.Data
	db     *sql.DB
}

func NewData(config *conf.Data) (data *Data, cleanup func(), err error) {
	cleanup = func() { log.Info("closing the data resources") }

	data = &Data{}

	return
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
