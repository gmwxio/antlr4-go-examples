package eval

import (
	"fmt"
	"strconv"

	"math"

	"reflect"

	antlr "github.com/wxio/antlr4-go"
	"github.com/wxio/antlr4-go-examples/ctree"
	"github.com/wxio/antlr4-go-examples/eval/parser"
	. "github.com/wxio/antlr4-go-examples/eval/walker"
)

type MyToken struct {
	antlr.Token
	TType int
}

func (t *MyToken) GetTokenType() int { return t.TType }

type ENode struct {
	MyToken
}
type BinaryNode struct {
	MyToken
	Operator string
}
type UnaryNode struct {
	MyToken
	Operator string
}
type ErrorAtom struct {
	MyToken
	Text string
}
type FloatAtom struct {
	MyToken
	Val float64
}
type IntegerAtom struct {
	MyToken
	Val int
}
type BoolAtom struct {
	MyToken
	Val bool
}
type StringAtom struct {
	MyToken
	Val string
}

func (e *BinaryNode) Eval(left, right Atom) Atom {

	switch e.Operator {
	case "+":
		left, right = convertTypes(left, right, MathPrecedences)
		return left.Add(right)
	case "-":
	case "*":
	case "/":
	case "^":
		left, right = convertTypes(left, right, MathPrecedences)
		v := left.Exp(right)
		return v
	case "&":
	case ">":
	case "<":
	case ">=":
	case "<=":
	case "=":
	case "<>":
	}
	panic("")
}

func convertTypes(left, right Atom, p Precendence) (Atom, Atom) {
	ltype := reflect.TypeOf(left)
	rtype := reflect.TypeOf(right)
	if ltype == rtype {
		return left, right
	}
	ctype := p.Tree().CommonAncestor(ltype, rtype)
	for ltype != ctype {
		left = p.ConvertUp(ltype)(left)
		ltype = reflect.TypeOf(left)
	}
	for rtype != ctype {
		right = p.ConvertUp(rtype)(right)
		rtype = reflect.TypeOf(right)
	}
	return left, right
}

func (e *UnaryNode) Eval(arg Atom) Atom {
	switch e.Operator {
	case "+":
		return arg
	case "-":
		return arg.Neg()
	case "%":
		return arg.Percent()
	}
	panic("")
}

// func (e *UnaryExpr) String() string        { return fmt.Sprintf("%s", e.Operator) }
// func (e *FunctionCallExpr) String() string { return "FUNC " + e.Name }
// func (e *BlockOrRefExpr) String() string   { return "BLOCKORUNION" }
func (e *ENode) String() string      { return "=" }
func (e *BinaryNode) String() string { return fmt.Sprintf("%s", e.Operator) }
func (e *UnaryNode) String() string  { return fmt.Sprintf("%s", e.Operator) }

func (e *ErrorAtom) String() string   { return "ERROR " + e.Text }
func (e *IntegerAtom) String() string { return strconv.Itoa(e.Val) }
func (e *FloatAtom) String() string   { return fmt.Sprintf("%f", e.Val) }
func (e *StringAtom) String() string  { return e.Val }
func (e *BoolAtom) String() string    { return fmt.Sprintf("%v", e.Val) }

// func (e *RangeRef) String() string    { return "RANGE" }
// func (e *IntersectRef) String() string { return "INTERSECT" }
// func (e *EnhancedRef) String() string  { return "EH" }
// func (e *HRangeRef) String() string    { return "HRANGE" }

func (a *ErrorAtom) Value() interface{}   { return a.Text }
func (a *FloatAtom) Value() interface{}   { return a.Val }
func (a *IntegerAtom) Value() interface{} { return a.Val }
func (a *StringAtom) Value() interface{}  { return a.Val }
func (a *BoolAtom) Value() interface{}    { return a.Val }

func (a *ErrorAtom) Add(b Atom) Atom   { return a }
func (a *FloatAtom) Add(b Atom) Atom   { return &FloatAtom{Val: a.Val + b.(*FloatAtom).Val} }
func (a *IntegerAtom) Add(b Atom) Atom { return &IntegerAtom{Val: a.Val + b.(*IntegerAtom).Val} }
func (a *StringAtom) Add(b Atom) Atom  { return &ErrorAtom{Text: "#Value!"} }
func (a *BoolAtom) Add(b2 Atom) Atom {
	var b *BoolAtom = b2.(*BoolAtom)
	if a.Val && b.Val {
		return &IntegerAtom{Val: 2}
	}
	if a.Val || b.Val {
		return &IntegerAtom{Val: 1}
	}
	return &IntegerAtom{Val: 0}
}

