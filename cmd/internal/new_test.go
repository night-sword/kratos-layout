package internal

import (
	"os"
	"strings"
	"testing"

	"github.com/night-sword/kratos-layout/internal/conf"
)

func Test_id(t *testing.T) {
	hostname, _ := os.Hostname()

	tests := []struct {
		name string
		addr string
		want string
	}{
		{
			name: "standard addr",
			addr: "0.0.0.0:9000",
			want: hostname + ":9000",
		},
		{
			name: "no colon",
			addr: "9000",
			want: hostname + ":9000",
		},
		{
			name: "empty addr",
			addr: "",
			want: hostname + ":",
		},
		{
			name: "ipv6 style with multiple colons",
			addr: "::1:9000",
			want: hostname + ":9000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bootstrap := &conf.Bootstrap{
				Server: &conf.Server{
					Grpc: &conf.Server_GRPC{
						Addr: tt.addr,
					},
				},
			}

			got := id(bootstrap)
			if !strings.HasSuffix(got, tt.want[len(hostname):]) {
				t.Errorf("id() = %q, want suffix %q", got, tt.want[len(hostname):])
			}
			if got != tt.want {
				t.Errorf("id() = %q, want %q", got, tt.want)
			}
		})
	}
}
