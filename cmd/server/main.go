package main

import (
	"flag"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	klog "github.com/night-sword/kratos-kit/log"
	_ "go.uber.org/automaxprocs"

	"github.com/night-sword/kratos-layout/cmd"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
)

func main() {
	flag.Parse()
	config, cfgCleanup := cmd.Config()
	bootstrap := cmd.Bootstrap(config)

	name, version := cmd.Name(Name), cmd.Version(Version)

	logger := cmd.Logger(Version, bootstrap.GetData().GetLog().GetLevel())
	klog.SetLogger(logger)
	log.SetLogger(logger)

	if bootstrap.GetBusiness().GetName() != "" {
		name = cmd.Name(bootstrap.GetBusiness().GetName())
	}

	app, cleanup, err := wireApp(
		name, version,
		logger, config,
		&bootstrap, bootstrap.Server, bootstrap.Data, bootstrap.Business,
	)

	if err != nil {
		panic(err)
	}

	defer func() {
		cfgCleanup()
		cleanup()
	}()

	if e := app.Run(); e != nil {
		panic(e)
	}
}

func NewServers(gs *grpc.Server, hs *http.Server) *cmd.Servers {
	ss := []transport.Server{gs, hs}
	return cmd.NewServers(ss...)
}
