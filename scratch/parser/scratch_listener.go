// Generated from Scratch.g4 by ANTLR 4.7.

package parser // Scratch

import "github.com/wxio/antlr4-go"

// ScratchListener is a complete listener for a parse tree produced by ScratchParser.
type ScratchListener interface {
	antlr.ParseTreeListener

	Start1EntryListener
	Start1ExitListener

	A1EntryListener
	A1ExitListener

	B1EntryListener
	B1ExitListener
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
