package data

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	_ "github.com/go-sql-driver/mysql"
	"github.com/night-sword/kratos-kit/log"
	googlegrpc "google.golang.org/grpc"

	"github.com/night-sword/kratos-layout/internal/conf"
	"github.com/night-sword/kratos-layout/internal/dao"
)

type Data struct {
	config *conf.Data
	db     *sql.DB
}

func NewData(config *conf.Data) (data *Data, cleanup func(), err error) {
	cleanup = func() {
		// TODO: close resources
		log.Info("closing the data resources")
	}

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
