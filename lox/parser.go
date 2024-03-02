package lox

import (
	"fmt"
	"log"
)

type ParserError struct {
    token Token
}

func (parserError ParserError) Error() string {
    return parserError.token.String()
}

func Parse(tokens []Token) Expr {
	current := 0

	peek := func() Token {
		return tokens[current]
	}

	isAtEnd := func() bool {
		return peek().tokenType == EOF
	}

	previous := func() Token {
		return tokens[current-1]
	}

	advance := func() Token {
		if !isAtEnd() {
			current++
		}
		return previous()
	}

	check := func(tokenType TokenType) bool {
		if isAtEnd() {
			return false
		}
		return peek().tokenType == tokenType
	}

	match := func(types ...TokenType) bool {
		for i := 0; i < len(types); i++ {
			if check(types[i]) {
				advance()
				return true
			}
		}
		return false
	}

    reportError := func(token Token, message string) {
        if token.tokenType == EOF {
            fmt.Printf("[line %d] Error %s: %s\n", token.line, "at end", message)
        } else {
            fmt.Printf("[line %d] Error %s: %s\n", token.line, fmt.Sprintf("at '%s'", token.lexeme), message)
        }
    }

	consume := func(tokenType TokenType, message string) (Token, error) {
		if check(tokenType) {
			return advance(), nil
		}

        // TODO: refactor error reporting to its own package/function
        reportError(peek(), message)
        return peek(), ParserError{}
	}

	// TODO: fix recursion
    var unary func() Expr
	var primary func() (Expr, error)

	unary = func() Expr {
		if match(BANG, MINUS) {
			operator := previous()
			right := unary()
			return Unary{operator, right}
		}

		expr, err := primary()
		if err != nil {
			log.Fatal("error occurred in unary")
		}
		return expr
	}

	factor := func() Expr {
		expr := unary()

		for match(SLASH, STAR) {
			operator := previous()
			right := unary()
			expr = Binary{expr, operator, right}
		}

		return expr
	}

	term := func() Expr {
		expr := factor()

		for match(MINUS, PLUS) {
			operator := previous()
			right := factor()
			expr = Binary{expr, operator, right}
		}

		return expr
	}

	comparison := func() Expr {
		expr := term()

		for match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
			operator := previous()
			right := term()
			expr = Binary{expr, operator, right}
		}
		return expr
	}

	equality := func() Expr {
		expr := comparison()

		for match(BANG_EQUAL, EQUAL_EQUAL) {
			operator := previous()
			right := comparison()
			expr = Binary{expr, operator, right}
		}

		return expr
	}

	expression := func() Expr {
		return equality()
	}

	primary = func() (Expr, error) {
		if match(FALSE) {
			return Literal{"false"}, nil
		}
		if match(TRUE) {
			return Literal{"true"}, nil
		}
		if match(NIL) {
			return Literal{""}, nil
		}

		if match(NUMBER, STRING) {
			return Literal{previous().literal}, nil
		}

		if match(LEFT_PAREN) {
			expr := expression()
			consume(RIGHT_PAREN, "Expect ')' after expression.")
			return Grouping{expr}, nil
		}

		// return token with error
        reportError(peek(), "Expect expression");
		return nil, ParserError{}
	}

    expr := expression()
    return expr
}
