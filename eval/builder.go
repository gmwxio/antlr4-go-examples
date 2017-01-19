package eval

import (
	"fmt"
	"strconv"

	"math"

	"reflect"

	"strings"

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

type ENode struct{ MyToken }
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

type RangeRef struct{ MyToken }
type IntersectRef struct{ MyToken }
type FuncNode struct {
	MyToken
	Name string
}

func (f *FuncNode) GetName() string {
	return f.Name
}

type ArgNode struct {
	MyToken
}

type BlockOrRef struct{ MyToken }
type A1Ref struct {
	MyToken
	Col    string
	Row    string
	ColAbs bool
	RowAbs bool
}
type NamedRangeRef struct {
	MyToken
	Name string
}
type HRangeRef struct{ MyToken }
type EHRef struct{ MyToken }

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

func (e *FuncNode) EvalLazy(strictArgs []Atom, lazyArgs []ctree.Tree) Atom {
	// IF
	if strings.ToLower(e.Name) == "if" {
		var choose bool
		test := strictArgs[0]
		switch test := test.(type) {
		case *BoolAtom:
			choose = test.Val
		case *IntegerAtom:
			choose = test.Val != 0
		case *FloatAtom:
			choose = test.Val != 0
		case *StringAtom:
			return &ErrorAtom{Text: "#Value!"}
		case *ErrorAtom:
			return test
		}
		var arg ctree.Tree
		if choose {
			arg = lazyArgs[0]
		} else {
			if len(lazyArgs) == 1 {
				return &BoolAtom{Val: false}
			}
			arg = lazyArgs[1]
		}
		got, err := WalkTreeFromStartExpr(arg)
		if err != nil {
			return &ErrorAtom{Text: err.Error()}
		}
		return got
	} else {
		panic("")
	}
}

func (e *FuncNode) EvalStrict(args []Atom) Atom {
	// TODO move to map
	if strings.ToLower(e.Name) == "sum" {
		var acc Atom = &IntegerAtom{}
		for _, a := range args {
			acc, a = convertTypes(acc, a, MathPrecedences)
			acc = acc.Add(a)
		}
		return acc
	}
	return nil
}

func (e *ENode) String() string      { return "=" }
func (e *BinaryNode) String() string { return fmt.Sprintf("%s", e.Operator) }
func (e *UnaryNode) String() string  { return fmt.Sprintf("%s", e.Operator) }

func (e *ErrorAtom) String() string   { return "ERROR " + e.Text }
func (e *IntegerAtom) String() string { return strconv.Itoa(e.Val) }
func (e *FloatAtom) String() string   { return fmt.Sprintf("%f", e.Val) }
func (e *StringAtom) String() string  { return e.Val }
func (e *BoolAtom) String() string    { return fmt.Sprintf("%v", e.Val) }

func (a *RangeRef) String() string      { return "RANGE" }
func (a *IntersectRef) String() string  { return "INTERSECT" }
func (a *FuncNode) String() string      { return "FUNC " + a.Name }
func (a *ArgNode) String() string       { return "ARG" }
func (a *BlockOrRef) String() string    { return "BLOCKORUNION" }
func (a *HRangeRef) String() string     { return "HRANGE" }
func (a *EHRef) String() string         { return "EH" }
func (a *NamedRangeRef) String() string { return a.Name }
func (a *A1Ref) String() string {
	if a.ColAbs && a.RowAbs {
		return "$" + a.Col + "$" + a.Row
	}
	if a.ColAbs && !a.RowAbs {
		return "$" + a.Col + a.Row
	}
	if !a.ColAbs && a.RowAbs {
		return a.Col + "$" + a.Row
	}
	return a.Col + a.Row
}

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
	warning string
	err     string
	debug   bool
}

