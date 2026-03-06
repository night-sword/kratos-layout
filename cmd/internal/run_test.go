package internal

import (
	"testing"
	"time"

	"github.com/night-sword/kratos-layout/internal/conf"
)

func Test_fixTimeZone(t *testing.T) {
	originalLocal := time.Local
	defer func() { time.Local = originalLocal }()

	tests := []struct {
		name     string
		cfg      *conf.Data_TimeZone
		wantName string
		changed  bool
	}{
		{
			name:    "nil config",
			cfg:     nil,
			changed: false,
		},
		{
			name:    "fix disabled",
			cfg:     &conf.Data_TimeZone{Location: "Asia/Shanghai", Fix: false},
			changed: false,
		},
		{
			name:    "empty location",
			cfg:     &conf.Data_TimeZone{Location: "", Fix: true},
			changed: false,
		},
		{
			name:     "valid location",
			cfg:      &conf.Data_TimeZone{Location: "Asia/Shanghai", Fix: true},
			wantName: "Asia/Shanghai",
			changed:  true,
		},
		{
			name:     "invalid location fallback to offset",
			cfg:      &conf.Data_TimeZone{Location: "Invalid/Zone", Fix: true, Offset: 8 * 3600},
			wantName: "CST",
			changed:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			time.Local = originalLocal
			before := time.Local

			fixTimeZone(tt.cfg)

			if tt.changed {
				if time.Local.String() != tt.wantName {
					t.Errorf("time.Local = %q, want %q", time.Local.String(), tt.wantName)
				}
			} else {
				if time.Local != before {
					t.Errorf("time.Local changed unexpectedly to %q", time.Local.String())
				}
			}
		})
	}
}