func (a *ErrorAtom) Exp(b Atom) Atom   { return a }
func (a *FloatAtom) Exp(b Atom) Atom   { return &FloatAtom{Val: math.Pow(a.Val, b.(*FloatAtom).Val)} }
func (a *IntegerAtom) Exp(b Atom) Atom { return &IntegerAtom{Val: power(a.Val, b.(*IntegerAtom).Val)} }
func power(x, n int) int {             //Exponentiation by squaring
	if n == 1 {
		return x
	}
	if n%2 == 0 {
		return power(x*x, n/2)
	}
	return x * power(x*x, (n-1)/2)
}
func (a *StringAtom) Exp(b Atom) Atom { return &ErrorAtom{Text: "#Value!"} }
func (a *BoolAtom) Exp(b2 Atom) Atom {
	var b *BoolAtom = b2.(*BoolAtom)
	if a.Val {
		return &IntegerAtom{Val: 1}
	}
	if !a.Val && b.Val {
		return &IntegerAtom{Val: 0}
	}
	return &ErrorAtom{Text: "#Num!"}
}

func (a *ErrorAtom) Neg() Atom   { return a }
func (a *FloatAtom) Neg() Atom   { return &FloatAtom{Val: -a.Val} }
func (a *IntegerAtom) Neg() Atom { return &IntegerAtom{Val: -a.Val} }
func (a *StringAtom) Neg() Atom  { return &ErrorAtom{Text: "#Value!"} }
func (a *BoolAtom) Neg() Atom {
	if a.Val {
		return &IntegerAtom{Val: -1}
	}
	return &IntegerAtom{Val: 0}
}

func (a *ErrorAtom) Percent() Atom   { return a }
func (a *FloatAtom) Percent() Atom   { return &FloatAtom{Val: a.Val / 100} }
func (a *IntegerAtom) Percent() Atom { return &FloatAtom{Val: float64(a.Val) / 100} }
func (a *StringAtom) Percent() Atom  { return &ErrorAtom{Text: "#Value!"} }
func (a *BoolAtom) Percent() Atom {
	if a.Val {
		return &FloatAtom{Val: 0.01}
	}
	return &FloatAtom{Val: 0}
}

// type Atom interface {
// 	Value() interface{}
// 	Add(b Atom) Atom
// 	Exp(b Atom) Atom
// 	// Mult(b Atom) Atom
// 	// Div(b Atom) Atom

// 	Neg() Atom
// 	Percent() Atom
// }

var (
	_ Atom = &ErrorAtom{}
	_ Atom = &FloatAtom{}
	_ Atom = &IntegerAtom{}
	_ Atom = &BoolAtom{}
	_ Atom = &StringAtom{}
)

type ExprBuildListener struct {
	*parser.BaseExprParserListener
	Builder ctree.WalkableBuilder
	err     string
	debug   bool
}

// EnterStart is called when production start is entered.
func (s *ExprBuildListener) EnterStart(ctx *parser.StartContext) {
}

// EnterCell is called when production cell is entered.
func (s *ExprBuildListener) EnterCell(ctx *parser.CellContext) {
	r := &ENode{MyToken: MyToken{Token: ctx.EQ().GetSymbol(), TType: parser.ExprParserROOT}}
	s.Builder = ctree.NewWalkableBuild("TREE", r)
}

// ExitCell is called when production cell is exited.
func (s *ExprBuildListener) ExitCell(ctx *parser.CellContext) {}

// EnterStartexpr is called when production startexpr is entered.
func (s *ExprBuildListener) EnterStartexpr(ctx *parser.StartexprContext) {
	r := &ENode{MyToken: MyToken{Token: ctx.EOF().GetSymbol(), TType: parser.ExprParserROOT}}
	s.Builder = ctree.NewWalkableBuild("TREE", r)
}

// EnterInfix is called when production Infix is entered.
func (s *ExprBuildListener) EnterInfix(ctx *parser.InfixContext) {
	n := &BinaryNode{MyToken: MyToken{Token: ctx.GetT(), TType: parser.ExprParserBINARYEXPR}, Operator: ctx.GetT().GetText()}
	s.Builder.Add(n)
	s.Builder.Down()
}

// ExitInfix is called when production Infix is exited.
func (s *ExprBuildListener) ExitInfix(ctx *parser.InfixContext) {
	s.Builder.Up()
}

// EnterReference is called when production Reference is entered.
func (s *ExprBuildListener) EnterReference(ctx *parser.ReferenceContext) {}

// ExitReference is called when production Reference is exited.
func (s *ExprBuildListener) ExitReference(ctx *parser.ReferenceContext) {}

