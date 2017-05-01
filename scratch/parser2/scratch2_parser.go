// Generated from Scratch2.g4 by ANTLR 4.7.

package parser2 // Scratch2
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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 8, 35, 4, 
	2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 3, 2, 3, 2, 3, 
	2, 3, 3, 3, 3, 5, 3, 18, 10, 3, 3, 3, 5, 3, 21, 10, 3, 3, 4, 3, 4, 3, 4, 
	5, 4, 26, 10, 4, 3, 5, 5, 5, 29, 10, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 2, 
	2, 7, 2, 4, 6, 8, 10, 2, 3, 3, 2, 3, 4, 2, 34, 2, 12, 3, 2, 2, 2, 4, 15, 
	3, 2, 2, 2, 6, 25, 3, 2, 2, 2, 8, 28, 3, 2, 2, 2, 10, 32, 3, 2, 2, 2, 12, 
	13, 5, 4, 3, 2, 13, 14, 7, 2, 2, 3, 14, 3, 3, 2, 2, 2, 15, 17, 9, 2, 2, 
	2, 16, 18, 5, 6, 4, 2, 17, 16, 3, 2, 2, 2, 17, 18, 3, 2, 2, 2, 18, 20, 
	3, 2, 2, 2, 19, 21, 5, 4, 3, 2, 20, 19, 3, 2, 2, 2, 20, 21, 3, 2, 2, 2, 
	21, 5, 3, 2, 2, 2, 22, 26, 7, 5, 2, 2, 23, 26, 7, 6, 2, 2, 24, 26, 5, 8, 
	5, 2, 25, 22, 3, 2, 2, 2, 25, 23, 3, 2, 2, 2, 25, 24, 3, 2, 2, 2, 26, 7, 
	3, 2, 2, 2, 27, 29, 5, 10, 6, 2, 28, 27, 3, 2, 2, 2, 28, 29, 3, 2, 2, 2, 
	29, 30, 3, 2, 2, 2, 30, 31, 7, 7, 2, 2, 31, 9, 3, 2, 2, 2, 32, 33, 7, 8, 
	2, 2, 33, 11, 3, 2, 2, 2, 6, 17, 20, 25, 28,
}
var deserializer = antlr.NewATNDeserializer(nil)
var deserializedATN = deserializer.DeserializeFromUInt16(parserATN)

var literalNames = []string{
	"", "'1'", "'0'", "'x'", "'y'", "'c'", "'d'",
}
var symbolicNames []string


var ruleNames = []string{
	"start", "a", "b", "c", "d",
}
var decisionToDFA = make([]*antlr.DFA, len(deserializedATN.DecisionToState))

