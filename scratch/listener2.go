package main

import (
	"fmt"

	"github.com/wxio/antlr4-go-examples/scratch/parser"
	"github.com/wxio/antlr4-go-examples/scratch/parser2"

	antlr "github.com/wxio/antlr4-go"
)

type Scratch2Listener struct {
	s1 *ScratchListener
}

var _ parser2.Scratch2Listener = &Scratch2Listener{}

func (s *Scratch2Listener) VisitTerminal(node antlr.TerminalNode) {
}
func (s *Scratch2Listener) VisitErrorNode(node antlr.ErrorNode) {
}
func (s *Scratch2Listener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Printf(".")
}
func (s *Scratch2Listener) ExitEveryRule(ctx antlr.ParserRuleContext) {
}

// Only implemented as needed.

func (s *Scratch2Listener) EnterStart1(ctx parser2.IStart1Context) {
	s.s1.EnterStart1(ctx.(parser.IStartContext))
}
func (s *Scratch2Listener) ExitStart1(ctx parser2.IStart1Context) {
	s.s1.ExitStart1(ctx.(parser.IStartContext))
}

func (s *Scratch2Listener) EnterA1(ctx parser2.IA1Context) {
	s.s1.EnterA1(ctx.(parser.IA1Context))
}
func (s *Scratch2Listener) ExitA1(ctx parser2.IA1Context) {
	s.s1.ExitA1(ctx.(parser.IA1Context))
}

func (s *Scratch2Listener) EnterB1(ctx parser2.IB1Context) {
	s.s1.EnterB1(ctx.(parser.IB1Context))
}
func (s *Scratch2Listener) ExitB1(ctx parser2.IB1Context) {
	s.s1.ExitB1(ctx.(parser.IB1Context))
}

func (s *Scratch2Listener) EnterB2(ctx parser2.IB2Context) {
	fmt.Printf(">2 %T\n", ctx)
}
func (s *Scratch2Listener) ExitB2(ctx parser2.IB2Context) {
	fmt.Printf("<2 %T\n", ctx)
}

func (s *Scratch2Listener) EnterC1(ctx parser2.IC1Context) {
	fmt.Printf(">2 %T\n", ctx)
}
func (s *Scratch2Listener) ExitC1(ctx parser2.IC1Context) {
	fmt.Printf(">2 %T\n", ctx)
}
