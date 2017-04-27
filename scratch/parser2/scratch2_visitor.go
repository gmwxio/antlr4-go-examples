// Generated from Scratch2.g4 by ANTLR 4.7.

package parser2 // Scratch2
import "github.com/wxio/antlr4-go"

// A complete Visitor for a parse tree produced by Scratch2Parser.
type Scratch2Visitor interface {
	antlr.ParseTreeVisitor
	Start1ContextVisitor
	A1ContextVisitor
	B1ContextVisitor
	B2ContextVisitor
	C1ContextVisitor
}

type Start1ContextVisitor interface {
	VisitStart1(ctx IStart1Context, delegate antlr.ParseTreeVisitor)
}
type A1ContextVisitor interface {
	VisitA1(ctx IA1Context, delegate antlr.ParseTreeVisitor)
}
type B1ContextVisitor interface {
	VisitB1(ctx IB1Context, delegate antlr.ParseTreeVisitor)
}
type B2ContextVisitor interface {
	VisitB2(ctx IB2Context, delegate antlr.ParseTreeVisitor)
}
type C1ContextVisitor interface {
	VisitC1(ctx IC1Context, delegate antlr.ParseTreeVisitor)
}
