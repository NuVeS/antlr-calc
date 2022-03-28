// test.g4
grammar Calculator;

// Tokens
MUL: '*';
DIV: '/';
ADD: '+';
SUB: '-';
FACT: '!';
OP_BRACETS: '(';
CL_BRACETS: ')';
POW: '^';
NUMBER: [0-9]+;
WHITESPACE: [ \r\n\t]+ -> skip;

// Rules
start : expression EOF;

expression
   : OP_BRACETS expression CL_BRACETS   # Bracet
   | expression FACT                    # FACTORIAL
   | expression POW expression          # POWER
   | expression op=('*'|'/') expression # MulDiv
   | expression op=('+'|'-') expression # AddSub
   | NUMBER                             # Number
   ;