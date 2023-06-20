package main

import (
	"flag"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
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
	config := cmd.Config()
	bootstrap := cmd.Bootstrap(config)
	name, version := cmd.Name(Name), cmd.Version(Version)
	logger := cmd.Logger(version)
	log.SetLogger(logger) // set default logger

	app, cleanup, err := wireApp(
		name, version,
		logger, config,
		&bootstrap, bootstrap.Server, bootstrap.Data, bootstrap.Business, bootstrap.Job,
	)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	if err := app.Run(); err != nil {
		panic(err)
	}
}

func NewServers(gs *grpc.Server, hs *http.Server) *cmd.Servers {
	ss := []transport.Server{gs, hs}
	return cmd.NewServers(ss...)
}
