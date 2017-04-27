// Generated from Scratch2.g4 by ANTLR 4.7.

package parser2 // Scratch2
import "github.com/wxio/antlr4-go"

// Scratch2Listener is a complete listener for a parse tree produced by Scratch2Parser.
type Scratch2Listener interface {
	antlr.ParseTreeListener

	Start1EntryListener
	Start1ExitListener

	A1EntryListener
	A1ExitListener

	B1EntryListener
	B1ExitListener

	B2EntryListener
	B2ExitListener

	C1EntryListener
	C1ExitListener
}

//
// Rules with unnamed alternatives
//

//
// Named alternatives
//
//
// From Rule 'start'
type Start1EntryListener interface {
	EnterStart1(c IStart1Context)
}
type Start1ExitListener interface {
	ExitStart1(c IStart1Context)
}

// From Rule 'a'
type A1EntryListener interface {
	EnterA1(c IA1Context)
}
type A1ExitListener interface {
	ExitA1(c IA1Context)
}

// From Rule 'b'
type B1EntryListener interface {
	EnterB1(c IB1Context)
}
type B1ExitListener interface {
	ExitB1(c IB1Context)
}

// From Rule 'b'
type B2EntryListener interface {
	EnterB2(c IB2Context)
}
type B2ExitListener interface {
	ExitB2(c IB2Context)
}

// From Rule 'c'
type C1EntryListener interface {
	EnterC1(c IC1Context)
}
type C1ExitListener interface {
	ExitC1(c IC1Context)
}
