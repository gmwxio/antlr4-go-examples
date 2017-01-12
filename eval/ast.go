package eval

import (
	"bytes"
	"fmt"

	antlr "github.com/wxio/antlr4-go"
	"github.com/wxio/antlr4-go-examples/ctree"
	"github.com/wxio/antlr4-go-examples/eval/parser"
	. "github.com/wxio/antlr4-go-examples/eval/walker"
)

func LexDebug(str string) string {
	errListener := &errorListener{}
	input := antlr.NewInputStream(str)
	lexer := parser.NewExprLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errListener)

	var buf bytes.Buffer
	i := 0
	for {
		i++
		if i == 100 {
			return "more than 100 tokens"
		}
		t := lexer.NextToken()
		// fmt.Printf("%v\n", t)
		buf.WriteString(fmt.Sprintf("%v\n", t))
		if t.GetTokenType() == antlr.TokenEOF {
			return buf.String()
		}
	}
}

func BuildEvalAST(exprRep string) (ctree.Tree, error) {
	errListener := &errorListener{}
	tbl := &ExprBuildListener{}

	input := antlr.NewInputStream(exprRep)
	lexer := parser.NewExprLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errListener)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewExprParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	p.AddErrorListener(tbl)
	p.BuildParseTrees = true
	tree := p.Startexpr()

	antlr.ParseTreeWalkerDefault.Walk(tbl, tree)
	if tbl.err != "" {
		return nil, fmt.Errorf("ERROR:%v", tbl.err)
	}
	return tbl.Builder.Build(), nil
}

type TTType struct{}

func (*TTType) Eof() int  { return ExprWalkerEOF }
func (*TTType) Down() int { return ExprWalkerDOWN }
func (*TTType) Up() int   { return ExprWalkerUP }

func WalkTree(tree ctree.Tree) (Atom, error) {
	var tttype *TTType
	var tts antlr.TokenStream = ctree.NewTreeTokenSource(tree, tttype)
	p := NewExprWalker(tts)
	p.RemoveErrorListeners()
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tbl := &WalkListener{}
	p.AddErrorListener(tbl)
	p.BuildParseTrees = false
	ret := p.Start()
	if tbl.err != nil {
		return nil, tbl.err
	}
	return ret.GetAtom(), nil
}

type WalkListener struct {
	debug bool
	err   error
}

func (tbl *WalkListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	if tbl.debug {
		fmt.Printf("SyntaxError %d:%d <%s>\n", line, column, msg)
	}
	tbl.err = fmt.Errorf("SyntaxError %d:%d <%s>", line, column, msg)
	// panic("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg)
}

func (tbl *WalkListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// panic("ReportAmbiguity")
	if tbl.debug {
		fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
	}
}

func (tbl *WalkListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// panic("ReportAttemptingFullContext")
	if tbl.debug {
		fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
	}
}

func (tbl *WalkListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	//	panic("ReportContextSensitivity")
}