// EnterPrefix is called when production Prefix is entered.
func (s *ExprBuildListener) EnterPrefix(ctx *parser.PrefixContext) {
	n := &UnaryNode{MyToken: MyToken{Token: ctx.GetPre(), TType: parser.ExprParserUNARYEXPR}, Operator: ctx.GetPre().GetText()}
	s.Builder.Add(n)
	s.Builder.Down()
}

// ExitPrefix is called when production Prefix is exited.
func (s *ExprBuildListener) ExitPrefix(ctx *parser.PrefixContext) {
	s.Builder.Up()
}

// EnterPostfix is called when production Postfix is entered.
func (s *ExprBuildListener) EnterPostfix(ctx *parser.PostfixContext) {
	n := &UnaryNode{MyToken: MyToken{Token: ctx.GetPost(), TType: parser.ExprParserUNARYEXPR}, Operator: ctx.GetPost().GetText()}
	s.Builder.Add(n)
	s.Builder.Down()
}

// ExitPostfix is called when production Postfix is exited.
func (s *ExprBuildListener) ExitPostfix(ctx *parser.PostfixContext) {
	s.Builder.Up()
}

// EnterError is called when production Error is entered.
func (s *ExprBuildListener) EnterError(ctx *parser.ErrorContext) {}

// ExitError is called when production Error is exited.
func (s *ExprBuildListener) ExitError(ctx *parser.ErrorContext) {}

// EnterInteger is called when production Integer is entered.
func (s *ExprBuildListener) EnterInteger(ctx *parser.IntegerContext) {
	v, err := strconv.Atoi(ctx.GetA().GetText())
	if err != nil {
		panic("")
	}
	n := &IntegerAtom{MyToken: MyToken{Token: ctx.GetA(), TType: parser.ExprParserINTEGERATOM}, Val: v}
	s.Builder.Add(n)
}

// EnterFloat is called when production Float is entered.
func (s *ExprBuildListener) EnterFloat(ctx *parser.FloatContext) {
	var v float64
	_, err := fmt.Sscanf(ctx.GetA().GetText(), "%f", &v)
	if err != nil {
		panic("")
	}
	n := &FloatAtom{MyToken: MyToken{Token: ctx.GetA(), TType: parser.ExprParserFLOATATOM}, Val: v}
	s.Builder.Add(n)
}

// EnterString is called when production String is entered.
func (s *ExprBuildListener) EnterString(ctx *parser.StringContext) {
	n := &StringAtom{MyToken: MyToken{Token: ctx.GetA(), TType: parser.ExprParserSTRINGATOM}, Val: ctx.GetA().GetText()}
	s.Builder.Add(n)
}

// EnterErrorType is called when production errorType is entered.
func (s *ExprBuildListener) EnterErrorType(ctx *parser.ErrorTypeContext) {
	e := &ErrorAtom{MyToken: MyToken{Token: ctx.GetA(), TType: parser.ExprParserERRORATOM}}
	if ctx.GetB() != nil {
		e.Text = ctx.GetA().GetText() + ctx.GetB().GetText() + ctx.GetC().GetText()
	}
	e.Text = ctx.GetA().GetText() + ctx.GetC().GetText()
	s.Builder.Add(e)
}

// EnterTrue is called when production True is entered.
func (s *ExprBuildListener) EnterTrue(ctx *parser.TrueContext) {
	n := &BoolAtom{MyToken: MyToken{Token: ctx.GetA(), TType: parser.ExprParserBOOLATOM}, Val: true}
	s.Builder.Add(n)
}

// EnterFalse is called when production False is entered.
func (s *ExprBuildListener) EnterFalse(ctx *parser.FalseContext) {
	n := &BoolAtom{MyToken: MyToken{Token: ctx.GetA(), TType: parser.ExprParserBOOLATOM}, Val: false}
	s.Builder.Add(n)
}

// EnterIntersect is called when production Intersect is entered.
func (s *ExprBuildListener) EnterIntersect(ctx *parser.IntersectContext) {}

// ExitIntersect is called when production Intersect is exited.
func (s *ExprBuildListener) ExitIntersect(ctx *parser.IntersectContext) {}

// EnterRange is called when production Range is entered.
func (s *ExprBuildListener) EnterRange(ctx *parser.RangeContext) {}

// ExitRange is called when production Range is exited.
func (s *ExprBuildListener) ExitRange(ctx *parser.RangeContext) {}

// EnterEHRef is called when production EHRef is entered.
func (s *ExprBuildListener) EnterEHRef(ctx *parser.EHRefContext) {}

// ExitEHRef is called when production EHRef is exited.
func (s *ExprBuildListener) ExitEHRef(ctx *parser.EHRefContext) {}

// EnterHRange is called when production HRange is entered.
func (s *ExprBuildListener) EnterHRange(ctx *parser.HRangeContext) {}

