package lox

import "strings"

type Expr interface {
    Print() string
}

func parenthesize(name string, exprs ...Expr) string {
    var builder strings.Builder

    builder.WriteString("(")
    builder.WriteString(name)
    for _, expr := range exprs {
        builder.WriteString(" ")
        builder.WriteString(expr.Print())
    }
    builder.WriteString(")")

    return builder.String()
}

type Binary struct{
	left Expr
	operator Token
	right Expr
}

func (b Binary) Print() string {
    return parenthesize(b.operator.lexeme, b.left, b.right)
}


type Grouping struct{
	expression Expr
}

func (g Grouping) Print() string {
    return parenthesize("group", g.expression)
}

type Literal struct{
	value string
}

func (l Literal) Print() string {
    if l.value == "" {
        return "nil"
    }
    return l.value
}

type Unary struct{
	operator Token
	right Expr
}

func (u Unary) Print() string {
    return parenthesize(u.operator.lexeme, u.right)
}

