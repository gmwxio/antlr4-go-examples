grammar Scratch;

options {
  language = Go;
  newCallback = true;
  runtimeImport = 'github.com/wxio/antlr4-go';  
}

start 
  : a EOF # Start1
;

a  
  : t=('1'|'0') b? a? # A1
;

b
  : t='x' # B1
  | t='y' # B1
;