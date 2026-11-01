// Package greeting builds short, human-readable greetings.
package greeting

import (
	"fmt"
	"strings"
)

// Greet returns a greeting for name. An empty or whitespace-only name yields a
// generic greeting instead of a blank one.
func Greet(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		return "Hello, world!"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
