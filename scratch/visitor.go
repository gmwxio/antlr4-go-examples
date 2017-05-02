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

var _ antlr.ParseTreeVisitor = &ScratchVisitor{}

// var _ antlr.AggregateResultVisitor = &ScratchVisitor{}
var _ antlr.VisitNextCheck = &ScratchVisitor{}
var _ antlr.VisitRestCheck = &ScratchVisitor{}

// var _ antlr.EnterEveryRuleVisitor = &ScratchVisitor{}
// var _ antlr.ExitEveryRuleVisitor = &ScratchVisitor{}

// var _ parser.Start1ContextVisitor = &ScratchVisitor{}
var _ parser.A1ContextVisitor = &ScratchVisitor{}
var _ parser.B1ContextVisitor = &ScratchVisitor{}

// Filter
func (v *ScratchVisitor) VisitNext(node antlr.Tree, current interface{}) bool {
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
func (v *ScratchVisitor) VisitRest(node antlr.RuleNode, current interface{}) bool {
	if v.seenY {
		v.seenY = true
		return false
	}
	return true
}

// never called because VisitNext filters all TerminalNode types
func (v *ScratchVisitor) VisitTerminal(node antlr.TerminalNode) {
	fmt.Printf("terminal %v\n", node.GetText())
	return
}
func (v *ScratchVisitor) VisitErrorNode(node antlr.ErrorNode) {
	fmt.Printf("ErrorNode %v\n", node)
	return
}

func (v *ScratchVisitor) VisitA1(ctx parser.IA1Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	// before children
	// Don't visit any more children after a "y"
	// if v.seenY {
	// 	return
	// }
	v.SExpr.WriteString(" (a ")
	v.SExpr.WriteString(fmt.Sprintf("%v", ctx.GetT().GetText()))
	fmt.Printf("%s\n", ctx.GetT().GetText())
	// v.lastAis1 = ctx.GetT().GetText() == "1"
	v.VisitChildren(ctx, delegate)
	// afer children
	v.SExpr.WriteString(")")
	return
}

func (v *ScratchVisitor) VisitB1(ctx parser.IB1Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	// before children
	v.SExpr.WriteString(" (b ")
	v.SExpr.WriteString(ctx.GetT().GetText())
	fmt.Printf("%s\n", ctx.GetT().GetText())
	v.seenY = ctx.GetT().GetText() == "y"
	v.VisitChildren(ctx, delegate)
	// afer children
	v.SExpr.WriteString(")")
	return
}
