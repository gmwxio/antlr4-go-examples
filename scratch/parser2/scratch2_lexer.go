// Generated from Scratch2.g4 by ANTLR 4.7.

package parser2
import (
	"fmt"
	"unicode"
)

import "github.com/wxio/antlr4-go"
// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter


var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 8, 27, 8, 
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 
	7, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 
	7, 2, 2, 8, 3, 3, 5, 4, 7, 5, 9, 6, 11, 7, 13, 8, 3, 2, 2, 2, 26, 2, 3, 
	3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 
	3, 2, 2, 2, 2, 13, 3, 2, 2, 2, 3, 15, 3, 2, 2, 2, 5, 17, 3, 2, 2, 2, 7, 
	19, 3, 2, 2, 2, 9, 21, 3, 2, 2, 2, 11, 23, 3, 2, 2, 2, 13, 25, 3, 2, 2, 
	2, 15, 16, 7, 51, 2, 2, 16, 4, 3, 2, 2, 2, 17, 18, 7, 50, 2, 2, 18, 6, 
	3, 2, 2, 2, 19, 20, 7, 122, 2, 2, 20, 8, 3, 2, 2, 2, 21, 22, 7, 123, 2, 
	2, 22, 10, 3, 2, 2, 2, 23, 24, 7, 101, 2, 2, 24, 12, 3, 2, 2, 2, 25, 26, 
	7, 102, 2, 2, 26, 14, 3, 2, 2, 2, 3, 2, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'1'", "'0'", "'x'", "'y'", "'c'", "'d'",
}

var lexerSymbolicNames []string


var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5",
}

type Scratch2Lexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewScratch2Lexer(input antlr.CharStream) *Scratch2Lexer {

	l := new(Scratch2Lexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Scratch2.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// Scratch2Lexer tokens.
const (
	Scratch2LexerT__0 = 1
	Scratch2LexerT__1 = 2
	Scratch2LexerT__2 = 3
	Scratch2LexerT__3 = 4
	Scratch2LexerT__4 = 5
	Scratch2LexerT__5 = 6
)

