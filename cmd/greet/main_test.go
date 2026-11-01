package main

import (
	"bytes"
	"errors"
	"flag"
	"strings"
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

func TestRunHelp(t *testing.T) {
	var out bytes.Buffer
	if err := run([]string{"-h"}, &out); err != nil {
		t.Fatalf("run(-h) returned error: %v", err)
	}
	if got := out.String(); !strings.Contains(got, "Usage of greet:") || !strings.Contains(got, "-greeting") {
		t.Errorf("run(-h) printed %q, want usage text listing the flags", got)
	}
}

func TestRunRejectsUnknownFlag(t *testing.T) {
	var out bytes.Buffer
	err := run([]string{"-nope"}, &out)
	if err == nil {
		t.Fatal("run with unknown flag returned nil error")
	}
	if errors.Is(err, flag.ErrHelp) {
		t.Fatalf("run with unknown flag returned %v, want a parse error other than flag.ErrHelp", err)
	}
	if got := out.String(); !strings.Contains(got, "Usage of greet:") {
		t.Errorf("run with unknown flag printed %q, want usage text", got)
	}
}
