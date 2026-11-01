# factory-sandbox

A small Go project used to exercise the [software
factory](https://github.com/maehld/software-factory) end to end.

## Layout

```
greeting/       library: Greet builds a greeting from a name
cmd/greet/      tiny CLI around the library
```

## Commands

```sh
go build ./...
go vet ./...
go test ./...
go run ./cmd/greet -name Ada
```

The factory's verify stage runs `go build ./...`, `go vet ./...`, and
`go test ./...`. Keep those green.