func init() {
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

type Scratch2Parser struct {
	*antlr.BaseParser
}

func NewScratch2Parser(input antlr.TokenStream) *Scratch2Parser {
	this := new(Scratch2Parser)

	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Scratch2.g4"

	return this
}

// Scratch2Parser tokens.
const (
	Scratch2ParserEOF = antlr.TokenEOF
	Scratch2ParserT__0 = 1
	Scratch2ParserT__1 = 2
	Scratch2ParserT__2 = 3
	Scratch2ParserT__3 = 4
	Scratch2ParserT__4 = 5
	Scratch2ParserT__5 = 6
)

// Scratch2Parser rules.
const (
	Scratch2ParserRULE_start = 0
	Scratch2ParserRULE_a = 1
	Scratch2ParserRULE_b = 2
	Scratch2ParserRULE_c = 3
	Scratch2ParserRULE_d = 4
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
	p.RuleIndex = Scratch2ParserRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Scratch2ParserRULE_start

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
    return s.GetToken(Scratch2ParserEOF, 0)
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




func (p *Scratch2Parser) Start() (localctx IStartContext) {
	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, Scratch2ParserRULE_start)

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
		p.SetState(10)
		p.A()
	}
	{
		p.SetState(11)
	    	p.Match(Scratch2ParserEOF)
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
	p.RuleIndex = Scratch2ParserRULE_a
	return p
}

func (*AContext) IsAContext() {}

func NewAContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AContext {
	var p = new(AContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Scratch2ParserRULE_a

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




func (p *Scratch2Parser) A() (localctx IAContext) {
	localctx = NewAContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, Scratch2ParserRULE_a)
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
	p.SetState(13)

	var _lt = p.GetTokenStream().LT(1)

	    localctx.(*A1Context).t = _lt
	 

	_la = p.GetTokenStream().LA(1)

	if !(_la == Scratch2ParserT__0 || _la == Scratch2ParserT__1) {
		var _ri = p.GetErrorHandler().RecoverInline(p)
	        localctx.(*A1Context).t = _ri
	     

	} else {
	    p.GetErrorHandler().ReportMatch(p)
		p.Consume()
	}
	p.SetState(15)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if (((_la) & -(0x1f+1)) == 0 && ((1 << uint(_la)) & ((1 << Scratch2ParserT__2) | (1 << Scratch2ParserT__3) | (1 << Scratch2ParserT__4) | (1 << Scratch2ParserT__5))) != 0) {
		{
			p.SetState(14)
			p.B()
		}

	}
	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == Scratch2ParserT__0 || _la == Scratch2ParserT__1 {
		{
			p.SetState(17)
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
	p.RuleIndex = Scratch2ParserRULE_b
	return p
}

func (*BContext) IsBContext() {}

func NewBContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BContext {
	var p = new(BContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Scratch2ParserRULE_b

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


type B2Context struct {
	*BContext
}

func NewB2Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *B2Context {
	var p = new(B2Context)

	p.BContext = NewEmptyBContext()
	p.parser = parser
	p.CopyFrom(ctx.(*BContext))

	return p
}

// Alts are separated into two interfaces.
// The intent is to allow two similar grammars to share Visitor or Listener implementations. 

type IB2ContextInternal interface {
    IB2Context
    //Gets for raw elements
    //  ruleGetterDecl
    C() ICContext  
    //  ruleListGetterDecl
    //  ruleListIndexedGetterDecl

    //  tokenGetterDecl
    //  tokenListGetterDecl
    //  tokenListIndexedGetterDecl


}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type IB2Context interface {
    //Current rule
    IBContext

    //Gets for labeled elements
    //  tokenDecls
    //  tokenTypeDecls
    //  tokenListDecls
    //  ruleContextDecls
    //  ruleContextListDecls
    //  attributeDecls

// TODO dispatchMethods (needed?)   
}

func (*B2Context) IsB2Context() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *B2Context) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *B2Context) C() ICContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICContext)
}


func (s *B2Context) EnterRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(B2EntryListener); ok {
        listenerT.EnterB2(s)
    }
}


func (s *B2Context) ExitRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(B2ExitListener); ok {
        listenerT.ExitB2(s)
    }
}


func (s *B2Context) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case B2ContextVisitor:
		return t.VisitB2(s, delegate, args)
	default:
		return delegate.VisitChildren(s, delegate, args)
	}
}

//END AltLabelStructDecl




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




func (p *Scratch2Parser) B() (localctx IBContext) {
	localctx = NewBContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, Scratch2ParserRULE_b)

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

	p.SetState(23)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case Scratch2ParserT__2:
		localctx = NewB1Context(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(20)
		    	var _m = p.Match(Scratch2ParserT__2)
		            localctx.(*B1Context).t = _m
		         
		}


	case Scratch2ParserT__3:
		localctx = NewB1Context(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(21)
		    	var _m = p.Match(Scratch2ParserT__3)
		            localctx.(*B1Context).t = _m
		         
		}


	case Scratch2ParserT__4, Scratch2ParserT__5:
		localctx = NewB2Context(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(22)
			p.C()
		}



	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}


	return localctx
}


type ICContextInternal interface {
    // embed exported interface
    ICContext

    //  ruleGetterDecl
    //  ruleListGetterDecl
    //  ruleListIndexedGetterDecl
    //  ruleContextDecls
    //  ruleContextListDecls
}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type ICContext interface {
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

    // IsCContext differentiates from other interfaces.
//copyStruct,GetRuleContext and ToStringTree  from embedded

//<if(dispatchMethods)>
//<dispatchMethods; separator="\n\n">
//<endif>

//<if(extensionMembers)>
//<extensionMembers; separator="\n\n">
//<endif>
}

type CContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCContext() *CContext {
	var p = new(CContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Scratch2ParserRULE_c
	return p
}

func (*CContext) IsCContext() {}

func NewCContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CContext {
	var p = new(CContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Scratch2ParserRULE_c

	return p
}

func (s *CContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *CContext) CopyFrom(ctx *CContext) {
	s.BaseParserRuleContext.CopyFrom(ctx.BaseParserRuleContext)
}

func (s *CContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
//extensionMembers



//Begin AltLabelStructDecl


type C1Context struct {
	*CContext
}

func NewC1Context(parser antlr.Parser, ctx antlr.ParserRuleContext) *C1Context {
	var p = new(C1Context)

	p.CContext = NewEmptyCContext()
	p.parser = parser
	p.CopyFrom(ctx.(*CContext))

	return p
}

// Alts are separated into two interfaces.
// The intent is to allow two similar grammars to share Visitor or Listener implementations. 

type IC1ContextInternal interface {
    IC1Context
    //Gets for raw elements
    //  ruleGetterDecl
    D() IDContext  
    //  ruleListGetterDecl
    //  ruleListIndexedGetterDecl

    //  tokenGetterDecl
    //  tokenListGetterDecl
    //  tokenListIndexedGetterDecl


}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type IC1Context interface {
    //Current rule
    ICContext

    //Gets for labeled elements
    //  tokenDecls
    //  tokenTypeDecls
    //  tokenListDecls
    //  ruleContextDecls
    //  ruleContextListDecls
    //  attributeDecls

// TODO dispatchMethods (needed?)   
}

func (*C1Context) IsC1Context() {}

//AltLabelStructDecl tokenDecls

//AltLabelStructDecl tokenTypeDecls

//AltLabelStructDecl tokenListDecls

//AltLabelStructDecl ruleContextDecls

//AltLabelStructDecl ruleContextListDecls

//AltLabelStructDecl attributeDecls

//getRuleContext
func (s *C1Context) GetRuleContext() antlr.RuleContext {
	return s
}

//getters
func (s *C1Context) D() IDContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDContext)
}


func (s *C1Context) EnterRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(C1EntryListener); ok {
        listenerT.EnterC1(s)
    }
}


func (s *C1Context) ExitRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(C1ExitListener); ok {
        listenerT.ExitC1(s)
    }
}


func (s *C1Context) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case C1ContextVisitor:
		return t.VisitC1(s, delegate, args)
	default:
		return delegate.VisitChildren(s, delegate, args)
	}
}

//END AltLabelStructDecl




func (p *Scratch2Parser) C() (localctx ICContext) {
	localctx = NewCContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, Scratch2ParserRULE_c)
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

	localctx = NewC1Context(p, localctx)
	p.EnterOuterAlt(localctx, 1)
	p.SetState(26)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)


	if _la == Scratch2ParserT__5 {
		{
			p.SetState(25)
			p.D()
		}

	}
	{
		p.SetState(28)
	    	p.Match(Scratch2ParserT__4)
	}



	return localctx
}


type IDContextInternal interface {
    // embed exported interface
    IDContext

    //  ruleGetterDecl
    //  ruleListGetterDecl
    //  ruleListIndexedGetterDecl
    //  ruleContextDecls
    //  ruleContextListDecls
}

// This can to be the 'Exported' interface (put in non export if it turn out to be an issue)
type IDContext interface {
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

    // IsDContext differentiates from other interfaces.
//copyStruct,GetRuleContext and ToStringTree  from embedded

//<if(dispatchMethods)>
//<dispatchMethods; separator="\n\n">
//<endif>

//<if(extensionMembers)>
//<extensionMembers; separator="\n\n">
//<endif>
}

type DContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDContext() *DContext {
	var p = new(DContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = Scratch2ParserRULE_d
	return p
}

func (*DContext) IsDContext() {}

func NewDContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DContext {
	var p = new(DContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = Scratch2ParserRULE_d

	return p
}

func (s *DContext) GetParser() antlr.Parser { return s.parser }

//StructDecl tokenDecls

//StructDecl tokenTypeDecls

//StructDecl tokenListDecls

//StructDecl ruleContextDecls

//StructDecl ruleContextListDecls

//StructDecl attributeDecls

// Getters

//provideCopyFrom
func (s *DContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

//dispatchMethods
func (s *DContext) EnterRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(DEntryListener); ok {
        listenerT.EnterD(s)
    }
}

func (s *DContext) ExitRule(listener antlr.ParseTreeListener) {
    if listenerT, ok := listener.(DExitListener); ok {
        listenerT.ExitD(s)
    }
}

func (s *DContext) Visit(delegate antlr.ParseTreeVisitor, args ...interface{}) (result interface{}) {
	switch t := delegate.(type) {
	case DContextVisitor:
		return t.VisitD(s, delegate, args)
	default:
		return delegate.VisitChildren(s, delegate, args)
	}
}

//extensionMembers



func (p *Scratch2Parser) D() (localctx IDContext) {
	localctx = NewDContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, Scratch2ParserRULE_d)

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

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(30)
	    	p.Match(Scratch2ParserT__5)
	}



	return localctx
}


