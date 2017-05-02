// Generated from Scratch.g4 by ANTLR 4.7.

package parser // Scratch
// See base listener file for example implementations

//import "github.com/wxio/antlr4-go"


// import "generate package"

//type ScratchVisitor struct {
//    *antlr.BaseParseTreeVisitor
//    indent string
//}

//var _ antlr.ParseTreeVisitor = &&ScratchVisitor{}
//var _ antlr.AggregateResultVisitor = &&ScratchVisitor{}
//var _ antlr.VisitNextCheck = &&ScratchVisitor{}
//var _ antlr.VisitRestCheck = &&ScratchVisitor{}
//var _ antlr.EnterEveryRuleVisitor = &&ScratchVisitor{}
//var _ antlr.ExitEveryRuleVisitor = &&ScratchVisitor{}

//var _ parser.Start1ContextVisitor = &ScratchVisitor{}
//var _ parser.A1ContextVisitor = &ScratchVisitor{}
//var _ parser.B1ContextVisitor = &ScratchVisitor{}


//func (v *ScratchVisitor) VisitNext(node antlr.Tree, current interface{}) bool {
//    return true
//}
//func (v *ScratchVisitor) VisitRest(node antlr.RuleNode, current interface{}) bool {
//    return true
//}
//func (v *ScratchVisitor) AggregateResult(aggregate, nextResult interface{}) (result interface{}) {
//    return nextResult
//}
//func (v *ScratchVisitor) VisitTerminal(node antlr.TerminalNode) {
//}
//func (v *ScratchVisitor) VisitErrorNode(node antlr.ErrorNode) {
//}
//func (v *ScratchVisitor) EnterEveryRule(ctx antlr.RuleNode) {
//    v.indent += " "
//    fmt.Printf("%s %T\n", v.indent, ctx)
//}
//func (v *ScratchVisitor) ExitEveryRule(ctx antlr.RuleNode) {
//    v.indent = v.indent[:len(v.indent)-1]
//}


//func (v *ScratchVisitor) VisitStart1(ctx parser.IStart1Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    v.VisitChildren(ctx, delegate);return }
//func (v *ScratchVisitor) VisitA1(ctx parser.IA1Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    v.VisitChildren(ctx, delegate);return }
//func (v *ScratchVisitor) VisitB1(ctx parser.IB1Context, delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}){    v.VisitChildren(ctx, delegate);return }


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

