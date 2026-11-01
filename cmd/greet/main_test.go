package main

import (
	"bytes"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		wantOut  string
		wantCode int
	}{
		{"default greeting", nil, "Hello, world!\n", 0},
		{"named greeting", []string{"-name", "Ada"}, "Hello, Ada!\n", 0},
		{"version", []string{"-version"}, version + "\n", 0},
		{"version wins over name", []string{"-version", "-name", "Ada"}, version + "\n", 0},
		{"help short flag", []string{"-h"}, "", 0},
		{"help long flag", []string{"-help"}, "", 0},
		{"unknown flag", []string{"-nope"}, "", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var stdout, stderr bytes.Buffer
			gotCode := run(tt.args, &stdout, &stderr)
			if gotCode != tt.wantCode {
				t.Errorf("run(%q) = exit code %d, want %d", tt.args, gotCode, tt.wantCode)
			}
			if got := stdout.String(); got != tt.wantOut {
				t.Errorf("run(%q) stdout = %q, want %q", tt.args, got, tt.wantOut)
			}
			if tt.wantCode != 0 && stderr.Len() == 0 {
				t.Errorf("run(%q) wrote nothing to stderr, want an error message", tt.args)
			}
		})
	}
}
