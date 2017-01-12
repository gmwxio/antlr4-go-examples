parser grammar ExprParser;

tokens { 
	DOWN, 
	UP, 
	ROOT,
	
	FORMULAEXPR, 
	BINARYEXPR, 
	UNARYEXPR, 
	FUNCEXPR,
	BLOCKORREFEXPR,

	ARG,

	ERRORATOM,
	FLOATATOM,
	INTEGERATOM,
	BOOLATOM,
	STRINGATOM,

	A1REF,
	NAMEDREF,
	RANGEREF,
	INTERSECTREF,
	ENHANCEDREF,
	HRANGEREF
}

options {
  tokenVocab = ExprLexer;
}

//TODO => | 
 
start : cell EOF 
;

cell : EQ x+=WS* expr y+=WS*  
     | .*? 
;

startexpr 
    : expr EOF
;

expr
    : pre=('+'|'-') ws+=WS* e=expr                          #Prefix
    | e=expr ws+=WS* post='%'                               #Postfix
    | e1=expr l+=WS* t='^' r+=WS* e2=expr                   #Infix // | <assoc=right> expr '^' expr #EXP // Excel is left associative - really!!!
    | e1=expr l+=WS* t='*' r+=WS* e2=expr                   #Infix
    | e1=expr l+=WS* t='/' r+=WS* e2=expr                   #Infix
    | e1=expr l+=WS* t='+' r+=WS* e2=expr                   #Infix
    | e1=expr l+=WS* t='-' r+=WS* e2=expr                   #Infix
    | e1=expr l+=WS* t='&' r+=WS* e2=expr                   #Infix
    | e1=expr l+=WS* t=(GT|LT|GTE|LTE|EQ|NEQ) r+=WS* e2=expr
                                                            #Infix
//    | LP l+=WS* elems r+=WS* RP                 			#BlockOrRef
    | a=atom                                                #Atomic
    | ref                                                   #Reference
    | errorType                                             #Error
;

atom
    :   a=INT												#Integer
    |   a=FLOAT												#Float
    |   a=STRING											#String
    |   a=TRUE												#True
    |   a=FALSE												#False
;

ref 
    : r1=ref COLON r2=ref                       			#Range
    | r1=ref ws+=WS+ r2=ref                     			#Intersect
//    | r1=ref l+=WS* COMMA r+=WS* r2=ref         			#Union
    | name=fname LP l+=WS* args? r+=WS* RP      			#FunctionCall
	//todo empty func - down up per arg
    | LP l+=WS* elems r+=WS* RP                 			#BlockOrRef
    | a1=id_a1                                  			#A1Ref
    | row1=INT COLON row2=INT                   			#HRange
	| enhancedRef                               			#EHRef
;

id_a1
    : CHAR INT                 #NameOrA1
	| first=(CHAR | UNDERSCORE | BSLASH) rest+=(CHAR | UNDERSCORE | BSLASH | INT )*
	  // A1 or NamedRange
															#NamedRange
    | DOLLAR CHAR INT          								#A1ColumnAbs
    | CHAR DOLLAR INT          								#A1RowAbs
    | DOLLAR CHAR DOLLAR INT   								#A1Abs
;

enhancedRef
	: enhancedRefParts+
;

enhancedRefParts
	: (LSQRP firstRef+=enhancedRefContent* RSQRP)
;

enhancedRefContent
	:  CHAR|INT|UNDERSCORE|SLASH|DOT|COLON|LP|RP|HASH|WS | (LSQRP enhancedRefContent* RSQRP)
;

args
    : a=arg (l+=WS* COMMA r+=WS* b+=args)? 
;

arg
	: e=expr
;

fname 
    : (a+=CHAR | a+=UNDERSCORE) (a+=CHAR | a+=UNDERSCORE | a+=INT | a+=DOT )*  
;

elems
    : a=expr (l+=WS* COMMA r+=WS* b+=elems)?
;

//id_r1c1
//    : 
//      r=r1c1_row c=r1c1_col   ->  ^(CELLREF_R1C1<XLPN> $r $c)
//    |  r=r1c1_row             ->  ^(CELLREF_R1C1<XLPN> $r ALL)
//    |  c=r1c1_col             ->  ^(CELLREF_R1C1<XLPN> ALL $c)
//    | {input.LT(1).getText().equals("RC")}?=>
//      CHAR cidx=r1c1_idx ->        ^(CELLREF_R1C1<XLPN> THIS $cidx)
//    | //TODO check that ID is a valid named range (i.e. not a valid a1 ref
//      (a+=CHAR | a+=UNDERSCORE | a+=BSLASH) 
//      (a+=CHAR | a+=UNDERSCORE | a+=BSLASH | a+=INT )*   
//                                     -> ^(ATOM ID<XLPN>[$a,null])
//;

//r1c1_row
//    : {input.LT(1).getText().equals("R")}?=>
//      CHAR r1c1_idx -> r1c1_idx
//;

//r1c1_col
//    : {input.LT(1).getText().equals("C")}?=>
//      CHAR r1c1_idx -> r1c1_idx 
//    | {input.LT(1).getText().equals("C")}?=>
//      CHAR -> ^(THIS)
//;

//r1c1_idx
//    : LSQRP DASH INT RSQRP -> ^(RELATIVE NEG INT)
//    | LSQRP INT RSQRP -> ^(RELATIVE INT)
//    | INT -> ^(ABS INT)
//;

errorType
    : HASH a=CHAR b=(CHAR | SLASH | INT )* c=(QMARK | XMARK)
;

