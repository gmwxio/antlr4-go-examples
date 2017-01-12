package ctree

import (
	"fmt"
	"strings"
	"testing"
)

// Testing structure
type mys struct {
	s string
}

func (m mys) String() string {
	return m.s
}

func TestExample(t *testing.T) {
	r := &mys{}
	ct := NewTree("T", r)
	a := mys{"a"}
	ct.Add(r, &a)
	b := mys{"b"}
	c := mys{"c"}
	ct.Add(&a, &b)
	ct.Add(&a, &c)
	ct.Add(&a, &c) // ignored as c already in the tree
	ct.Walk(p)

	fmt.Println()

	slice := &myslice{s: []int{1, 2, 3, 4}}
	BuildTree_MutableNodes("T", &mys{}).
		Add(mys{"a"}).Down().
		/**/ Add(&mys{"b"}).
		/**/ Add(mys{"c"}).
		Add(slice).
		Build().Walk(p)
}

type myslice struct {
	s []int
}

func (m myslice) String() string {
	return fmt.Sprintf("%v", m.s)
}

// Function given to the walker for printing nodes
func p(d int, n INode) bool {
	s := strings.Repeat(" ", d)
	//	fmt.Println( reflect.TypeOf( n ) )
	switch v := n.(type) {
	case *mys:
		fmt.Printf("1)%s%v\n", s, v.s)
	default:
		fmt.Printf("*)%s%v - %T\n", s, n, n)
	}
	return true
}

func TestSExpr(t *testing.T) {
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
		{"simple", BuildTree("", r).Add(e).Down().Add(p).Add(o).Add(two).Up().Build().(*ctree), " (= (+ 1 2))"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rec := tt.input.SExpr(tt.input.Children(tt.input.Root())[0])
			if tt.expected != rec {
				t.Errorf(`
Expected %s
Received %s				
				`, tt.expected, rec)
			}
		})
	}
}
