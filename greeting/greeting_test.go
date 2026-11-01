package greeting

import "testing"

func TestGreet(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Ada", "Hello, Ada!"},
		{"", "Hello, world!"},
		{"   ", "Hello, world!"},
		{"  Grace  ", "Hello, Grace!"},
	}

	for _, tt := range tests {
		if got := Greet(tt.name); got != tt.want {
			t.Errorf("Greet(%q) = %q, want %q", tt.name, got, tt.want)
		}
	}
}
