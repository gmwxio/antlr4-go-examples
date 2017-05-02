// Generated from Scratch2.g4 by ANTLR 4.7.

package parser2 // Scratch2
//import "github.com/wxio/antlr4-go"
//import "generated code if in another package"

//// Commented out basic implementation for your convenience.

//func Example(s string) {
//  // Setup
//  input := antlr.NewInputStream(s)
//  lexer := parser2.NewScratch2Lexer(input)
//  stream := antlr.NewCommonTokenStream(lexer, 0)
//  p := parser2.NewScratch2Parser(stream)

//  // Antlr error listener - turns reports (ambiguity etc) into syntax errors
//  p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

//  // Custom error listener, register before the parse
//  el := &Scratch2ErrorListener{}
//  p.AddErrorListener(el)

//  // Parse - start rule
//  tree := p.Start()

//  // Antlr provided parse tree representation
//  sexpr := antlr.TreesStringTree(tree, nil, p)
//  fmt.Printf("%s\n", sexpr)

//  // Custom listener
//  l := &Scratch2Listener{}
//  antlr.ParseTreeWalkerDefault.Walk(l, tree)

//  // Custom visitor
//  v := &Scratch2Visitor{}
//  tree.Accept(v)
// }

//// implemented all listeners methods
//var _ parser2.Scratch2Listener = &Scratch2Listener{}
//// implemented specific
//var _ parser2.Start1EntryListener = &Scratch2Listener{}
//var _ parser2.Start1ExitListener = &Scratch2Listener{}

//var _ parser2.A1EntryListener = &Scratch2Listener{}
//var _ parser2.A1ExitListener = &Scratch2Listener{}

//var _ parser2.B1EntryListener = &Scratch2Listener{}
//var _ parser2.B1ExitListener = &Scratch2Listener{}

//var _ parser2.B2EntryListener = &Scratch2Listener{}
//var _ parser2.B2ExitListener = &Scratch2Listener{}

//var _ parser2.C1EntryListener = &Scratch2Listener{}
//var _ parser2.C1ExitListener = &Scratch2Listener{}

//var _ parser2.DEntryListener = &Scratch2Listener{}
//var _ parser2.DExitListener = &Scratch2Listener{}
 

//type Scratch2Listener struct {
//}

//type Scratch2ErrorListener struct {
//    Warning string
//    Err     error
//    Debug   bool
//}

// func (cb *Scratch2ErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
//  if cb.Debug {
//      fmt.Printf("SyntaxError %d:%d <%s>\n", line, column, msg)
//  }
//  if strings.HasPrefix(msg, "report") { // TODO remove NewDiagnosticErrorListener and move warning to ReportAmbiguity etc. when getDecisionDescription is make public
//      cb.Warning = fmt.Sprintf("At %d:%d <%s>", line, column, msg)
//  } else {
//      cb.Err = fmt.Errorf("SyntaxError %d:%d <%s>", line, column, msg)
//  }
// }
// func (cb *Scratch2ErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int,
//  exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
//  }
// }
// func (cb *Scratch2ErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
//  }
// }
// func (cb *Scratch2ErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportContextSensitivity rec:%v dfs:%v start:%d stop:%d, prediction:%v configs:%v\n", recognizer, dfa, startIndex, stopIndex, prediction, configs)
//  }
// }


//// antlr.ParseTreeListener implementation.
//// All required.

//func (s *Scratch2Listener ) VisitTerminal(node  antlr.TerminalNode) {   }
//func (s *Scratch2Listener ) VisitErrorNode(node antlr.ErrorNode)    {   }
//func (s *Scratch2Listener ) EnterEveryRule(ctx antlr.ParserRuleContext) {  }
//func (s *Scratch2Listener ) ExitEveryRule(ctx antlr.ParserRuleContext) {  }

//// Only implemented as needed.

//func (s *Scratch2Listener) EnterStart1(ctx parser2.IStart1Context) {}
//func (s *Scratch2Listener) ExitStart1(ctx parser2.IStart1Context) {}

//func (s *Scratch2Listener) EnterA1(ctx parser2.IA1Context) {}
//func (s *Scratch2Listener) ExitA1(ctx parser2.IA1Context) {}

//func (s *Scratch2Listener) EnterB1(ctx parser2.IB1Context) {}
//func (s *Scratch2Listener) ExitB1(ctx parser2.IB1Context) {}

//func (s *Scratch2Listener) EnterB2(ctx parser2.IB2Context) {}
//func (s *Scratch2Listener) ExitB2(ctx parser2.IB2Context) {}

//func (s *Scratch2Listener) EnterC1(ctx parser2.IC1Context) {}
//func (s *Scratch2Listener) ExitC1(ctx parser2.IC1Context) {}

//func (s *Scratch2Listener) EnterD(ctx parser2.IDContext) {}
//func (s *Scratch2Listener) ExitD(ctx parser2.IDContext) {}


