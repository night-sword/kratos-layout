package main

import (
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	_ "go.uber.org/automaxprocs"

	cmd "github.com/night-sword/kratos-layout/cmd/internal"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string
	// Version is the version of the compiled software.
	Version string
)

func main() {
	cmd.Run(&Name, Version, wireApp)
}

func newServers(gs *grpc.Server, hs *http.Server) *cmd.Servers {
	ss := []transport.Server{gs, hs}
	return cmd.NewServers(ss...)
}
