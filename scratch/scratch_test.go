package main

import (
	"testing"

	"fmt"

	antlr "github.com/wxio/antlr4-go"
	"github.com/wxio/antlr4-go-examples/scratch/parser"
	"github.com/wxio/antlr4-go-examples/scratch/parser2"
)

func TestVisit(t *testing.T) {
	input := antlr.NewInputStream("101x0x01y11111")
	lexer := parser.NewScratchLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewScratchParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	tree := p.Start()
	sexpr := antlr.TreesStringTree(tree, nil, p)
	fmt.Printf("%s\n", sexpr)
	v := &ScratchVisitor{}
	tree.Visit(v)
	fmt.Printf("%s\n", v.SExpr.String())
	fmt.Println()
}

func TestVisit2(t *testing.T) {
	input := antlr.NewInputStream("101c0x01y11111")
	lexer := parser2.NewScratch2Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser2.NewScratch2Parser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	tree := p.Start()
	sexpr := antlr.TreesStringTree(tree, nil, p)
	fmt.Printf("%s\n", sexpr)
	v := &Scratch2Visitor{super: &ScratchVisitor{}}
	tree.Visit(v)
	fmt.Printf("%s\n", v.super.SExpr.String())
	fmt.Println()
}

func TestScratchListener(t *testing.T) {
	input := antlr.NewInputStream("101x0x01y11111")
	lexer := parser.NewScratchLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewScratchParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	p.AddErrorListener(&ScratchErrorListener{})
	tree := p.Start()
	sexpr := antlr.TreesStringTree(tree, nil, p)
	fmt.Printf("%s\n", sexpr)

	l := &ScratchListener{}
	antlr.ParseTreeWalkerDefault.Walk(l, tree)

	fmt.Println()
}

func TestScratch2Listener(t *testing.T) {
	input := antlr.NewInputStream("101c0x01y11111")
	lexer := parser2.NewScratch2Lexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser2.NewScratch2Parser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	p.AddErrorListener(&ScratchErrorListener{})
	tree := p.Start()
	sexpr := antlr.TreesStringTree(tree, nil, p)
	fmt.Printf("%s\n", sexpr)

	l := &Scratch2Listener{s1: &ScratchListener{}}
	antlr.ParseTreeWalkerDefault.Walk(l, tree)

	fmt.Println()
}
