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
	if err := run(os.Args[1:], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}

// run parses args and writes the greeting to out. The flag set runs in
// ContinueOnError mode so tests can observe its results; usage text is written
// to out. Asking for help is not an error: -h prints usage and run returns nil.
func run(args []string, out io.Writer) error {
	fs := flag.NewFlagSet("greet", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.Usage = func() {
		// flag reports parse errors and the usage block through the same
		// writer. Keep the error message out of out (main reports it on
		// stderr) and send only the usage block here.
		fs.SetOutput(out)
		defer fs.SetOutput(io.Discard)
		fmt.Fprintf(out, "Usage of %s:\n", fs.Name())
		fs.PrintDefaults()
	}
	name := fs.String("name", "", "who to greet")
	text := fs.String("greeting", greeting.DefaultGreeting, "greeting text")
	if err := fs.Parse(args); err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return nil
		}
		return err
	}
	_, err := fmt.Fprintln(out, greeting.GreetWith(*text, *name))
	return err
}
