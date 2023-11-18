package internal

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	klog "github.com/night-sword/kratos-kit/log"
	etcdv3 "go.etcd.io/etcd/client/v3"
	_ "go.uber.org/automaxprocs"

	"github.com/night-sword/kratos-layout/internal/conf"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// flagconf is the config flag.
	flagconf string
)

func NewKratos(opts []kratos.Option, servers *Servers, bootstrap *conf.Bootstrap) *kratos.App {
	options := []kratos.Option{
		kratos.Server(servers.Gets()...),
		kratos.Metadata(map[string]string{}),
		kratos.ID(id(bootstrap)),
	}
	options = append(options, opts...)

	if len(bootstrap.GetData().GetRegistrar().GetEndpoints()) > 0 {
		// only grpc or http server registrar
		for _, server := range servers.Gets() {
			switch server.(type) {
			case *grpc.Server, *http.Server:
				options = append(options, kratos.Registrar(registrar(bootstrap)))
			default:
				continue
			}
		}
	}

	return kratos.New(options...)
}

func Bootstrap() (bootstrap *conf.Bootstrap, cleanup func()) {
	c, cleanup := cfg()

	bootstrap = &conf.Bootstrap{}
	if err := c.Scan(bootstrap); err != nil {
		panic(err)
	}

	fn := func(key string, value config.Value) {
		log.Info("business config changed")

		// only auto refresh business config
		b := &conf.Bootstrap{}
		if err := c.Scan(b); err != nil {
			log.Errorw("scan fail, skip update bootstrap", err)
			return
		}
		bootstrap.Business = b.GetBusiness()

		log.Info("bootstrap.Business refresh success")
	}
	if err := c.Watch("business", fn); err != nil {
		panic(err)
	}

	return
}

func Logger(version string, level string) (logger log.Logger) {
	logger = klog.NewLogger(level, nil, []any{"VER", version})
	klog.SetLogger(logger)
	log.SetLogger(logger)
	return
}

func id(bootstrap *conf.Bootstrap) string {
	hostname, _ := os.Hostname()
	grpcPort := strings.Split(bootstrap.GetServer().GetGrpc().GetAddr(), ":")[1]

	return fmt.Sprintf("%s:%s", hostname, grpcPort)
}

// new registrar with etcd client
func registrar(bootstrap *conf.Bootstrap) registry.Registrar {
	// new etcd client
	client, err := etcdv3.New(etcdv3.Config{
		Endpoints: bootstrap.GetData().GetRegistrar().GetEndpoints(),
	})
	if err != nil {
		panic(err)
	}
	// new reg with etcd client
	return etcd.New(client)
}

func cfg() (cfg config.Config, cancel func()) {
	cfg = config.New(
		config.WithSource(file.NewSource(flagconf)),
	)
	cancel = func() { _ = cfg.Close() }

	if err := cfg.Load(); err != nil {
		panic(err)
	}
	return
}

func init() {
	flag.StringVar(&flagconf, "conf", "./configs/config.yaml", "config path, eg: -conf config.yaml")
}
