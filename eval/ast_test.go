package eval

import (
	"fmt"
	"strings"
	"testing"
)

const debugtest = true

func TestBuildExprASTExpr(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"simple" /*      */, "1+2", "(+ (1 2))"},
		{"precedents" /*  */, "1+2^3", "(+ (1 ^ (2 3)))"},
		{"precedents" /*  */, "1^2+3", "(+ (^ (1 2) 3))"},
		{"precedents" /*  */, "1^2^3", "(^ (^ (1 2) 3))"},
		{"neg" /* 		  */, "--1", "(- (- (1)))"},
		{"neg" /* 		  */, "+-1", "(+ (- (1)))"},
		{"neg" /* 		  */, "-1+2", "(+ (- (1) 2))"},
		{"percent" /* 	  */, "1%+2", "(+ (% (1) 2))"},
		{"compare" /* 	  */, "1>=2", "(>= (1 2))"},
		{"compare" /* 	  */, "1<>2=3", "(= (<> (1 2) 3))"},
		{"compare" /* 	  */, "1=2<>3", "(<> (= (1 2) 3))"},
		{"atom" /*		  */, "1", "(1)"},
		// {"atom" /*		  */, "1.00", "1.000000"},
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
		{"name" /*		  */, "\"ab12cd\"", `("ab12cd")`},
		// {"a1ref" /*		  */, "ab12cd", `("ab12cd")`},
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
		// {"block" /*		  */, "(1+2)^3", "(^ (BLOCKORUNION (+ 1 2)) 3)"},
		{"error" /*		  */, "#ERROR!", "(ERROR ERROR!)"},
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

	// first
	// context required
	// filter expects a single

	const debugName = ""
	if debugName != "" {
		for _, tt := range tests {
			if debugName != strings.Replace(tt.name, " ", "_", -1) {
				continue
			}
			fmt.Println("**************")
			fmt.Println(tt.name)
			fmt.Printf("Input: \n\"\"\"\n%s' \n\"\"\"\n", tt.input)
			fmt.Println("Expecting: " + tt.expected)
			fmt.Println("-------------")
			fmt.Println(LexDebug(tt.input))
			fmt.Println("-------------")
			got, err := BuildEvalAST(tt.input)
			fmt.Printf("Input: '%+#v' \nReturned: '%s' error %v\n", tt.input, got, err)
			if err == nil {
				fmt.Printf("Returned: '%v'\n", got)
			}
			fmt.Println("**************")
		}
		return
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			func() {
				defer func() {
					if e := recover(); e != nil {
						debugLex(tt.input, e, t)
					}
				}()
				got, err := BuildEvalAST(tt.input)
				if err != nil {
					if err.Error() != tt.expected {
						t.Errorf(`Builder Error
received '%v', 
received '%v', 
expected '%v'
`, err.Error(), got, tt.expected)
					}
					return
				}
				var sexpr string
				sexpr = fmt.Sprintf("%v", got.SExpr(got.Root()))
				sexpr = sexpr[4:]
				sexpr = sexpr[:len(sexpr)-1]
				// sexpr = got.DebugTreeString()
				if sexpr != tt.expected {
					tokens := LexDebug(tt.input)
					t.Errorf(`Output mismatch   
expected '%v'
received '%v'
received '%v'
received '%+#v'
tokens   
%s`, tt.expected, sexpr, got.SExpr(got.Root()), got, tokens)
				}
			}()
		})
	}
}

func debugLex(input string, e interface{}, t *testing.T) {
	defer func() {
		if e2 := recover(); e2 != nil {
			t.Errorf("LEXER PANIC: %v\n", e2)
			if debugtest {
				fmt.Printf("LEXER PANIC: %v\n", e)
				panic(e2)
			}
		}
		if debugtest {
			panic(e)
		}
	}()
	tokens := LexDebug(input)
	t.Errorf("PARSER PANIC: %v\nLex Tokens:\n%s\n", e, tokens)
	if debugtest {
		fmt.Printf("PARSER PANIC: %v\nLex Tokens:\n%s\n", e, tokens)
	}
}
