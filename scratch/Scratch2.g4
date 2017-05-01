grammar Scratch2;

options {
  language = Go;
  newCallback = true;
  runtimeImport = 'github.com/wxio/antlr4-go';
}

start : a EOF # Start1
;

a  
  : t=('1'|'0') b? a? # A1
;

b
  : t='x' # B1
  | t='y' # B1
  | c     # B2
;

c : d? 'c' # C1
;

d : 'd';