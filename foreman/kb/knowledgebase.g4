/*------------------------------------------------------------------
 *
 * Foreman for Go
 *
 * Copyright (C) 2017 Satoshi Konno
 *
 * This is licensed under BSD-style license, see file COPYING.
 *
 *------------------------------------------------------------------*/

grammar knowledgebase;

/*------------------------------------------------------------------
 * Rule
 *------------------------------------------------------------------*/

knowledgebase
   : (SB)? clauses (CB)?
   ;

clauses
	: clause (AND clause)*
	; 

clause
	: (SB)? formulas (CB)?
	; 

formulas
	: formula (OR formula)*
	; 

formula
	: (SB)? operand operator operand (CB)?
	; 

operand
	: NUMBER
	| REAL
	| IDENTIFIER
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

SB
	: '('
	;

CB
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
	: [a-zA-Z0-9_\-.*]+
    ;

STRING  
	: '"' ('""'|~'"')* '"'
    ;

NUMBER 
	: '0'..'9'+
	;

REAL
	:   ('0'..'9')+ '.' ('0'..'9')* EXPONENT?
	|   '.' ('0'..'9')+ EXPONENT?
	|   ('0'..'9')+ EXPONENT
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