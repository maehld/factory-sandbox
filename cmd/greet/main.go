package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/maehld/factory-sandbox/greeting"
)

func main() {
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}

func run(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("greet", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	name := fs.String("name", "", "who to greet")
	text := fs.String("greeting", greeting.DefaultGreeting, "greeting text")
	if err := fs.Parse(args); err != nil {
		return err
	}
	_, err := fmt.Fprintln(out, greeting.GreetWith(*text, *name))
	return err
}
