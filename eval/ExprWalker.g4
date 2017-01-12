parser grammar ExprWalker;

options {
  tokenVocab = ExprParser;
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

}

start returns[Atom atom] 
	: DOWN ROOT DOWN a=expr {$atom=$a.atom} UP UP EOF ;

expr returns[Atom atom]
	: e=BINARYEXPR DOWN l=expr r=expr UP				{ $atom=$e.(BinaryEval).Eval($l.atom, $r.atom) }
	| e=UNARYEXPR DOWN l=expr UP						{ $atom=$e.(UnaryEval).Eval($l.atom) }
	| e=ERRORATOM										{ $atom=$e.(Atom) }
	| e=FLOATATOM										{ $atom=$e.(Atom) }
	| e=INTEGERATOM										{ $atom=$e.(Atom) }
	| e=BOOLATOM										{ $atom=$e.(Atom) }
	| e=STRINGATOM										{ $atom=$e.(Atom) }
;