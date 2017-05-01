// Generated from Scratch.g4 by ANTLR 4.7.

package parser // Scratch
import (
	"fmt"
	"reflect"
	"strconv"
)
import "github.com/wxio/antlr4-go"


// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa


var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 6, 26, 4, 
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 3, 2, 3, 2, 3, 2, 3, 3, 3, 3, 5, 3, 14, 
	10, 3, 3, 3, 5, 3, 17, 10, 3, 3, 3, 5, 3, 20, 10, 3, 3, 4, 3, 4, 5, 4, 
	24, 10, 4, 3, 4, 2, 2, 5, 2, 4, 6, 2, 2, 2, 26, 2, 8, 3, 2, 2, 2, 4, 13, 
	3, 2, 2, 2, 6, 23, 3, 2, 2, 2, 8, 9, 5, 4, 3, 2, 9, 10, 7, 2, 2, 3, 10, 
	3, 3, 2, 2, 2, 11, 14, 7, 3, 2, 2, 12, 14, 7, 4, 2, 2, 13, 11, 3, 2, 2, 
	2, 13, 12, 3, 2, 2, 2, 14, 16, 3, 2, 2, 2, 15, 17, 5, 6, 4, 2, 16, 15, 
	3, 2, 2, 2, 16, 17, 3, 2, 2, 2, 17, 19, 3, 2, 2, 2, 18, 20, 5, 4, 3, 2, 
	19, 18, 3, 2, 2, 2, 19, 20, 3, 2, 2, 2, 20, 5, 3, 2, 2, 2, 21, 24, 7, 5, 
	2, 2, 22, 24, 7, 6, 2, 2, 23, 21, 3, 2, 2, 2, 23, 22, 3, 2, 2, 2, 24, 7, 
	3, 2, 2, 2, 6, 13, 16, 19, 23,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'1'", "'0'", "'x'", "'y'",
}
var symbolicNames []string


var ruleNames = []string{
	"start", "a", "b",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type ScratchParser struct {
	*antlr.BaseParser
}

func NewScratchParser(input antlr.TokenStream) *ScratchParser {
	this := new(ScratchParser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Scratch.g4"

	return this
}

// ScratchParser tokens.
const (
	ScratchParserEOF = antlr.TokenEOF
	ScratchParserT__0 = 1
	ScratchParserT__1 = 2
	ScratchParserT__2 = 3
	ScratchParserT__3 = 4
)

// ScratchParser rules.
const (
	ScratchParserRULE_start = 0
	ScratchParserRULE_a = 1
	ScratchParserRULE_b = 2
)

type IStartContextInternal interface {
    // embed exported interface
    IStartContext

    //  ruleGetterDecl
    //  ruleListGetterDecl
    //  ruleListIndexedGetterDecl
    //  ruleContextDecls
    //  ruleContextListDecls
}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type IStartContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

    //Gets for labeled elements
    //tokenDecls	
    //tokenTypeDecls
    //tokenListDecls
    //attributeDecls
    //tokenGetterDecl
    //tokenListGetterDecl
    //tokenListIndexedGetterDecl

    // IsStartContext differentiates from other interfaces.
//copyStruct,GetRuleContext and ToStringTree  from embedded

//<if(dispatchMethods)>
//<dispatchMethods; separator="\n\n">
//<endif>

//<if(extensionMembers)>
//<extensionMembers; separator="\n\n">
//<endif>
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScratchParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScratchParserRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *StartContext) CopyFrom(ctx *StartContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers



//Begin AltLabelStructDecl


type Start1Context struct {
	*StartContext
}

func NewStart1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *Start1Context {
	var p = new(Start1Context)

	p.StartContext = NewEmptyStartContext()
	p.parser = parser
	p.CopyFrom(ctx.(*StartContext))

	return p
}

// Alts are separated into two interfaces.
// The intent is to allow two similar grammars to share Visitor or Listener implementations. 

type IStart1ContextInternal interface {
    IStart1Context
    //Gets for raw elements
    //  ruleGetterDecl
    A() IAContext  
    //  ruleListGetterDecl
    //  ruleListIndexedGetterDecl

    //  tokenGetterDecl
    EOF() antlr.TerminalNode
    //  tokenListGetterDecl
    //  tokenListIndexedGetterDecl


}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type IStart1Context interface {
    //Current rule
    IStartContext

    //Gets for labeled elements
    //  tokenDecls
    //  tokenTypeDecls
    //  tokenListDecls
    //  ruleContextDecls
    //  ruleContextListDecls
    //  attributeDecls

// TODO dispatchMethods (needed?)   
}

func (*Start1Context) IsStart1Context() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *Start1Context) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *Start1Context) A() IAContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAContext)
}

