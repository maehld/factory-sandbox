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

func TestGreetWith(t *testing.T) {
	tests := []struct {
		text string
		name string
		want string
	}{
		{"Hi", "Ada", "Hi, Ada!"},
		{"Hi", "", "Hi, world!"},
		{"  Hi  ", "  Ada  ", "Hi, Ada!"},
		{"", "Ada", "Hello, Ada!"},
		{"   ", "Ada", "Hello, Ada!"},
		{DefaultGreeting, "Ada", "Hello, Ada!"},
	}

	for _, tt := range tests {
		if got := GreetWith(tt.text, tt.name); got != tt.want {
			t.Errorf("GreetWith(%q, %q) = %q, want %q", tt.text, tt.name, got, tt.want)
		}
	}
}
