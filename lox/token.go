package lox

import "fmt"

type TokenType int64

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
    literal string
    line int
}

func (token Token) String() string {
    return fmt.Sprintf("%v %v %v %d", token.tokenType, token.lexeme, token.literal, token.line)
}