// EnterEveryRule is called when any rule is entered.
func (s *ExprBuildListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	switch ctx := ctx.(type) {
	case *parser.StartContext:
	case *parser.CellContext:
		r := &ENode{MyToken: MyToken{Token: ctx.EQ().GetSymbol(), TType: parser.ExprParserROOT}}
		s.Builder = ctree.NewWalkableBuild("TREE", r)
	case *parser.StartexprContext:
		r := &ENode{MyToken: MyToken{Token: ctx.EOF().GetSymbol(), TType: parser.ExprParserROOT}}
		s.Builder = ctree.NewWalkableBuild("TREE", r)
	// Expressions
	case *parser.InfixContext:
		n := &BinaryNode{MyToken: MyToken{Token: ctx.GetT(), TType: parser.ExprParserBINARYEXPR}, Operator: ctx.GetT().GetText()}
		s.Builder.Add(n)
		s.Builder.Down()
	case *parser.ReferenceContext:
	case *parser.PrefixContext:
		n := &UnaryNode{MyToken: MyToken{Token: ctx.GetPre(), TType: parser.ExprParserUNARYEXPR}, Operator: ctx.GetPre().GetText()}
		s.Builder.Add(n)
		s.Builder.Down()
	case *parser.PostfixContext:
		n := &UnaryNode{MyToken: MyToken{Token: ctx.GetPost(), TType: parser.ExprParserUNARYEXPR}, Operator: ctx.GetPost().GetText()}
		s.Builder.Add(n)
		s.Builder.Down()
	// Atoms
	case *parser.ErrorContext:
	case *parser.IntegerContext:
		v, err := strconv.Atoi(ctx.GetA().GetText())
		if err != nil {
			panic("")
		}
		n := &IntegerAtom{MyToken: MyToken{Token: ctx.GetA(), TType: parser.ExprParserINTEGERATOM}, Val: v}
		s.Builder.Add(n)
	case *parser.FloatContext:
		var v float64
		_, err := fmt.Sscanf(ctx.GetA().GetText(), "%f", &v)
		if err != nil {
			panic("")
		}
		n := &FloatAtom{MyToken: MyToken{Token: ctx.GetA(), TType: parser.ExprParserFLOATATOM}, Val: v}
		s.Builder.Add(n)
	case *parser.StringContext:
		n := &StringAtom{MyToken: MyToken{Token: ctx.GetA(), TType: parser.ExprParserSTRINGATOM}, Val: ctx.GetA().GetText()}
		s.Builder.Add(n)
	case *parser.ErrorTypeContext:
		e := &ErrorAtom{MyToken: MyToken{Token: ctx.GetA(), TType: parser.ExprParserERRORATOM}}
		if ctx.GetB() != nil {
			e.Text = ctx.GetA().GetText() + ctx.GetB().GetText() + ctx.GetC().GetText()
		}
		e.Text = ctx.GetA().GetText() + ctx.GetC().GetText()
		s.Builder.Add(e)
	case *parser.TrueContext:
		n := &BoolAtom{MyToken: MyToken{Token: ctx.GetA(), TType: parser.ExprParserBOOLATOM}, Val: true}
		s.Builder.Add(n)
	case *parser.FalseContext:
		n := &BoolAtom{MyToken: MyToken{Token: ctx.GetA(), TType: parser.ExprParserBOOLATOM}, Val: false}
		s.Builder.Add(n)
	// REFS
	case *parser.IntersectContext:
		n := &IntersectRef{
			MyToken: MyToken{Token: ctx.WS(0).GetSymbol(), TType: parser.ExprParserINTERSECTREF}}
		s.Builder.Add(n)
		s.Builder.Down()
	case *parser.RangeContext:
		n := &RangeRef{
			MyToken: MyToken{Token: ctx.COLON().GetSymbol(), TType: parser.ExprParserRANGEREF}}
		s.Builder.Add(n)
		s.Builder.Down()
		//	case *parser.A1RefContext:
	case *parser.BlockOrRefContext:
		n := &BlockOrRef{MyToken: MyToken{Token: ctx.LP().GetSymbol(), TType: parser.ExprParserBLOCKORREFEXPR}}
		s.Builder.Add(n)
		s.Builder.Down()
	case *parser.FunctionCallContext:
		name := ctx.GetName().GetText()
		n := &FuncNode{MyToken: MyToken{
			Token: ctx.LP().GetSymbol(), TType: parser.ExprParserFUNCEXPR,
		},
			Name: name,
		}
		n.SetText("FUNCEXPR")
		s.Builder.Add(n)
		s.Builder.Down()
	case *parser.ArgContext:
		// n := &ArgNode{MyToken: MyToken{
		// 	Token: ctx.GetE().GetStart(), TType: parser.ExprParserARG,
		// }}
		// s.Builder.Add(n)
		// s.Builder.Down()
	case *parser.NamedRangeContext:
		name := ctx.GetFirst().GetText()
		for _, r := range ctx.GetRest() {
			name = name + r.GetText()
		}
		n := &NamedRangeRef{MyToken: MyToken{
			Token: ctx.GetFirst(), TType: parser.ExprParserNAMEDREF,
		},
			Name: name,
		}
		s.Builder.Add(n)
	case *parser.NameOrA1Context:
		n := &A1Ref{MyToken: MyToken{
			Token: ctx.CHAR().GetSymbol(), TType: parser.ExprParserA1REF,
		},
			Col: ctx.CHAR().GetSymbol().GetText(),
			Row: ctx.INT().GetSymbol().GetText(),
		}
		s.Builder.Add(n)
	case *parser.A1ColumnAbsContext:
		n := &A1Ref{MyToken: MyToken{
			Token: ctx.CHAR().GetSymbol(), TType: parser.ExprParserA1REF,
		},
			Col:    ctx.CHAR().GetSymbol().GetText(),
			Row:    ctx.INT().GetSymbol().GetText(),
			ColAbs: true,
		}
		s.Builder.Add(n)
	case *parser.A1RowAbsContext:
		n := &A1Ref{MyToken: MyToken{
			Token: ctx.CHAR().GetSymbol(), TType: parser.ExprParserA1REF,
		},
			Col:    ctx.CHAR().GetSymbol().GetText(),
			Row:    ctx.INT().GetSymbol().GetText(),
			RowAbs: true,
		}
		s.Builder.Add(n)
	case *parser.A1AbsContext:
		n := &A1Ref{MyToken: MyToken{
			Token: ctx.CHAR().GetSymbol(), TType: parser.ExprParserA1REF,
		},
			Col:    ctx.CHAR().GetSymbol().GetText(),
			Row:    ctx.INT().GetSymbol().GetText(),
			ColAbs: true,
			RowAbs: true,
		}
		s.Builder.Add(n)

	case *parser.EHRefContext:
	case *parser.HRangeContext:
		n := &HRangeRef{MyToken: MyToken{
			Token: ctx.GetRow1(), TType: parser.ExprParserHRANGEREF,
		},
		}
		s.Builder.Add(n)
		s.Builder.Down()
		if ctx.GetRow1() == nil || ctx.GetRow2() == nil { // this only happen under antlr error recovery
			return
		}
		{
			v, err := strconv.Atoi(ctx.GetRow1().GetText())
			if err != nil {
				panic("")
			}
			n := &IntegerAtom{MyToken: MyToken{Token: ctx.GetRow1(), TType: parser.ExprParserINTEGERATOM}, Val: v}
			s.Builder.Add(n)
		}
		{
			v, err := strconv.Atoi(ctx.GetRow2().GetText())
			if err != nil {
				panic("")
			}
			n := &IntegerAtom{MyToken: MyToken{Token: ctx.GetRow2(), TType: parser.ExprParserINTEGERATOM}, Val: v}
			s.Builder.Add(n)
		}
	case *parser.EnhancedRefContext:
	case *parser.EnhancedRefPartsContext:
	case *parser.EnhancedRefContentContext:
	case *parser.ArgsContext:
	case *parser.FnameContext:
	case *parser.ElemsContext:
	}
}

