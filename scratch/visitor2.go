package main

import (
	"fmt"

	antlr "github.com/wxio/antlr4-go"
	"github.com/wxio/antlr4-go-examples/scratch/parser"
	"github.com/wxio/antlr4-go-examples/scratch/parser2"
)

type Scratch2Visitor struct {
	*antlr.BaseParseTreeVisitor
	super *ScratchVisitor
}

// func (v *Scratch2Visitor) VisitNext(node antlr.Tree) bool {
// 	return true
// }
// func (v *Scratch2Visitor) VisitTerminal(node antlr.TerminalNode) {
// }
// func (v *Scratch2Visitor) VisitErrorNode(node antlr.ErrorNode) {
// }

func (v *Scratch2Visitor) VisitStart1(ctx parser2.IStart1Context, delegate antlr.ParseTreeVisitor) {
	v.VisitChildren(ctx, delegate)
}

func (v *Scratch2Visitor) VisitA1(ctx parser2.IA1Context, delegate antlr.ParseTreeVisitor) {
	v.super.VisitA1(ctx.(parser.IA1Context), delegate)
}

func (v *Scratch2Visitor) VisitB1(ctx parser2.IB1Context, delegate antlr.ParseTreeVisitor) {
	v.super.VisitB1(ctx.(parser.IB1Context), delegate)
}

func (v *Scratch2Visitor) VisitC1(ctx parser2.IC1Context, delegate antlr.ParseTreeVisitor) {
	fmt.Printf("C\n")
	v.VisitChildren(ctx, delegate)
}
