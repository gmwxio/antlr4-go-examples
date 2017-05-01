// Generated from Scratch.g4 by ANTLR 4.7.

package parser // Scratch
import "github.com/wxio/antlr4-go"
// A complete Visitor for a parse tree produced by ScratchParser.
type ScratchVisitor interface {
    antlr.ParseTreeVisitor
    Start1ContextVisitor
    A1ContextVisitor
    B1ContextVisitor
}

type Start1ContextVisitor interface {
    VisitStart1(ctx IStart1Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type A1ContextVisitor interface {
    VisitA1(ctx IA1Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}
type B1ContextVisitor interface {
    VisitB1(ctx IB1Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{})
}