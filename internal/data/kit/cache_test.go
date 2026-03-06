package kit

import "testing"

func TestCache_Key(t *testing.T) {
	tests := []struct {
		name   string
		prefix string
		key    string
		want   string
	}{
		{
			name:   "normal",
			prefix: "myapp",
			key:    "user:1",
			want:   "myapp:user:1",
		},
		{
			name:   "empty prefix",
			prefix: "",
			key:    "user:1",
			want:   ":user:1",
		},
		{
			name:   "empty key",
			prefix: "myapp",
			key:    "",
			want:   "myapp:",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Cache{prefix: tt.prefix}
			if got := c.Key(tt.key); got != tt.want {
				t.Errorf("Key() = %q, want %q", got, tt.want)
			}
		})
	}
}
