parser grammar ExprWalker;

options {
  tokenVocab = ExprParser;
}

@header {
	import 	"strings"
	import 	"github.com/wxio/antlr4-go-examples/ctree"
}

@members {	
type Atom interface {
	Value() interface{}
	Add(b Atom) Atom
	Exp(b Atom) Atom
	// Mult(b Atom) Atom
	// Div(b Atom) Atom
	Neg() Atom
	Percent() Atom
}

type BinaryEval interface {
	Eval(Atom, Atom) Atom
}
type UnaryEval interface {
	Eval(Atom) Atom
}
type LazyEval interface {
	EvalLazy(strictArgs []Atom, lazyArgs []ctree.Tree) Atom	
}
type StrictEval interface {
	EvalStrict(args []Atom) Atom	
}

type Namer interface {
	GetName() string
}

func (p *ExprWalker) isIf(f antlr.Token) bool {
	n := f.(Namer).GetName()
	ret := strings.ToLower(n) == "if"
	return ret
}

}

start returns[Atom atom] 
	: DOWN ROOT a=startexpr {$atom=$a.atom} UP EOF ;

startexpr returns[Atom atom]
	: DOWN a=expr {$atom=$a.atom} UP ;

expr returns[Atom atom]
	: e=BINARYEXPR DOWN l=expr r=expr UP				{ $atom=$e.(BinaryEval).Eval($l.atom, $r.atom) }
	| e=UNARYEXPR DOWN l=expr UP						{ $atom=$e.(UnaryEval).Eval($l.atom) }
	| BLOCKORREFEXPR DOWN a=ref b+=ref+ UP				{ fmt.Printf("%v %v\n", $a.atom, localctx.(*ExprContext).GetB()) }
	| BLOCKORREFEXPR DOWN ex=expr UP					{ $atom=$ex.atom }
	| e=ERRORATOM										{ $atom=$e.(Atom) }
	| e=FLOATATOM										{ $atom=$e.(Atom) }
	| e=INTEGERATOM										{ $atom=$e.(Atom) }
	| e=BOOLATOM										{ $atom=$e.(Atom) }
	| e=STRINGATOM										{ $atom=$e.(Atom) }
	| re=ref											{ $atom=$re.atom }
;

ref returns[Atom atom]
@init { var ret Atom }
	: A1REF			{ $atom = ret }
	| NAMEDREF		{ $atom = ret }
	| RANGEREF		{ $atom = ret }
	| INTERSECTREF	{ $atom = ret }
	| ENHANCEDREF	{ $atom = ret }
	| HRANGEREF		{ $atom = ret }
	| f=FUNCEXPR {p.isIf($f)}? DOWN a=expr b=lazyarg UP	
	{ 
		$atom = $f.(LazyEval).EvalLazy([]Atom{$a.atom}, []ctree.Tree{$b.tree} ) 
	}
	| f=FUNCEXPR {p.isIf($f)}? DOWN a=expr b=lazyarg c=lazyarg UP	
	{ 
		$atom = $f.(LazyEval).EvalLazy([]Atom{$a.atom}, []ctree.Tree{$b.tree, $c.tree} ) 
	}
	| f=FUNCEXPR DOWN e+=expr+ UP
	{
		exprs := localctx.(*RefContext).GetE()
		atoms := make([]Atom, len(exprs))
		for i, e := range exprs {
			atoms[i] = e.GetAtom()
		}
		$atom = localctx.(*RefContext).GetF().(StrictEval).EvalStrict(atoms)
	}	
	| f=FUNCEXPR 	
	{ 
		$atom = $f.(StrictEval).EvalStrict([]Atom{})
	}	
;

lazyarg returns[ctree.Tree tree]
@init { builder := ctree.BuildTree("", &struct{}{}) }
@after{ $tree = builder.Build() }
	: lazyexpr[ builder ]
;

// TODO could save the start and end index of the stream and use those.
lazyexpr[ctree.Builder builder]
//	: DOWN le=lazyexpr+ UP					{ $node=$le.node }
	: n=~(DOWN|UP)			{ builder.Add($n) } 
		DOWN 				{ builder.Down() } 
		lazyexpr[builder]+ 
		UP 					{ builder.Up() }
	| n=~(DOWN|UP) 			{ builder.Add($n) } 	
;
