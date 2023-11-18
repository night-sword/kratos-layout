package internal

import "github.com/go-kratos/kratos/v2/transport"

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
