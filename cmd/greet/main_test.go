package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{"default", []string{"-name", "Ada"}, "Hello, Ada!\n"},
		{"repeat 3", []string{"-name", "Ada", "-repeat", "3"}, "Hello, Ada!\nHello, Ada!\nHello, Ada!\n"},
		{"repeat equals form", []string{"-name", "Ada", "-repeat=2"}, "Hello, Ada!\nHello, Ada!\n"},
		{"repeat with empty name", []string{"-name", "  ", "-repeat", "2"}, "Hello, world!\nHello, world!\n"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			if err := run(tt.args, &stdout, &stderr); err != nil {
				t.Fatalf("run(%q) returned error: %v", tt.args, err)
			}
			if got := stdout.String(); got != tt.want {
				t.Errorf("run(%q) stdout = %q, want %q", tt.args, got, tt.want)
			}
			if got := stderr.String(); got != "" {
				t.Errorf("run(%q) stderr = %q, want empty", tt.args, got)
			}
		})
	}
}

func TestRunRejectsNonPositiveRepeat(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{"zero", "0"},
		{"negative", "-1"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			err := run([]string{"-name", "Ada", "-repeat", tt.value}, &stdout, &stderr)
			if err == nil {
				t.Fatalf("run with -repeat %s returned nil error", tt.value)
			}
			if stdout.Len() != 0 {
				t.Errorf("run with -repeat %s wrote %q to stdout, want nothing", tt.value, stdout.String())
			}
			if stderr.Len() == 0 {
				t.Errorf("run with -repeat %s wrote nothing to stderr, want a message", tt.value)
			}
			if !strings.Contains(stderr.String(), "repeat") {
				t.Errorf("run with -repeat %s stderr = %q, want it to mention -repeat", tt.value, stderr.String())
			}
		})
	}
}

func TestRunFlagParseError(t *testing.T) {
	var stdout, stderr bytes.Buffer
	if err := run([]string{"-nope"}, &stdout, &stderr); err == nil {
		t.Fatal("run with an unknown flag returned nil error")
	}
	if stdout.Len() != 0 {
		t.Errorf("run with an unknown flag wrote %q to stdout, want nothing", stdout.String())
	}
	if stderr.Len() == 0 {
		t.Error("run with an unknown flag wrote nothing to stderr, want usage text")
	}
}