func (s *Start1Context) EOF() antlr.TerminalNode {
    return s.GetToken(ScratchParserEOF, 0)
}


func (s *Start1Context) EnterRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(Start1EntryListener); ok {
        listenerT.EnterStart1(s)
    }
}


func (s *Start1Context) ExitRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(Start1ExitListener); ok {
        listenerT.ExitStart1(s)
    }
}


func (s *Start1Context) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case Start1ContextVisitor:
		return t.VisitStart1(s, delegate, args)
	default:
		return delegate.VisitChildren(s, delegate, args)
	}
}

//END AltLabelStructDecl




func (p *ScratchParser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ScratchParserRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewStart1Context(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(6)
		p.A()
	}
	{
		p.SetState(7)
	    	p.Match(ScratchParserEOF)
	}



	return localctx
}


type IAContextInternal interface {
    // embed exported interface
    IAContext

    //  ruleGetterDecl
    //  ruleListGetterDecl
    //  ruleListIndexedGetterDecl
    //  ruleContextDecls
    //  ruleContextListDecls
}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type IAContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

    //Gets for labeled elements
    //tokenDecls	
    //tokenTypeDecls
    //tokenListDecls
    //attributeDecls
    //tokenGetterDecl
    //tokenListGetterDecl
    //tokenListIndexedGetterDecl

    // IsAContext differentiates from other interfaces.
//copyStruct,GetRuleContext and ToStringTree  from embedded

//<if(dispatchMethods)>
//<dispatchMethods; separator="\n\n">
//<endif>

//<if(extensionMembers)>
//<extensionMembers; separator="\n\n">
//<endif>
}

type AContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAContext() *AContext {
	var p = new(AContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScratchParserRULE_a
	return p
}

func (*AContext) IsAContext() {}

func NewAContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AContext {
	var p = new(AContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScratchParserRULE_a

	return p
}

func (s *AContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *AContext) CopyFrom(ctx *AContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *AContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers



//Begin AltLabelStructDecl


type A1Context struct {
	*AContext
	//TokenDecl
	t antlr.Token
}

func NewA1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *A1Context {
	var p = new(A1Context)

	p.AContext = NewEmptyAContext()
	p.parser = parser
	p.CopyFrom(ctx.(*AContext))

	return p
}

// Alts are separated into two interfaces.
// The intent is to allow two similar grammars to share Visitor or Listener implementations. 

type IA1ContextInternal interface {
    IA1Context
    //Gets for raw elements
    //  ruleGetterDecl
    B() IBContext  
    A() IAContext  
    //  ruleListGetterDecl
    //  ruleListIndexedGetterDecl

    //  tokenGetterDecl
    //  tokenListGetterDecl
    //  tokenListIndexedGetterDecl


}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type IA1Context interface {
    //Current rule
    IAContext

    //Gets for labeled elements
    //  tokenDecls
    GetT() antlr.Token
     
    //  tokenTypeDecls
    //  tokenListDecls
    //  ruleContextDecls
    //  ruleContextListDecls
    //  attributeDecls

// TODO dispatchMethods (needed?)   
}

func (*A1Context) IsA1Context() {}

//AltLabelStructDecl tokenDecls
func (s *A1Context) GetT() antlr.Token { return s.t }
func (s *A1Context) SetT(v antlr.Token) { s.t = v }


//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *A1Context) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *A1Context) B() IBContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBContext)
}

func (s *A1Context) A() IAContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAContext)
}


func (s *A1Context) EnterRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(A1EntryListener); ok {
        listenerT.EnterA1(s)
    }
}


func (s *A1Context) ExitRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(A1ExitListener); ok {
        listenerT.ExitA1(s)
    }
}


func (s *A1Context) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case A1ContextVisitor:
		return t.VisitA1(s, delegate, args)
	default:
		return delegate.VisitChildren(s, delegate, args)
	}
}

//END AltLabelStructDecl




