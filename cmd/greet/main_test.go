package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{"default greeting", []string{"-name", "Ada"}, "Hello, Ada!\n"},
		{"custom greeting", []string{"-name", "Ada", "-greeting", "Hi"}, "Hi, Ada!\n"},
		{"empty greeting falls back", []string{"-name", "Ada", "-greeting", ""}, "Hello, Ada!\n"},
		{"no name", nil, "Hello, world!\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var out bytes.Buffer
			if err := run(tt.args, &out); err != nil {
				t.Fatalf("run(%q) returned error: %v", tt.args, err)
			}
			if got := out.String(); got != tt.want {
				t.Errorf("run(%q) printed %q, want %q", tt.args, got, tt.want)
			}
		})
	}
}

func TestRunRejectsUnknownFlag(t *testing.T) {
	var out bytes.Buffer
	if err := run([]string{"-nope"}, &out); err == nil {
		t.Fatal("run with unknown flag returned nil error")
	}
}
