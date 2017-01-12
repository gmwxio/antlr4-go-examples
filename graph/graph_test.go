package graph

import (
	"fmt"
	"testing"

	"bytes"
)

type Nodie struct {
	Name string
}

func (n *Nodie) String() string {
	return n.Name
}

func TestGraph(t *testing.T) {
	a := &Nodie{"A"}
	b := &Nodie{"B"}
	c := &Nodie{"C"}
	d := &Nodie{"D"}
	g := BuildGraph(a).AddEdge(b).AddRoot(c).AddEdge(d).Build()
	if "[A C]" != fmt.Sprintf("%v", g.Roots()) {
		t.Errorf("Expect '[A C]' received '%v'", g.Roots())
	}
	r1c := g.Children(g.Roots()[0])
	if "[B]" != fmt.Sprintf("%v", r1c) {
		t.Errorf("Expect '[B]' received '%v'", r1c)
	}
	if "[]" != fmt.Sprintf("%v", g.Children(r1c[0])) {
		t.Errorf("expected '[]' received '%v'", g.Children(r1c[0]))
	}
	r2c := g.Children(g.Roots()[1])
	if "[D]" != fmt.Sprintf("%v", r2c) {
		t.Errorf("Expected '[D]' received '%v'", r2c)
	}
	if "[]" != fmt.Sprintf("%v", g.Children(r2c[0])) {
		t.Errorf("Expected '[]' received '%v'", g.Children(r2c[0]))
	}
}

func TestGraphBuilder(t *testing.T) {
	a := &Nodie{"A"}
	b := &Nodie{"B"}
	c := &Nodie{"C"}
	d := &Nodie{"D"}

	tests := []struct {
		name     string
		input    RootWalker
		expected string
	}{
		{
			"roots",
			BuildGraph(a).AddEdge(b).AddRoot(c).AddEdge(d).BFW(),
			"A C B D ",
		},
		{
			"level1",
			BuildGraph(a).AddEdge(b).AddEdge(c).AddEdge(d).BFW(),
			"A B C D ",
		},
		{
			"tree two levels",
			BuildGraph(a).AddEdge(b).Follow().AddEdge(c).Back().AddEdge(d).BFW(),
			"A B D C ",
		},
		{
			"deroot",
			BuildGraph(a).AddRoot(b).Back().AddEdge(d).Back().AddEdge(b).BFW(),
			"A D B ",
		},
		{
			"bfw!=topsort",
			BuildGraph(a).AddEdge(b).AddEdge(c).Follow().AddEdge(b).BFW(),
			"A B C ",
		},
		{
			"dfw = topsort",
			BuildGraph(a).AddEdge(b).AddEdge(c).Follow().AddEdge(b).DFW(),
			"B C A ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			coll := &collector{}
			tt.input.Walk(coll)
			if tt.expected != coll.String() {
				t.Errorf("TestGraphBuilder/%s expected:'%s' received:'%s'", tt.name, tt.expected, coll.String())
			}
		})
	}
}

type printer struct{}

func (p printer) Visit(n Node) {
	fmt.Printf("%s ", n)
}

type collector struct {
	bytes.Buffer
}

func (c *collector) Visit(n Node) {
	c.WriteString(fmt.Sprintf("%s ", n))
}

func (c *collector) String() string {
	return c.Buffer.String()
}
