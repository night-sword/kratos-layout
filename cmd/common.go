package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport"
	clientv3 "go.etcd.io/etcd/client/v3"
	_ "go.uber.org/automaxprocs"

	"github.com/night-sword/kratos-layout/internal/conf"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// flagconf is the config flag.
	flagconf string
)

type Name string
type Version string

type Servers struct {
	servers []transport.Server
}

func NewServers(server ...transport.Server) *Servers {
	return &Servers{
		servers: server,
	}
}

func (inst *Servers) Gets() []transport.Server {
	return inst.servers
}

func NewKratos(name Name, version Version, logger log.Logger, servers *Servers, bootstrap *conf.Bootstrap) *kratos.App {
	return kratos.New(
		kratos.ID(id(bootstrap)),
		kratos.Name(string(name)),
		kratos.Version(string(version)),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(servers.Gets()...),

		// kratos.Registrar(registrar(bootstrap)),
	)
}

func id(bootstrap *conf.Bootstrap) string {
	hostname, _ := os.Hostname()
	grpcPort := strings.Split(bootstrap.GetServer().GetGrpc().GetAddr(), ":")[1]

	return fmt.Sprintf("%s:%s", hostname, grpcPort)
}

// new registrar with etcd client
func registrar(bootstrap *conf.Bootstrap) registry.Registrar {
	// new etcd client
	client, err := clientv3.New(clientv3.Config{
		Endpoints: bootstrap.GetData().GetRegistrar().GetEndpoints(),
	})
	if err != nil {
		panic(err)
	}
	// new reg with etcd client
	return etcd.New(client)
}

func Logger(version Version) log.Logger {
	options := []any{
		"ts", log.Timestamp("20060102-150405"),
		"caller", log.DefaultCaller,
		"version", version,
		// "service.id", id,
		// "service.name", Name,
		// "trace.id", tracing.TraceID(),
		// "span.id", tracing.SpanID(),
	}

	return log.With(log.NewStdLogger(os.Stdout), options...)
}

func Config() config.Config {
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	// defer func() { _ = c.Close() }()
	if err := c.Load(); err != nil {
		panic(err)
	}

	return c
}

func Bootstrap(c config.Config) (bootstrap conf.Bootstrap) {
	if err := c.Scan(&bootstrap); err != nil {
		panic(err)
	}
	return
}

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")

	locale, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		locale = time.FixedZone("CST", 8*3600)
	}

	time.Local = locale
}
