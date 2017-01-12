lexer grammar ExprLexer;

//TODO ? & | { }
DQ  : '\"'  ;
DOT : '.'   ;
SQ  : '\''  ;
GTE : '>='  ;
LTE : '<='  ;
NEQ : '<>'  ;
PER : '%'   ;
EQ  : '='   ;
GT  : '>'   ;
LT  : '<'   ;
AMP : '&'   ;
HAT : '^'   ;
COMMA    : ','   ;
DASH     : '-'   ;
PLUS     : '+'   ;
SLASH    : '/'   ;
BSLASH   : '\\'   ;
STAR     : '*'   ;
LP       : '('   ;
RP       : ')'   ;
LSQRP    : '['   ;
RSQRP    : ']'   ;
DOLLAR   : '$'   ;
COLON    : ':'   ;
SEMI     : ';'   ;
HASH     : '#'   ;
QMARK    : '?'   ;
XMARK    : '!'   ;

TRUE : 'T' 'R' 'U' 'E';
FALSE : 'F' 'A' 'L' 'S' 'E';

CHAR  :   ('a'..'z'|'A'..'Z')+;
UNDERSCORE : '_';

//) ('a'..'z'|'A'..'Z'|'0'..'9'|'_'|ESC_SEQ)*
//    ;
//ID  :	DIGITS? ('a'..'z'|'A'..'Z'|'_') ('a'..'z'|'A'..'Z'|'0'..'9'|'_'|ESC_SEQ)*
//    ;

INT :	DIGITS ;
FLOAT
    :   ('0'..'9')+ '.' ('0'..'9')*
    |   '.' ('0'..'9')+
    ;
STRING
    :  '"' ( '"' '"' | ~('\\'|'"') )* '"'
    ;
WS  :   ( ' ' | '\t' | '\r' | '\n' | '\u00A0' )
    ;
fragment
DIGITS :  '0'..'9'+ ;
fragment
ESC_SEQ
    :   '\\' ('\\'|'@'|'['|']'|'.'|'#'|'+'|'-'|'!'|':'|' ')
    ;

