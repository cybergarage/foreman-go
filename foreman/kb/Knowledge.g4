/*------------------------------------------------------------------
 *
 * Foreman for Go
 *
 * Copyright (C) 2017 Satoshi Konno
 *
 * This is licensed under BSD-style license, see file COPYING.
 *
 *------------------------------------------------------------------*/

grammar Knowledge;

/*------------------------------------------------------------------
 * DNF (Disjunctive normal form) Rule
 *------------------------------------------------------------------*/

knowledge
   : (BS)? clauses (BE)?
   ;

clauses
	: clause (OR clause)*
	; 

clause
	: (BS)? formulas (BE)?
	; 

formulas
	: formula (AND formula)*
	; 

formula
	: (BS)? leftOperand operator rightOperand (BE)?
	; 

leftOperand
	: operand
	;

rightOperand
	: operand
	;

operand
	: literalOperand
	| variableOperand
	| functionOperand
	;

literalOperand
	: NUMBER
	| REAL
	;

variableOperand
	: IDENTIFIER
	;

functionOperand
	: functionName '(' parameter (',' parameter)* ')'
	;

functionName
	: IDENTIFIER
	;

parameter
	: IDENTIFIER
	| NUMBER
	| REAL
	;

operator
	: DEQ
	| NEQ
	| LT
	| LE
	| GT
	| GE	
	;

/*------------------------------------------------------------------
 * Tokens (Noun)
 *------------------------------------------------------------------*/

/*------------------------------------------------------------------
 * Single Tokens
 *------------------------------------------------------------------*/

ASTERISK
	: '*'
	;

BS
	: '('
	;

BE
	: ')'
	;

DEQ
	: '=='
	;

NEQ
	: '!='
	;

LT
	: '<'
	;
	
LE
	: '<='
	;
	
GT
	: '>'
	;

GE	
	: '>='
	;
	
AND
	: '&'
	;

OR
	: '|'
	;

COMMA
	: ','
	;

SEMICOLON
	: ';'
	;

/*------------------------------------------------------------------
* General Tokens
------------------------------------------------------------------*/

WS  :   ( ' '
        | '\t'
        | '\r'
        | '\n'
        )
    -> skip
	;
	
IDENTIFIER  
	: ([A-Za-z]) ([a-zA-Z0-9_\-.*])*
    ;

STRING  
	: '"' ('""'|~'"')* '"'
    ;

NUMBER 
	: ([+-])? ('0'..'9')+
	;

REAL
	:   ([+-])? ('0'..'9')+ '.' ('0'..'9')* EXPONENT?
	|   ([+-])? '.' ('0'..'9')+ EXPONENT?
	|   ([+-])? ('0'..'9')+ EXPONENT
	;


fragment
EXPONENT : ('e'|'E') ('+'|'-')? ('0'..'9')+ ;

fragment
HEX_DIGIT : ('0'..'9'|'a'..'f'|'A'..'F') ;

/*------------------------------------------------------------------
 * Fragment single Tokens
 *------------------------------------------------------------------*/

fragment 
A	: 'A'
	| 'a'
	;

fragment 
B	: 'B'
	| 'b'
	;

fragment 
C	: 'C'
	| 'c'
	;

fragment 
D	: 'D'
	| 'd'
	;

fragment 
E	: 'E'
	| 'e'
	;

fragment 
F	: 'F'
	| 'f'
	;

fragment 
G	: 'G'
	| 'g'
	;

fragment 
H	: 'H'
	| 'h'
	;

fragment 
I	: 'I'
	| 'i'
	;

fragment 
J	: 'J'
	| 'j'
	;

fragment 
K	: 'K'
	| 'k'
	;

fragment 
L	: 'L'
	| 'l'
	;

fragment 
M	: 'M'
	| 'm'
	;

fragment 
N	: 'N'
	| 'n'
	;

fragment 
O	: 'O'
	| 'o'
	;

fragment 
P	: 'P'
	| 'p'
	;

fragment 
Q	: 'Q'
	| 'q'
	;

fragment 
R	: 'R'
	| 'r'
	;

fragment 
S	: 'S'
	| 's'
	;

fragment 
T	: 'T'
	| 't'
	;

fragment 
U	: 'U'
	| 'u'
	;

fragment 
V	: 'V'
	| 'v'
	;

fragment 
W	: 'W'
	| 'w'
	;

fragment 
X	: 'X'
	| 'x'
	;

fragment 
Y	: 'Y'
	| 'y'
	;

fragment 
Z	: 'Z'
	| 'z'
	;