func (p *ScratchParser) A() (localctx IAContext) {
	localctx = NewAContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ScratchParserRULE_a)
	var //TokenTypeDecl
	_la int


	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	localctx = NewA1Context(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(11)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScratchParserT__0:
		{
			p.SetState(9)
		    	var _m = p.Match(ScratchParserT__0)
		            localctx.(*A1Context).t = _m
		         
		}


	case ScratchParserT__1:
		{
			p.SetState(10)
		    	var _m = p.Match(ScratchParserT__1)
		            localctx.(*A1Context).t = _m
		         
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.SetState(14)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ScratchParserT__2 || _la == ScratchParserT__3 {
		{
			p.SetState(13)
			p.B()
		}

	}
	p.SetState(17)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == ScratchParserT__0 || _la == ScratchParserT__1 {
		{
			p.SetState(16)
			p.A()
		}

	}



	return localctx
}


type IBContextInternal interface {
    // embed exported interface
    IBContext

    //  ruleGetterDecl
    //  ruleListGetterDecl
    //  ruleListIndexedGetterDecl
    //  ruleContextDecls
    //  ruleContextListDecls
}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type IBContext interface {
	antlr.ParserRuleContext
	// GetParser returns the parser.
	GetParser() antlr.Parser

    //Gets for labeled elements
    //tokenDecls	
    //tokenTypeDecls
    //tokenListDecls
    //attributeDecls
    //tokenGetterDecl
    //tokenListGetterDecl
    //tokenListIndexedGetterDecl

    // IsBContext differentiates from other interfaces.
//copyStruct,GetRuleContext and ToStringTree  from embedded

//<if(dispatchMethods)>
//<dispatchMethods; separator="\n\n">
//<endif>

//<if(extensionMembers)>
//<extensionMembers; separator="\n\n">
//<endif>
}

type BContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBContext() *BContext {
	var p = new(BContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ScratchParserRULE_b
	return p
}

func (*BContext) IsBContext() {}

func NewBContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BContext {
	var p = new(BContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ScratchParserRULE_b

	return p
}

func (s *BContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *BContext) CopyFrom(ctx *BContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *BContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers



//Begin AltLabelStructDecl


type B1Context struct {
	*BContext
	//TokenDecl
	t antlr.Token
}

func NewB1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *B1Context {
	var p = new(B1Context)

	p.BContext = NewEmptyBContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BContext))

	return p
}

// Alts are separated into two interfaces.
// The intent is to allow two similar grammars to share Visitor or Listener implementations. 

type IB1ContextInternal interface {
    IB1Context
    //Gets for raw elements
    //  ruleGetterDecl
    //  ruleListGetterDecl
    //  ruleListIndexedGetterDecl

    //  tokenGetterDecl
    //  tokenListGetterDecl
    //  tokenListIndexedGetterDecl


}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type IB1Context interface {
    //Current rule
    IBContext

    //Gets for labeled elements
    //  tokenDecls
    GetT() antlr.Token
     
    //  tokenTypeDecls
    //  tokenListDecls
    //  ruleContextDecls
    //  ruleContextListDecls
    //  attributeDecls

// TODO dispatchMethods (needed?)   
}

func (*B1Context) IsB1Context() {}

//AltLabelStructDecl tokenDecls
func (s *B1Context) GetT() antlr.Token { return s.t }
func (s *B1Context) SetT(v antlr.Token) { s.t = v }


//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *B1Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *B1Context) EnterRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(B1EntryListener); ok {
        listenerT.EnterB1(s)
    }
}


func (s *B1Context) ExitRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(B1ExitListener); ok {
        listenerT.ExitB1(s)
    }
}


func (s *B1Context) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case B1ContextVisitor:
		return t.VisitB1(s, delegate, args)
	default:
		return delegate.VisitChildren(s, delegate, args)
	}
}

//END AltLabelStructDecl




func (p *ScratchParser) B() (localctx IBContext) {
	localctx = NewBContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ScratchParserRULE_b)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(21)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ScratchParserT__2:
		localctx = NewB1Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(19)
		    	var _m = p.Match(ScratchParserT__2)
		            localctx.(*B1Context).t = _m
		         
		}


	case ScratchParserT__3:
		localctx = NewB1Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(20)
		    	var _m = p.Match(ScratchParserT__3)
		            localctx.(*B1Context).t = _m
		         
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


