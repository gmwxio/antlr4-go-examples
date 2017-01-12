package graph

// graoh package implements a Digraph

import (
	"errors"
	"fmt"
	"reflect"
)

// Graph a digraph interface
type Graph interface {
	// Returns the roots digraph
	Roots() []Node
	// AddRoot
	AddRoot(n Node) (bool, error)
	// AddEdge create a directed edge between parent and child
	// Returns false is the edge already existed
	AddEdge(parent Node, child Node) (added bool) // TODO potential returns (success bool, cycle err)
	// Get the parents of the provided node
	Parents(child Node) (parents []Node)
	// Children of parent
	Children(parent Node) []Node

	// // Breadth first walk
	// BFW(listener func(Node))
	// // Depth first walk
	// DFW(listener func(Node))

	// Contains
	Contains(Node) bool
}

type RootWalker interface {
	Walk(v Visitor) Visitor
}

// Visitor interface
type Visitor interface {
	Visit(Node)
}

// Node an interface{} type store in the tree.
// Tider then writing interface{} everywhere
type Node interface{}

// Builder Fluent API for building a Graph
type GraphBuilder interface {
	AddRoot(n Node) GraphBuilder
	AddEdge(n Node) GraphBuilder
	Follow() GraphBuilder
	Back() GraphBuilder
	Build() Graph
	BFW() RootWalker
	DFW() RootWalker
}

// NewGraph create a new digraph
func NewGraph(root Node) Graph {
	if reflect.TypeOf(root).Kind() != reflect.Ptr {
		panic("Root node is not a pointer")
	}
	g := &graph{
		p2:    make(map[Node][]Node),
		c2:    make(map[Node][]Node),
		roots: make([]Node, 1, 1),
	}
	g.roots[0] = root
	g.p2[root] = nil
	return g
}

func NewEmptyGraph() Graph {
	g := &graph{
		p2:    make(map[Node][]Node),
		c2:    make(map[Node][]Node),
		roots: make([]Node, 0),
	}
	return g
}

// The internal structure of the centralised tree
type graph struct {
	// Map of parent to slice of children
	p2 map[Node][]Node
	// Map of child to slice of parents
	c2 map[Node][]Node
	// roots
	roots []Node
}

type graphArray struct {
	nodes []Node
	p2    map[Node]int
	c2    map[Node]int
	roots []int
}

func (g *graph) AddRoot(newr Node) (bool, error) {
	parents, exist := g.p2[newr]
	if !exist {
		g.roots = append(g.roots, newr)
		g.p2[newr] = nil
		return true, nil
	}
	if exist && parents == nil {
		return false, nil
	}
	return false, errors.New("Exists in Graph and is already a root, can't add a root multiple times")
}

func (g *graph) AddEdge(newp Node, newc Node) bool {
	if reflect.TypeOf(newc).Kind() != reflect.Ptr {
		panic("Child node is not a pointer")
	}
	if reflect.TypeOf(newp).Kind() != reflect.Ptr {
		panic("Parent node is not a pointer")
	}

	parents, exist := g.p2[newc]
	if !exist {
		g.p2[newp] = append(g.p2[newp], newc)
		g.c2[newc] = append(g.c2[newc], newp)
		return true
	}

	if exist && parents == nil { // its a root
		g.p2[newp] = append(g.p2[newp], newc)
		g.c2[newc] = append(g.c2[newc], newp)
		for i, root := range g.roots {
			if root == newc {
				// Delete
				// g.roots = append(g.roots[:idx], g.roots[idx+1:]...)
				// or without the potential memory leak
				copy(g.roots[i:], g.roots[i+1:])
				g.roots[len(g.roots)-1] = nil // or the zero value of T
				g.roots = g.roots[:len(g.roots)-1]
				return true
			}
		}
		panic(fmt.Sprintf("malformed tree. Parent should be a root but isn't"))
	}

	//TODO allow for multiple edge from p->c eg =A1+A1
	for _, kid := range g.p2[newp] {
		if kid == newc {
			return false // this edge already exists
		}
	}
	g.p2[newp] = append(g.p2[newp], newc)
	g.c2[newc] = append(g.c2[newc], newp)
	return true
}

func (g *graph) Parents(c Node) []Node {
	return g.c2[c]
}
func (g *graph) Children(p Node) []Node {
	return g.p2[p]
}

func (g *graph) Contains(c Node) bool {
	var exist bool
	_, exist = g.c2[c]
	return exist
}

// Return the roots of the graph
func (g *graph) Roots() []Node {
	return g.roots
}

type bfw struct {
	g       Graph
	visited map[Node]bool
}

// NewBFW breadth first walk
func NewBFW(g Graph) RootWalker {
	bfw := &bfw{
		g:       g,
		visited: make(map[Node]bool),
	}
	return bfw
}

func (bfw *bfw) Walk(v Visitor) Visitor {
	bfw.walklevel(bfw.g.Roots(), v.Visit)
	return v
}

func (bfw *bfw) walklevel(queued []Node, listener func(Node)) {
	nextlevel := make([]Node, 0)
	for _, n := range queued {
		listener(n)
		bfw.visited[n] = true
		for _, kid := range bfw.g.Children(n) {
			if !bfw.visited[kid] {
				nextlevel = append(nextlevel, kid)
			}
		}
	}
	if len(nextlevel) == 0 {
		return
	}
	bfw.walklevel(nextlevel, listener)
}

type colour int

const (
	white colour = iota
	grey
	black
)

type dfs struct {
	g      Graph
	colour map[Node]colour
	isDag  bool
}

// NewDFS depth first walker
func NewDFW(g Graph) RootWalker {
	dfs := &dfs{
		g:      g,
		colour: make(map[Node]colour),
		isDag:  true,
	}
	return dfs
}

func (dfs *dfs) Walk(v Visitor) Visitor {
	for _, n := range dfs.g.Roots() {
		dfs.walk(n, v.Visit)
	}
	return v
}

func (dfs *dfs) walk(n Node, listener func(Node)) {
	colour := dfs.colour[n]
	switch colour {
	case black:
		return
	case grey:
		dfs.isDag = false
		return
	case white:
		dfs.colour[n] = grey
		for _, kid := range dfs.g.Children(n) {
			dfs.walk(kid, listener)
		}
		dfs.colour[n] = black
		listener(n)
	}
}

// Builder's state
type graphbuilder struct {
	graph Graph
	curr  Node
	last  Node
	path  []Node
}

// BuildGraph creates a new fluent graphbuilder
func BuildGraph(root Node) GraphBuilder {
	g := NewGraph(root)
	path := make([]Node, 1, 1)
	path[0] = root
	b := &graphbuilder{g, root, root, path}
	return b
}

func (b *graphbuilder) AddRoot(n Node) GraphBuilder {
	b.graph.AddRoot(n)
	b.last = n
	b.Follow()
	return b
}
func (b *graphbuilder) AddEdge(n Node) GraphBuilder {
	b.graph.AddEdge(b.curr, n)
	b.last = n
	return b
}
func (b *graphbuilder) Follow() GraphBuilder {
	if b.curr == b.last {
		return b
	}
	b.path = append(b.path, b.curr)
	b.curr = b.last
	return b
}
func (b *graphbuilder) Back() GraphBuilder {
	b.curr = b.path[len(b.path)-1]
	b.path = b.path[:len(b.path)]
	b.last = b.curr
	return b
}
func (b *graphbuilder) Build() Graph {
	return b.graph
}
func (b *graphbuilder) BFW() RootWalker {
	return NewBFW(b.graph)
}
func (b *graphbuilder) DFW() RootWalker {
	return NewDFW(b.graph)
}
