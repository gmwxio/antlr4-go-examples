// Generated from Scratch.g4 by ANTLR 4.7.

package parser // Scratch
//import "github.com/wxio/antlr4-go"
//import "generated code if in another package"

//// Commented out basic implementation for your convenience.

//func Example(s string) {
//  // Setup
//  input := antlr.NewInputStream(s)
//  lexer := parser.NewScratchLexer(input)
//  stream := antlr.NewCommonTokenStream(lexer, 0)
//  p := parser.NewScratchParser(stream)

//  // Antlr error listener - turns reports (ambiguity etc) into syntax errors
//  p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

//  // Custom error listener, register before the parse
//  el := &ScratchErrorListener{}
//  p.AddErrorListener(el)

//  // Parse - start rule
//  tree := p.Start()

//  // Antlr provided parse tree representation
//  sexpr := antlr.TreesStringTree(tree, nil, p)
//  fmt.Printf("%s\n", sexpr)

//  // Custom listener
//  l := &ScratchListener{}
//  antlr.ParseTreeWalkerDefault.Walk(l, tree)

//  // Custom visitor
//  v := &ScratchVisitor{}
//  tree.Accept(v)
// }

//// implemented all listeners methods
//var _ parser.ScratchListener = &ScratchListener{}
//// implemented specific
//
//    

//    

//    
//
//var _ parser.Start1EntryListener = &ScratchListener{}
//var _ parser.Start1ExitListener = &ScratchListener{}

//var _ parser.A1EntryListener = &ScratchListener{}
//var _ parser.A1ExitListener = &ScratchListener{}

//var _ parser.B1EntryListener = &ScratchListener{}
//var _ parser.B1ExitListener = &ScratchListener{}
 

//type ScratchListener struct {
//}

//type ScratchErrorListener struct {
//    Warning string
//    Err     error
//    Debug   bool
//}

// func (cb *ScratchErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
//  if cb.Debug {
//      fmt.Printf("SyntaxError %d:%d <%s>\n", line, column, msg)
//  }
//  if strings.HasPrefix(msg, "report") { // TODO remove NewDiagnosticErrorListener and move warning to ReportAmbiguity etc. when getDecisionDescription is make public
//      cb.Warning = fmt.Sprintf("At %d:%d <%s>", line, column, msg)
//  } else {
//      cb.Err = fmt.Errorf("SyntaxError %d:%d <%s>", line, column, msg)
//  }
// }
// func (cb *ScratchErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int,
//  exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
//  }
// }
// func (cb *ScratchErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
//  }
// }
// func (cb *ScratchErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
//  if cb.Debug {
//      fmt.Printf("ReportContextSensitivity rec:%v dfs:%v start:%d stop:%d, prediction:%v configs:%v\n", recognizer, dfa, startIndex, stopIndex, prediction, configs)
//  }
// }


//// antlr.ParseTreeListener implementation.
//// All required.

//func (s *ScratchListener ) VisitTerminal(node  antlr.TerminalNode) {   
//}
//func (s *ScratchListener ) VisitErrorNode(node antlr.ErrorNode)    {   
//}
//func (s *ScratchListener ) EnterEveryRule(ctx antlr.ParserRuleContext) {  
//}
//func (s *ScratchListener ) ExitEveryRule(ctx antlr.ParserRuleContext) {  
//}

//// Only implemented as needed.

//func (s *ScratchListener) EnterStart1(ctx parser.IStart1Context) {
//}
//func (s *ScratchListener) ExitStart1(ctx parser.IStart1Context) {
//}

//func (s *ScratchListener) EnterA1(ctx parser.IA1Context) {
//}
//func (s *ScratchListener) ExitA1(ctx parser.IA1Context) {
//}

//func (s *ScratchListener) EnterB1(ctx parser.IB1Context) {
//}
//func (s *ScratchListener) ExitB1(ctx parser.IB1Context) {
//}


