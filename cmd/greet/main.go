package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/maehld/factory-sandbox/greeting"
)

func main() {
	if err := run(os.Args[1:], os.Stdout, os.Stderr); err != nil {
		// flag.ErrHelp means usage was already printed on request; keep the
		// help path exiting successfully and reserve a non-zero status for
		// real failures.
		if errors.Is(err, flag.ErrHelp) {
			return
		}
		os.Exit(1)
	}
}

// run greets the given name repeat times, writing to stdout. Usage and error
// messages go to stderr so stdout carries greetings only.
func run(args []string, stdout, stderr io.Writer) error {
	fs := flag.NewFlagSet("greet", flag.ContinueOnError)
	fs.SetOutput(stderr)
	name := fs.String("name", "", "who to greet")
	repeat := fs.Int("repeat", 1, "how many times to print the greeting")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *repeat < 1 {
		err := fmt.Errorf("greet: -repeat must be at least 1, got %d", *repeat)
		fmt.Fprintln(stderr, err)
		return err
	}
	msg := greeting.Greet(*name)
	for i := 0; i < *repeat; i++ {
		fmt.Fprintln(stdout, msg)
	}
	return nil
}
