package eval

import (
	"reflect"
	"testing"

	"github.com/wxio/antlr4-go-examples/eval/walker"
)

func TestEval(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected walker.Atom
	}{
		{"simple" /*      */, "1+2", &IntegerAtom{Val: 3}},
		{"simple" /*      */, "1+2+4", &IntegerAtom{Val: 7}},
		{"float float" /*      */, "1.0+2.0", &FloatAtom{Val: 3}},
		{"int float" /*      */, "1+2.0", &FloatAtom{Val: 3}},
		{"precedents" /*  */, "1+2^3", &IntegerAtom{Val: 9}},
		{"precedents" /*  */, "1^2+3", &IntegerAtom{Val: 4}},
		{"precedents" /*  */, "1^2^3", &IntegerAtom{Val: 1}},
		{"neg" /* 		  */, "--1", &IntegerAtom{Val: 1}},
		{"neg" /* 		  */, "+-1", &IntegerAtom{Val: -1}},
		{"neg" /* 		  */, "-1+2", &IntegerAtom{Val: 1}},
		{"percent" /* 	  */, "1%", &FloatAtom{Val: 0.01}},
		{"percent" /* 	  */, "1%+2", &FloatAtom{Val: 2.01}},
		// {"compare" /* 	  */, "1>=2", "(>= 1 2)"},
		// {"compare" /* 	  */, "1+2+4", "(+ (+ 1 2) 4)"},
		// {"compare" /* 	  */, "1=2<>3", "(<> (= 1 2) 3)"},
		{"atom" /*		  */, "1", &IntegerAtom{Val: 1}},
		{"atom" /*		  */, "1.00", &FloatAtom{Val: 1}},
		// {"range00" /*		  */, "a1:x10", "(RANGE a1 x10)"},
		// {"range00" /*		  */, "a1:x10 +  a1:x10", "(+ (RANGE a1 x10) RANGE a1 x10)"},
		// {"range01" /*		  */, "a:a", "(RANGE a a)"},
		// {"range02" /*		  */, "a:1", "ERROR:SyntaxError 1:3 <mismatched input '<EOF>' expecting ':'>"},
		// {"range" /*		  */, "1:1", "(HRANGE 1 1)"},
		// {"intersect" /*	  */, "a1 x10", "(INTERSECT a1 x10)"},
		// {"intersect" /*	  */, "a1:x10 x9", "(INTERSECT (RANGE a1 x10) x9)"},
		// // {"union" /*		  */, "a1,x10", "(ERROR Syntax)"},
		// // {"union" /*		  */, "a,b,c", "(ERROR Syntax)"}, // currently ReportContextSensitivity
		// {"union" /*		  */, "(a1,x10)", "(BLOCKORUNION a1 x10)"},
		// {"union" /*		  */, "(a,b,c)", "(BLOCKORUNION a b c)"}, // currently ReportContextSensitivity
		{"name" /*		  */, `"ab12cd"`, &StringAtom{Val: `"ab12cd"`}},
		// {"func" /*		  */, "SUM()", "(FUNC SUM)"},
		// {"func" /*		  */, "SUM(1)", "(FUNC SUM 1)"},
		// {"func" /*		  */, "SUM(1,2)", "(FUNC SUM 1 2)"},
		// {"func" /*		  */, "SUM(a1:x10)", "(FUNC SUM (RANGE a1 x10))"},
		// {"func" /*		  */, "SUM(a1:x10)+1", "(+ (FUNC SUM (RANGE a1 x10)) 1)"},
		// {"func" /*		  */, "Index():x10", "(RANGE (FUNC Index) x10)"},
		// {"block" /*		  */, "(A1)", "(BLOCKORUNION A1)"},
		// {"block" /*		  */, "((A1))", "(BLOCKORUNION (BLOCKORUNION A1))"},
		// {"block" /*		  */, "(A1,A2)", "(BLOCKORUNION A1 A2)"},       // ???
		// {"block" /*		  */, "(A1,A2,A3)", "(BLOCKORUNION A1 A2 A3)"}, // ???
		{"block" /*		  */, "(1+2)^3", &IntegerAtom{Val: 27}},
		{"error" /*		  */, "#ERROR!", &ErrorAtom{Text: "ERROR!"}},
		// {"error" /*		  */, "#REALLY?", "(ERROR REALLY?)"},
		// {"space" /*		  */, "-    1", "(- 1)"},
		// {"space" /*		  */, "5   %", "(% 5)"},
		// {"space" /*		  */, "1   ^    2", "(^ 1 2)"},
		// {"space" /*		  */, "1   =    2", "(= 1 2)"},
		// {"space" /*		  */, "(a   ,    b)", "(BLOCKORUNION a b)"},
		// {"space" /*		  */, "SUM(    a   ,    b)", "(FUNC SUM a b)"},
		// {"enhancedRefs" /*		  */, "[abc][123]", "(EH [abc] [123])"},
		// {"enhancedRefs" /*		  */, "[a[b]c][123]", "(EH [a[b]c] [123])"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			func() {
				got, err := BuildEvalAST(tt.input)
				if err != nil {
					t.Errorf(`Builder Error
received '%v', 
received '%v', 
expected '%v'
`, err.Error(), got, tt.expected)
					return
				}
				value, err := WalkTree(got)
				if err != nil {
					t.Errorf("Output err %v\n%s\n", err, got.SExpr(nil))
				}
				if reflect.TypeOf(value) != reflect.TypeOf(tt.expected) || value.Value() != tt.expected.Value() {
					t.Errorf(`Output mismatch
expected '%+#v'
received '%+#v'
received '%v'
received '%+#v'
				`, tt.expected, value, got.TreeString(), got)
				}
			}()
		})
	}
}
