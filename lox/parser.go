package lox

import (
	"errors"
	"log"
)

func Parse(tokens []Token) (Expr, error) {
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

	consume := func(tokenType TokenType, message string) (Token, error) {
		if check(tokenType) {
			return advance(), nil
		}

		// return peek(), errors.New(message)
		// should return peek as part of error
		return peek(), errors.New(message)
	}

	// TODO: fix recursion
	unary := func() Expr {
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

	primary := func() (Expr, error) {
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
		return nil, errors.New("Expect expression")
	}
	return nil, nil
}
