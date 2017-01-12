package ctree

import (
	"fmt"
)

type stackIterator []Iterator

func (s *stackIterator) push(v Iterator) {
	*s = append(*s, v)
}
func (s *stackIterator) pop() Iterator {
	l := len(*s)
	res := (*s)[l-1]
	*s = (*s)[:l-1]
	return res
}
func (s *stackIterator) peek() Iterator {
	return (*s)[len(*s)-1]
}
func (s *stackIterator) empty() bool {
	return len(*s) == 0
}

type PreOrderTreeIterator struct {
	stack stackIterator
	tree  Tree
}

type Iterator interface {
	HasNext() bool
	Next() INode
}

type INodeIterator struct {
	list        []INode
	pos, length int
}

func (iter *INodeIterator) String() string {
	return fmt.Sprintf("INodeIterator pos %d len %d hasnext %v list %v", iter.pos, iter.length, iter.HasNext(), iter.list)
}

func (iter *INodeIterator) Next() INode {
	x := iter.list[iter.pos]
	iter.pos++
	return x
}

func (iter *INodeIterator) HasNext() bool {
	return iter.pos < iter.length
}

// Constructs a preorder enumeration with the tree
// to traverse, and the starting node */
func NewPreOrderTreeIterator(tree Tree, node INode) *PreOrderTreeIterator {
	iter := &PreOrderTreeIterator{}
	iter.tree = tree
	iter.stack = stackIterator{}

	if node == tree.Root() {
		list := []INode{down, node, up, eof}
		stack := &INodeIterator{list: list, pos: 0, length: 4}
		iter.stack.push(stack)
	} else {
		kids := iter.tree.Children(iter.tree.Parent(node))
		if len(kids) != 0 {
			nkids := make([]INode, len(kids)+3)
			copy(nkids[1:], kids)
			nkids[0] = down
			nkids[len(nkids)-2] = up
			nkids[len(nkids)-1] = eof
			kids = nkids
		}
		kidsIter := &INodeIterator{list: kids, pos: 0, length: len(kids)}
		for node != kidsIter.Next() {
		}
		kidsIter.pos--
		kidsIter.pos--
		kidsIter.list[kidsIter.pos] = down

		iter.stack.push(kidsIter)
	}

	return iter
}

// INMIND It would be handy to get traversal information UP DOWN etc
func (iter *PreOrderTreeIterator) HasNext() bool {
	return !iter.stack.empty() && iter.stack.peek().HasNext()
}

type Down struct{}
type Up struct{}
type Eof struct{}

func (*Down) String() string { return "(" }
func (*Up) String() string   { return ")" }
func (*Eof) String() string  { return "<EOF>" }

var down *Down
var up *Up
var eof *Eof

// func init() {
// 	down = Down{}
// 	up = Up{}
// }

func (iter *PreOrderTreeIterator) Next() INode {
	currIter := iter.stack.peek()
	node := currIter.Next()

	kids := iter.tree.Children(node)
	if len(kids) != 0 {
		nkids := make([]INode, len(kids)+2)
		copy(nkids[1:], kids)
		nkids[0] = down
		nkids[len(nkids)-1] = up
		kids = nkids
	}
	kidsIter := &INodeIterator{list: kids, pos: 0, length: len(kids)}

	if !currIter.HasNext() {
		iter.stack.pop()
	}
	if kidsIter.HasNext() {
		iter.stack.push(kidsIter)
	}
	return node
}

func (iter *PreOrderTreeIterator) getChildrenIterator(parent INode) Iterator {
	kids := iter.tree.Children(parent)
	return &INodeIterator{list: kids, pos: 0, length: len(kids)}
}
