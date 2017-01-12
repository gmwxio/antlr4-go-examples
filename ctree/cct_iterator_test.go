package ctree

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func TestSize(t *testing.T) {
	r := &mys{"root"}
	e := &mys{"="}
	p := &mys{"+"}
	o := &mys{"1"}
	two := &mys{"2"}
	tests := []struct {
		name     string
		input    Tree
		expected int
	}{
		{"simple", BuildTree("", r).Add(e).Down().Add(p).Add(o).Add(two).Up().Build(), 5},
		{"simple", BuildTree("", r).Add(e).Down().Add(p).Down().Add(o).Down().Add(two).Up().Build(), 5},
		{"simple", BuildTree("", r).Add(e).Add(&mys{"a"}).Build(), 3},
		{"simple", BuildTree("", r).Build(), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.input.Size() != tt.expected {
				t.Errorf(`
Expected %d
Received %d
`, tt.expected, tt.input.Size())
			}
		})
	}
}

func TestPathAsPos(t *testing.T) {
	r := &mys{"root"}
	e := &mys{"="}
	p := &mys{"+"}
	o := &mys{"1"}
	two := &mys{"2"}
	tests := []struct {
		name     string
		input    Tree
		node     INode
		expected []int
	}{
		{"0-0", BuildTree("", r).Add(e).Down().Add(p).Add(o).Add(two).Up().Build(), e, []int{0}},
		{"1000", BuildTree("", r).Add(e).Down().Add(p).Down().Add(o).Add(two).Up().Build(), two, []int{0, 0, 1}},
		{"1000", BuildTree("", r).Add(e).Down().Add(p).Add(o).Add(two).Up().Build(), two, []int{0, 2}},
		{"0#00", BuildTree("", r).Add(e).Add(&mys{"a"}).Build(), r, []int{}},
		{"0#01", BuildTree("", r).Build(), r, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.input.PathAsPosition(tt.node), tt.expected) {
				t.Errorf(`
Expected %v
Received %v
`, tt.expected, tt.input.PathAsPosition(tt.node))
			}
		})
	}
}

func TestCommonAncestor(t *testing.T) {
	r := &mys{"root"}
	a := &mys{"a"}
	a2 := &mys{"a2"}
	b := &mys{"b"}
	c := &mys{"c"}
	d := &mys{"d"}
	e := &mys{"e"}
	f := &mys{"f"}
	tree := BuildTree("", r).
		Add(a).Down().
		/**/ Add(b).Down().
		/* */ Add(c).Add(d).
		/**/ Up().
		/**/ Add(e).Add(f).
		Up().
		Add(a2).
		Build()
	tests := []struct {
		name     string
		n1, n2   INode
		expected INode
	}{
		{"0#00", r, r, r},
		{"0#01", c, d, b},
		{"0#02", f, d, a},
		{"0#03", a2, e, r},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tree.CommonAncestor(tt.n1, tt.n2) != tt.expected {
				t.Errorf(`
Expected %v
Received %v
`, tt.expected, tree.CommonAncestor(tt.n1, tt.n2))
			}
		})
	}
}

func TestIterator(t *testing.T) {
	r := &mys{"root"}
	e := &mys{"="}
	p := &mys{"+"}
	o := &mys{"1"}
	two := &mys{"2"}
	tests := []struct {
		name     string
		input    *ctree
		expected string
	}{
		{"simple", BuildTree("", r).Add(e).Down().Add(p).Add(o).Add(two).Up().Build().(*ctree), " (= (+ 1 2)) <EOF>"},
		{"simple", BuildTree("", r).Add(e).Add(&mys{"a"}).Build().(*ctree), " (= a) <EOF>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := bytes.Buffer{}
			iter := NewPreOrderTreeIterator(tt.input, e)
			var start bool
			for iter.HasNext() {
				// fmt.Printf("1.%v\n", iter.stack)
				e := iter.Next()
				if _, ok := e.(*Down); ok {
					start = true
					buf.WriteString(fmt.Sprintf(" %v", e))
				} else if _, ok := e.(*Up); ok {
					buf.WriteString(fmt.Sprintf("%v", e))
				} else if start {
					buf.WriteString(fmt.Sprintf("%v", e))
					start = false
				} else {
					buf.WriteString(fmt.Sprintf(" %v", e))
				}
			}
			rec := buf.String()
			if tt.expected != rec {
				t.Errorf(`
Expected %s
Received '%s'				
				`, tt.expected, rec)
			}
		})
	}
}

func TestSExpr02(t *testing.T) {
	r := &mys{"root"}
	e := &mys{"="}
	p := &mys{"+"}
	o := &mys{"1"}
	two := &mys{"2"}
	tests := []struct {
		name      string
		input     *ctree
		expected  string
		expected2 string
	}{
		{"simple",
			BuildTree("", r).Add(e).Down().Add(&mys{"a"}).
				Add(&mys{"b"}).
				Add(&mys{"c"}).
				Add(&mys{"d"}).
				Add(p).Add(o).Add(two).Up().Build().(*ctree),
			" (root (= (a b c d + 1 2)))",
			" (+ 1 2)",
		},
		{"simple",
			BuildTree("", r).Add(e).
				Down().Add(&mys{"a"}).
				Down().Add(&mys{"b"}).
				Down().Add(&mys{"c"}).
				Down().Add(&mys{"d"}).
				Down().Add(p).
				Down().Add(o).
				Down().Add(two).
				Build().(*ctree),
			" (root (= (a (b (c (d (+ (1 (2)))))))))",
			" (+ (1 (2)))",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.expected != tt.input.SExpr(r) {
				t.Errorf(`
Expected %s
Received '%s'				
				`, tt.expected, tt.input.SExpr(r))
			}
			if tt.expected2 != tt.input.SExpr(p) {
				t.Errorf(`
2. 
Expected %s
Received '%s'				
				`, tt.expected2, tt.input.SExpr(p))
			}
		})
	}
}

//' = ( + 1 2 )'
