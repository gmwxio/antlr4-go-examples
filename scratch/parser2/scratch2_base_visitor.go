// Generated from Scratch2.g4 by ANTLR 4.7.

package parser2 // Scratch2
// See base listener file for example implementations

//import "github.com/wxio/antlr4-go"


// import "generate package"

//type Scratch2Visitor struct {
//    *antlr.BaseParseTreeVisitor
//    indent string
//}

//var _ antlr.ParseTreeVisitor = &&Scratch2Visitor{}
//var _ antlr.AggregateResultVisitor = &&Scratch2Visitor{}
//var _ antlr.VisitNextCheck = &&Scratch2Visitor{}
//var _ antlr.VisitRestCheck = &&Scratch2Visitor{}
//var _ antlr.EnterEveryRuleVisitor = &&Scratch2Visitor{}
//var _ antlr.ExitEveryRuleVisitor = &&Scratch2Visitor{}

//var _ parser2.Start1ContextVisitor = &Scratch2Visitor{}
//var _ parser2.A1ContextVisitor = &Scratch2Visitor{}
//var _ parser2.B1ContextVisitor = &Scratch2Visitor{}
//var _ parser2.B2ContextVisitor = &Scratch2Visitor{}
//var _ parser2.C1ContextVisitor = &Scratch2Visitor{}
//var _ parser2.DContextVisitor = &Scratch2Visitor{}


//func (v *Scratch2Visitor) VisitNext(node antlr.Tree, current interface{}) bool {
//    return true
//}
//func (v *Scratch2Visitor) VisitRest(node antlr.RuleNode, current interface{}) bool {
//    return true
//}
//func (v *Scratch2Visitor) AggregateResult(aggregate, nextResult interface{}) (result interface{}) {
//    return nextResult
//}
//func (v *Scratch2Visitor) VisitTerminal(node antlr.TerminalNode) {
//}
//func (v *Scratch2Visitor) VisitErrorNode(node antlr.ErrorNode) {
//}
//func (v *Scratch2Visitor) EnterEveryRule(ctx antlr.RuleNode) {
//    v.indent += " "
//    fmt.Printf("%s %T\n", v.indent, ctx)
//}
//func (v *Scratch2Visitor) ExitEveryRule(ctx antlr.RuleNode) {
//    v.indent = v.indent[:len(v.indent)-1]
//}


//func (v *Scratch2Visitor) VisitStart1(ctx parser2.IStart1Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    v.VisitChildren(ctx, delegate);return }
//func (v *Scratch2Visitor) VisitA1(ctx parser2.IA1Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    v.VisitChildren(ctx, delegate);return }
//func (v *Scratch2Visitor) VisitB1(ctx parser2.IB1Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    v.VisitChildren(ctx, delegate);return }
//func (v *Scratch2Visitor) VisitB2(ctx parser2.IB2Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    v.VisitChildren(ctx, delegate);return }
//func (v *Scratch2Visitor) VisitC1(ctx parser2.IC1Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    v.VisitChildren(ctx, delegate);return }
//func (v *Scratch2Visitor) VisitD(ctx parser2.IDContext, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    v.VisitChildren(ctx, delegate);return }


//  TODO list rules here
//  Visit rules manually
//  eg a : b c* | d;
//  if ctx.GetB() != nil {
//    result1 = ctx.GetB(ctx, delegate, args)
//    for _, c := range ctx.GetC() {
//      resultS = c.GetC(ctx, delegate, args)
//    }
//  } else { ... }
//  OR visit all children rules
//  // before children
//  v.VisitChildren(ctx, delegate)
//  // afer children
//
//  return result

