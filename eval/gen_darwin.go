package eval

//go:generate java -jar $GOPATH/src/github.com/wxio/antlr4-go/lib/antlr4-4.6.1-SNAPSHOT-complete.jar ExprLexer.g4 ExprParser.g4 -o parser -Dlanguage=Go
//go:generate sh -c "(cd parser; sed -i '' -e 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/wxio/antlr4-go!' *.go )"

//go:generate java -jar $GOPATH/src/github.com/wxio/antlr4-go/lib/antlr4-4.6.1-SNAPSHOT-complete.jar -lib ./parser ExprWalker.g4 -o walker -no-listener -package walker -Dlanguage=Go
//go:generate sh -c "(cd walker; sed -i '' -e 's!github.com/antlr/antlr4/runtime/Go/antlr!github.com/wxio/antlr4-go!' *.go )"
