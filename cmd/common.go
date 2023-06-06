package cmd

import (
	"flag"
	"os"
	"time"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
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

func NewKratos(name Name, version Version, logger log.Logger, servers *Servers) *kratos.App {
	id, _ := os.Hostname()

	return kratos.New(
		kratos.ID(id),
		kratos.Name(string(name)),
		kratos.Version(string(version)),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(servers.Gets()...),
	)
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

func Bootstrap() (bootstrap conf.Bootstrap) {
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer func() { _ = c.Close() }()
	if err := c.Load(); err != nil {
		panic(err)
	}

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
