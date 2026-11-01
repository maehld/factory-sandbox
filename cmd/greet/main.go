package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/maehld/factory-sandbox/greeting"
)

// version is reported by -version. Override at build time with
// go build -ldflags "-X main.version=v1.2.3" ./cmd/greet
var version = "dev"

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

// run parses args and writes the output for a single invocation. It returns
// the process exit code: 0 on success (including -h/-help), 2 on a flag
// parsing error.
func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("greet", flag.ContinueOnError)
	fs.SetOutput(stderr)
	showVersion := fs.Bool("version", false, "print the version and exit")
	name := fs.String("name", "", "who to greet")
	if err := fs.Parse(args); err != nil {
		// -h/-help is not an error: usage has already been written to stderr.
		if errors.Is(err, flag.ErrHelp) {
			return 0
		}
		return 2
	}

	if *showVersion {
		fmt.Fprintln(stdout, version)
		return 0
	}

	fmt.Fprintln(stdout, greeting.Greet(*name))
	return 0
}
