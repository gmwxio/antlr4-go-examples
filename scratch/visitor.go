package main

import (
	"fmt"

	"bytes"

	antlr "github.com/wxio/antlr4-go"
	"github.com/wxio/antlr4-go-examples/scratch/parser"
)

type ScratchVisitor struct {
	*antlr.BaseParseTreeVisitor
	// fields to control the interation
	lastAis1 bool
	seenY    bool

	// result fields
	SExpr bytes.Buffer
}

// Can be used to check that all Visit methods are "overriden"
//var _ parser.ScratchVisitor = &ScratchVisitor{}

// Filter
func (v *ScratchVisitor) VisitNext(node antlr.Tree) bool {
	// Don't visit terminal node
	if _, ok := node.(antlr.TerminalNode); ok {
		return false
	}
	// Don't visit rule B is the last rule A token was a 0
	if _, ok := node.(*parser.BContext); !v.lastAis1 && ok {
		return false
	}
	return true
}

// never called because VisitNext filters all TerminalNode types
func (v *ScratchVisitor) VisitTerminal(node antlr.TerminalNode) {
	fmt.Printf("terminal %v\n", node.GetText())
}
func (v *ScratchVisitor) VisitErrorNode(node antlr.ErrorNode) {
}

func (v *ScratchVisitor) VisitA1(ctx parser.IA1Context, delegate antlr.ParseTreeVisitor) {
	// before children
	// Don't visit any more children after a "y"
	if v.seenY {
		return
	}
	v.SExpr.WriteString(" (a ")
	v.SExpr.WriteString(ctx.GetT().GetText())
	fmt.Printf("%s\n", ctx.GetT().GetText())
	v.lastAis1 = ctx.GetT().GetText() == "1"
	v.VisitChildren(ctx, delegate)
	// afer children
	v.SExpr.WriteString(")")
}

func (v *ScratchVisitor) VisitB1(ctx parser.IB1Context, delegate antlr.ParseTreeVisitor) {
	// before children
	v.SExpr.WriteString(" (b ")
	v.SExpr.WriteString(ctx.GetT().GetText())
	fmt.Printf("%s\n", ctx.GetT().GetText())
	v.seenY = ctx.GetT().GetText() == "y"
	v.VisitChildren(ctx, delegate)
	// afer children
	v.SExpr.WriteString(")")
}
