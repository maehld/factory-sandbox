// Package greeting builds short, human-readable greetings.
package greeting

import (
	"fmt"
	"strings"
)

// DefaultGreeting is the greeting used when none is supplied.
const DefaultGreeting = "Hello"

// Greet returns a greeting for name. An empty or whitespace-only name yields a
// generic greeting instead of a blank one.
func Greet(name string) string {
	return GreetWith(DefaultGreeting, name)
}

// GreetWith returns text and name combined as a greeting, for example
// "Hi, Ada!". A blank text falls back to DefaultGreeting. An empty or
// whitespace-only name yields a generic greeting instead of a blank one.
func GreetWith(text, name string) string {
	text = strings.TrimSpace(text)
	if text == "" {
		text = DefaultGreeting
	}
	name = strings.TrimSpace(name)
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("%s, %s!", text, name)
}
