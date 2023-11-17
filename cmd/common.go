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
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	klog "github.com/night-sword/kratos-kit/log"
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
	opts := []kratos.Option{
		kratos.ID(id(bootstrap)),
		kratos.Name(string(name)),
		kratos.Version(string(version)),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(servers.Gets()...),
	}

	// only grpc or http server registrar
	for _, server := range servers.Gets() {
		switch server.(type) {
		case *grpc.Server, *http.Server:
			opts = append(opts, kratos.Registrar(registrar(bootstrap)))
		default:
			continue
		}
	}

	return kratos.New(opts...)
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

func Logger(version string, level string) log.Logger {
	return klog.NewLogger(level, nil, []any{"VER", version})
}

func Config() (cfg config.Config, cancel func()) {
	cfg = config.New(
		config.WithSource(file.NewSource(flagconf)),
	)
	cancel = func() { _ = cfg.Close() }

	if err := cfg.Load(); err != nil {
		panic(err)
	}
	return
}

func Bootstrap(c config.Config) (bootstrap conf.Bootstrap) {
	if err := c.Scan(&bootstrap); err != nil {
		panic(err)
	}
	return
}

func init() {
	flag.StringVar(&flagconf, "conf", "./configs/config.yaml", "config path, eg: -conf config.yaml")

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		location = time.FixedZone("CST", 8*3600)
	}

	time.Local = location
}