func (s *ExprBuildListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
	switch ctx.(type) {
	case *parser.CellContext:
	case *parser.InfixContext:
		s.Builder.Up()
	case *parser.ReferenceContext:
	case *parser.PrefixContext:
		s.Builder.Up()
	case *parser.PostfixContext:
		s.Builder.Up()
	case *parser.ErrorContext:
	case *parser.IntersectContext:
		s.Builder.Up()
	case *parser.RangeContext:
		s.Builder.Up()
	case *parser.EHRefContext:
	case *parser.HRangeContext:
		s.Builder.Up()
	case *parser.A1RefContext:
	case *parser.BlockOrRefContext:
		s.Builder.Up()
	case *parser.FunctionCallContext:
		s.Builder.Up()
	case *parser.ArgContext:
		// s.Builder.Up()
	case *parser.EnhancedRefContext:
	case *parser.EnhancedRefPartsContext:
	case *parser.EnhancedRefContentContext:
	case *parser.ArgsContext:
	case *parser.FnameContext:
	case *parser.ElemsContext:
	}
}

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
	if strings.HasPrefix(msg, "report") { // TODO remove NewDiagnosticErrorListener and move warning to ReportAmbiguity etc. when getDecisionDescription is make public
		tbl.warning = fmt.Sprintf("At %d:%d <%s>", line, column, msg)
	} else {
		tbl.err = fmt.Sprintf("SyntaxError %d:%d <%s>", line, column, msg)
	}
	// panic("line " + strconv.Itoa(line) + ":" + strconv.Itoa(column) + " " + msg)
}

func (tbl *ExprBuildListener) ReportAmbiguity(
	recognizer antlr.Parser,
	dfa *antlr.DFA,
	startIndex, stopIndex int,
	exact bool,
	ambigAlts *antlr.BitSet,
	configs antlr.ATNConfigSet) {
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