// ExitHRange is called when production HRange is exited.
func (s *ExprBuildListener) ExitHRange(ctx *parser.HRangeContext) {}

// EnterA1Ref is called when production A1Ref is entered.
func (s *ExprBuildListener) EnterA1Ref(ctx *parser.A1RefContext) {}

// ExitA1Ref is called when production A1Ref is exited.
func (s *ExprBuildListener) ExitA1Ref(ctx *parser.A1RefContext) {}

// EnterBlockOrRef is called when production BlockOrRef is entered.
func (s *ExprBuildListener) EnterBlockOrRef(ctx *parser.BlockOrRefContext) {}

// ExitBlockOrRef is called when production BlockOrRef is exited.
func (s *ExprBuildListener) ExitBlockOrRef(ctx *parser.BlockOrRefContext) {}

// EnterFunctionCall is called when production FunctionCall is entered.
func (s *ExprBuildListener) EnterFunctionCall(ctx *parser.FunctionCallContext) {}

// ExitFunctionCall is called when production FunctionCall is exited.
func (s *ExprBuildListener) ExitFunctionCall(ctx *parser.FunctionCallContext) {}

// EnterId_a1 is called when production id_a1 is entered.
func (s *ExprBuildListener) EnterId_a1(ctx *parser.Id_a1Context) {}

// ExitId_a1 is called when production id_a1 is exited.
func (s *ExprBuildListener) ExitId_a1(ctx *parser.Id_a1Context) {}

// EnterEnhancedRef is called when production enhancedRef is entered.
func (s *ExprBuildListener) EnterEnhancedRef(ctx *parser.EnhancedRefContext) {}

// ExitEnhancedRef is called when production enhancedRef is exited.
func (s *ExprBuildListener) ExitEnhancedRef(ctx *parser.EnhancedRefContext) {}

// EnterEnhancedRefParts is called when production enhancedRefParts is entered.
func (s *ExprBuildListener) EnterEnhancedRefParts(ctx *parser.EnhancedRefPartsContext) {}

// ExitEnhancedRefParts is called when production enhancedRefParts is exited.
func (s *ExprBuildListener) ExitEnhancedRefParts(ctx *parser.EnhancedRefPartsContext) {}

// EnterEnhancedRefContent is called when production enhancedRefContent is entered.
func (s *ExprBuildListener) EnterEnhancedRefContent(ctx *parser.EnhancedRefContentContext) {}

// ExitEnhancedRefContent is called when production enhancedRefContent is exited.
func (s *ExprBuildListener) ExitEnhancedRefContent(ctx *parser.EnhancedRefContentContext) {}

// EnterArgs is called when production args is entered.
func (s *ExprBuildListener) EnterArgs(ctx *parser.ArgsContext) {}

// ExitArgs is called when production args is exited.
func (s *ExprBuildListener) ExitArgs(ctx *parser.ArgsContext) {}

// EnterFname is called when production fname is entered.
func (s *ExprBuildListener) EnterFname(ctx *parser.FnameContext) {}

// ExitFname is called when production fname is exited.
func (s *ExprBuildListener) ExitFname(ctx *parser.FnameContext) {}

// EnterElems is called when production elems is entered.
func (s *ExprBuildListener) EnterElems(ctx *parser.ElemsContext) {}

// ExitElems is called when production elems is exited.
func (s *ExprBuildListener) ExitElems(ctx *parser.ElemsContext) {}

type errorListener struct {
}

func (d *errorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	panic("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg)
}

func (d *errorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// panic("ReportAmbiguity")
	fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

func (d *errorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	//	panic("ReportAttemptingFullContext")
	fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

func (d *errorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	// panic("ReportContextSensitivity")
}

func (tbl *ExprBuildListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	if tbl.debug {
		fmt.Printf("SyntaxError %d:%d <%s>\n", line, column, msg)
	}
	tbl.err = fmt.Sprintf("SyntaxError %d:%d <%s>", line, column, msg)
	// panic("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg)
}

func (tbl *ExprBuildListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// panic("ReportAmbiguity")
	if tbl.debug {
		fmt.Printf("ReportAmbiguity rec:%v dfs:%v start:%d stop:%d, exact:%v, ambigAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
	}
}

func (tbl *ExprBuildListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs antlr.ATNConfigSet) {
	// panic("ReportAttemptingFullContext")
	if tbl.debug {
		fmt.Printf("ReportAttemptingFullContext rec:%v dfs:%v start:%d stop:%d, conflictingAlts:%v config:%v\n", recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
	}
}

func (tbl *ExprBuildListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs antlr.ATNConfigSet) {
	//	panic("ReportContextSensitivity")
}
