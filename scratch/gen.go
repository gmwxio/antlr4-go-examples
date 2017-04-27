package main

//go:generate java -jar $GOPATH/src/github.com/wxio/antlr4-go/lib/wxio/antlr4-4.7.1-SNAPSHOT-complete.jar Scratch.g4 -o parser -visitor

//go:generate java -jar $GOPATH/src/github.com/wxio/antlr4-go/lib/wxio/antlr4-4.7.1-SNAPSHOT-complete.jar Scratch2.g4 -o parser2 -package parser2 -visitor

func main() {

}
