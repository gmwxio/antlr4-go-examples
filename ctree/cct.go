package ctree

import (
	"bytes"
	"reflect"

	"fmt"
)

type CallbackFunc func(depth int, node INode) (more bool)

type Tree interface {
	Name() string
	// Returns the root Inode of the tree
	Root() INode
	// Adds a child to the provided parent.
	// Returns true, nil if added successfully,
	// false, current parent if
	// the Inode already exists in the tree
	Add(parent INode, child INode) (success bool, otherParent INode)
	// Get the parent of the provided Inode
	Parent(child INode) (parent INode)
	// Depth first walk of the tree calling the provided function for each Inode visited
	Walk(CallbackFunc)

	// Remove(Inode INode) (bool)
	// GoWalk() chan<- Visit

	// Children of parent
	Children(parent INode) []INode

	Size() int
	PathAsPosition(node INode) []int
	CommonAncestor(a, b INode) INode

	// Contains
	Contains(INode) bool

	TreeString() string
	DebugTreeString() string
	SExpr(start INode) string
	//
}

// INode, an interface{} type store in the tree.
// Tider then writing interface{} everywhere
type INode interface {
}

// Fluent API for building a tree
type Builder interface {
	Add(n INode) Builder
	Down() Builder
	Up() Builder
	Build() Tree
	Current() INode
}

// Create a new tree with a hidden root and return the root and the tree
func NewTree(name string, root INode) Tree {
	if reflect.TypeOf(root).Kind() != reflect.Ptr {
		panic("Root Inode is not a pointer")
	}
	t := &ctree{
		name:    name,
		p2:      make(map[INode][]INode),
		c2:      make(map[INode]INode),
		root:    root,
		ptrOnly: true,
	}
	t.c2[t.root] = nil
	return t
}

// Creates a tree that allow non pointer Inode.
// Client beware, do not mutate object or pass non keyable objects
func NewTree_MutableINodes(name string, root INode) (Tree, INode) {
	t := &ctree{
		name:    name,
		p2:      make(map[INode][]INode),
		c2:      make(map[INode]INode),
		root:    root,
		ptrOnly: false,
	}
	t.c2[(INode)(t.root)] = nil
	return (Tree)(t), t.root
}

// The internal structure of the centralised tree
type ctree struct {
	name string
	// Map of parent to slice of children
	p2 map[INode][]INode
	// Map of child to parent
	c2 map[INode]INode
	// Hidden root Inode of the tree
	root    INode
	ptrOnly bool
}

func (t *ctree) Name() string {
	return t.name
}

// Add a child to a parent and create the p2c and c2p pointers
// If the child already exists anywhere in the tree the operations fails and the
// existing parent is returned with the boolean status
func (t *ctree) Add(p INode, c INode) (bool, INode) {
	if t.ptrOnly {
		if reflect.TypeOf(c) == nil {
			panic(fmt.Sprintf("Child kind is nil %T", c))
		}
		if reflect.TypeOf(p) == nil {
			panic(fmt.Sprintf("Parent kind is nil %T", c))
		}
		if reflect.TypeOf(c).Kind() != reflect.Ptr {
			panic("Child Inode is not a pointer")
		}
		if reflect.TypeOf(p).Kind() != reflect.Ptr {
			panic("Parent Inode is not a pointer")
		}
	}
	if orginal, exist := t.c2[c]; exist {
		return false, orginal
	}
	t.p2[p] = append(t.p2[p], c)
	t.c2[c] = p
	return true, nil
}

func (t *ctree) Parent(c INode) INode {
	return t.c2[c]
}

func (t *ctree) Contains(c INode) bool {
	var exist bool
	_, exist = t.c2[c]
	return exist
}

func (t *ctree) Size() int {
	return len(t.c2)
}
func (t *ctree) PathAsPosition(node INode) []int {
	depth := 0
	p, ex := t.c2[node]
	if !ex {
		return []int{-1}
	}
	for p != nil {
		p = t.c2[p]
		depth++
	}
	ret := make([]int, depth, depth)
	for node != t.Root() {
		p := t.Parent(node)
		ks := t.Children(p)
		for i, k := range ks {
			if node == k {
				depth--
				ret[depth] = i
				break
			}
		}
		node = p
	}
	return ret
}
func (t *ctree) CommonAncestor(a, b INode) INode {
	if a == b {
		return a
	}
	ap := t.PathAsPosition(a)
	bp := t.PathAsPosition(b)
	l := 0
	if len(ap) < len(bp) {
		l = len(ap)
	} else {
		l = len(bp)
	}
	path := make([]int, 0, l)
	for i := 0; i < l; i++ {
		if ap[i] != bp[i] {
			break
		}
		path = append(path, ap[i])
	}
	if len(path) == 0 {
		return t.Root()
	}
	cur := t.Root()
	for _, idx := range path {
		cur = t.Children(cur)[idx]
	}
	return cur
}

