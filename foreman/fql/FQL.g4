/*------------------------------------------------------------------
 *
 * Foreman for Go
 *
 * Copyright (C) 2017 Satoshi Konno
 *
 * This is licensed under BSD-style license, see file COPYING.
 *
 *------------------------------------------------------------------*/

grammar FQL;

/*------------------------------------------------------------------
 * query
 *------------------------------------------------------------------*/

fql
   : queryList
   ;

/*------------------------------------------------------------------
 * query
 *------------------------------------------------------------------*/

queryList
	: query (SEMICOLON query)*
	;	

query
	: insertQuery
	| setQuery
	| selectQuery
	| exportQuery
	| deleteQuery
	| analyzeQuery
	| executeQuery
	;

/*------------------------------------------------------------------
 * INSERT
 *------------------------------------------------------------------*/

insertQuery
	: INSERT INTO target ('(' columns ')')? VALUES '(' values ')'
	;

/*------------------------------------------------------------------
 * SET (Alias of 'INSERT INTO target VALUES')
 *------------------------------------------------------------------*/

setQuery
	: SET '(' values ')' INTO target
	;

/*------------------------------------------------------------------
 * SELECT
 *------------------------------------------------------------------*/

selectQuery
	: SELECT (ASTERISK|('(' columns ')')) FROM target (WHERE conditions)?
	;

/*------------------------------------------------------------------
 * EXPORT (Alias of 'SELECT *')
 *------------------------------------------------------------------*/

exportQuery
	: EXPORT target (WHERE conditions)?
	;

/*------------------------------------------------------------------
 * DELETE
 *------------------------------------------------------------------*/

deleteQuery
	: DELETE FROM target (WHERE conditions)?
	;

/*------------------------------------------------------------------
 * ANALYZE
 *------------------------------------------------------------------*/

analyzeQuery
	: ANALYZE column FROM target
	;

/*------------------------------------------------------------------
 * EXECUTE
 *------------------------------------------------------------------*/

executeQuery
	: EXECUTE target ('(' columns ')')? (VALUES '(' values ')')?
	;

/*------------------------------------------------------------------
 * Common Objects
 *------------------------------------------------------------------*/

target
	: IDENTIFIER
	;

value
	: IDENTIFIER
	| STRING
	| NUMBER
	| REAL
	;

values
	: value (',' value)*
	;

column
	: IDENTIFIER
	;

columns
	: column (',' column)*
	;

conditions
	: condition (AND condition)*
	;

condition
	: leftOperand operator rightOperand
	;

operator
	: DOUBLE_EQ
	| OP_LT
	| LE
	| GT
	| GE	
	| NOTEQ
	;

leftOperand
	: IDENTIFIER
	;

rightOperand
	: IDENTIFIER
	| STRING
	| NUMBER
	| REAL
	;

/*------------------------------------------------------------------
 * Tokens (Verb)
 *------------------------------------------------------------------*/

SELECT
	: S E L E C T
	;

INSERT
	: I N S E R T
	;

DELETE
	: D E L E T E
	;

SET
	: S E T
	;

ANALYZE
	: A N A L Y Z E
	;

EXPORT
	: E X P O R T
	;

EXECUTE
	: E X E C U T E
	;

INTO
	: I N T O
	;

FROM
	: F R O M
	;

VALUES
	: V A L U E S
	;

WHERE
	: W H E R E
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
	
SINGLE_EQ
	: '='
	;

DOUBLE_EQ
	: '=='
	;

OP_LT
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
	
NOTEQ
	: '!='
	;

AND
	: A N D
	| '&'
	;

OR
	: O R
	| '|'
	;

COMMA
	: ','
	;

SEMICOLON
	: ';'
	;

/*------------------------------------------------------------------
*
* General Tokens
*
------------------------------------------------------------------*/

WS  :   ( ' '
        | '\t'
        | '\r'
        | '\n'
        )
    -> skip
	;
	
IDENTIFIER  
	: [a-zA-Z_] [a-zA-Z_0-9]*
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