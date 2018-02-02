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
 * Statement
 *------------------------------------------------------------------*/

fql
   : statement_list
   ;

/*------------------------------------------------------------------
 * Statement
 *------------------------------------------------------------------*/

statement_list
	: statement (SEMICOLON statement)*
	;	

statement
	: insert_stmt
	| set_stmt
	| select_stmt
	| export_stmt
	| del_stmt
	;

/*------------------------------------------------------------------
 * INSERT
 *------------------------------------------------------------------*/

insert_stmt
	: INSERT INTO target VALUES '(' values ')' # InsertQuery
	;

/*------------------------------------------------------------------
 * SET (Alias of 'INSERT INTO target VALUES')
 *------------------------------------------------------------------*/

set_stmt
	: SET '(' values ')' INTO target # SetQuery
	;

/*------------------------------------------------------------------
 * SELECT
 *------------------------------------------------------------------*/

select_stmt
	: SELECT ASTERISK FROM target # SelectQuery
	;

/*------------------------------------------------------------------
 * EXPORT (Alias of 'SELECT *')
 *------------------------------------------------------------------*/

export_stmt
	: EXPORT target # ExportQuery
	;

/*------------------------------------------------------------------
 * DELETE
 *------------------------------------------------------------------*/

del_stmt
	: DELETE value FROM target # DeleteQuery
	;

/*------------------------------------------------------------------
 * Object
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

EXPORT
	: E X P O R T
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