func (t *ctree) Children(p INode) []INode {
	return t.p2[p]
}

func (t *ctree) String() string {
	if t.name == "" {
		return treeString(t, appendStrNormal)
	}
	return t.name + "::" + treeString(t, appendStrNormal)
}
func (t *ctree) TreeString() string {
	if t.name == "" {
		return treeString(t, appendStrNormal)
	}
	return t.name + "::" + treeString(t, appendStrNormal)
}
func (t *ctree) DebugTreeString() string {
	if t.name == "" {
		return treeString(t, appendStrDebug)
	}
	return t.name + "::" + treeString(t, appendStrDebug)
}

func treeString(t *ctree, appendStr func(buf *bytes.Buffer, node INode)) string {
	var buf bytes.Buffer
	currDepth := 0
	first := true
	t.Walk(func(d int, n INode) (more bool) {
		if d > currDepth {
			buf.WriteString("(")
			appendStr(&buf, n)
			first = false
			currDepth = d
		} else if d < currDepth {
			for d < currDepth {
				buf.WriteString(")")
				currDepth--
			}
			appendStr(&buf, n)
			first = false
		} else {
			//					fmt.Printf("%s %d %d %s\n", Inode.Name, d, currDepth, first)
			if !first {
				buf.WriteString(",")
			} else {
				first = false
			}
			appendStr(&buf, n)
		}
		return true
	})
	for 0 < currDepth {
		buf.WriteString(")")
		currDepth--
	}
	return buf.String()
}

func appendStrDebug(buf *bytes.Buffer, node INode) {
	buf.WriteString(fmt.Sprintf("%+#v", node))
}
func appendStrNormal(buf *bytes.Buffer, node INode) {
	switch node := node.(type) {
	case string:
		buf.WriteString(node)
	case fmt.Stringer:
		buf.WriteString(node.String())
	default:
		buf.WriteString(fmt.Sprintf("%v", node))
	}
}

func (t *ctree) SExpr(startNode INode) string {
	buf := bytes.Buffer{}
	var iter Iterator
	if startNode == nil {
		iter = NewPreOrderTreeIterator(t, t.Root())
	} else {
		iter = NewPreOrderTreeIterator(t, startNode)
	}
	var start bool
	for iter.HasNext() {
		e := iter.Next()
		if _, ok := e.(*Eof); ok {
			continue
		}
		if _, ok := e.(*Down); ok {
			start = true
			buf.WriteString(fmt.Sprintf(" %s", e))
		} else if _, ok := e.(*Up); ok {
			buf.WriteString(fmt.Sprintf("%s", e))
		} else if start {
			buf.WriteString(fmt.Sprintf("%s", e))
			start = false
		} else {
			buf.WriteString(fmt.Sprintf(" %s", e))
		}
	}
	return buf.String()
}

// A depth first walk of the tree calling the provided function at each Inode
func (t *ctree) Walk(f CallbackFunc) {
	for _, o := range t.p2[t.root] {
		walk(t, o, 0, f)
	}
}

func walk(t *ctree, Inode INode, depth int, f CallbackFunc) {
	more := f(depth, Inode)
	if !more {
		return
	}
	for _, o := range t.p2[Inode] {
		walk(t, o, depth+1, f)
	}
}

// Return the root of the tree
func (t *ctree) Root() INode {
	return t.root
}

// Builder's state
type builder struct {
	tree Tree
	curr INode
	last INode
}

func BuildTree(name string, root INode) Builder {
	t := NewTree(name, root)
	b := &builder{t, root, nil}
	return b
}
func BuildTree_MutableNodes(name string, root INode) Builder {
	t, r := NewTree_MutableINodes(name, root)
	b := &builder{t, r, nil}
	return b
}
func (b *builder) Add(n INode) Builder {
	ok, _ := b.tree.Add(b.curr, n)
	if !ok {
		panic("Can't replace in a builder")
	}
	b.last = n
	return b
}
func (b *builder) Current() INode {
	return b.last
}
func (b *builder) Down() Builder {
	//	if b.last != nil {
	b.curr = b.last
	//	}
	return b
}
func (b *builder) Up() Builder {
	b.curr = b.tree.Parent(b.curr)
	b.last = b.curr
	return b
}
func (b *builder) Build() Tree {
	return b.tree
}
