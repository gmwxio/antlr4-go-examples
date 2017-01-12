# Antlr4 Go Examples

## Getting started

```
mkdir antlr4test
cd antlr4test
mkdir src
export GOPATH=`pwd`
go get github.com/wxio/antlr4-go
go get github.com/wxio/antlr4-go-examples
cd src/github.com/wxio/antlr4-go-examples
go generate ./...
go test ./...
```

The `go generate` command read the source and executes commands specified by in `//go:generate xxx` lines.
The example has such a line, which calls the Antlr4 tool and then `sed` to replace the import statement in the generated code.

## Issues
This only works on OSX as the `sed` for OSX is different.
To get this to work in Linux or Windows, change the sed command or edit the creating and modifying a generate file (eg. gen_linux.go).
