package lox

import "fmt"

type TokenType int64

func (tokenType TokenType) String() string {
    return tokeTypeToString[tokenType]
}

const (
    // Single-character tokens.
    LEFT_PAREN TokenType = iota
    RIGHT_PAREN 
    LEFT_BRACE
    RIGHT_BRACE
    COMMA
    DOT
    MINUS
    PLUS
    SEMICOLON
    SLASH
    STAR

    // One or two character tokens.
    BANG
    BANG_EQUAL
    EQUAL
    EQUAL_EQUAL
    GREATER
    GREATER_EQUAL
    LESS
    LESS_EQUAL

    // Literals.
    IDENTIFIER
    STRING
    NUMBER

    // Keywords.
    AND
    CLASS
    ELSE
    FALSE
    FUN
    FOR
    IF
    NIL
    OR
    PRINT
    RETURN
    SUPER
    THIS
    TRUE
    VAR
    WHILE

    EOF
)

var tokeTypeToString = map[TokenType]string {
    LEFT_PAREN: "LEFT_PAREN",
    RIGHT_PAREN: "RIGHT_PAREN",
    LEFT_BRACE: "LEFT_BRACE",
    RIGHT_BRACE: "RIGHT_BRACE",
    COMMA: "COMMA",
    DOT: "DOT",
    MINUS: "MINUS",
    PLUS: "PLUS",
    SEMICOLON: "SEMICOLON",
    SLASH: "SLASH",
    STAR: "STAR",

    // One or two character tokens.
    BANG: "BANG",
    BANG_EQUAL: "BANG_EQUAL",
    EQUAL: "EQUAL",
    EQUAL_EQUAL: "EQUAL_EQUAL",
    GREATER: "GREATER",
    GREATER_EQUAL: "GREATER_EQUAL",
    LESS: "LESS",
    LESS_EQUAL: "LESS_EQUAL",

    // Literals.
    IDENTIFIER: "IDENTIFIER",
    STRING: "STRING",
    NUMBER: "NUMBER",

    // Keywords.
    AND: "AND",
    CLASS: "CLASS",
    ELSE: "ELSE",
    FALSE: "FALSE",
    FUN: "FUN",
    FOR: "FOR",
    IF: "IF",
    NIL: "NIL",
    OR: "OR",
    PRINT: "PRINT",
    RETURN: "RETURN",
    SUPER: "SUPER",
    THIS: "THIS",
    TRUE: "TRUE",
    VAR: "VAR",
    WHILE: "WHILE",

    EOF: "EOF",
}

var keywords = map[string]TokenType{
    "and":    AND,
    "class":  CLASS,
    "else":   ELSE,
    "false":  FALSE,
    "for":    FOR,
    "fun":    FUN,
    "if":     IF,
    "nil":    NIL,
    "or":     OR,
    "print":  PRINT,
    "return": RETURN,
    "super":  SUPER,
    "this":   THIS,
    "true":   TRUE,
    "var":    VAR,
    "while":  WHILE,
}

type Token struct {
    tokenType TokenType
    lexeme string
    literal any
    line int
}

func (token Token) String() string {
    return fmt.Sprintf("tokenType: %v, lexeme: %s, literal: %v, line: %d\n", token.tokenType, token.lexeme, token.literal, token.line)
}
