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

    reportError := func(token Token, message string) {
        if token.tokenType == EOF {
            log.Printf("[line %d] Error %s: %s\n", token.line, "at end", message)
        } else {
            log.Printf("[line %d] Error %s: %s\n", token.line, fmt.Sprintf("at '%s'", token.lexeme), message)
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
    var unary func() (Expr, error)
	var primary func() (Expr, error)

	unary = func() (Expr, error) {
		if match(BANG, MINUS) {
			operator := previous()
			right, err := unary()

            if err != nil {
                return nil, ParserError{}
            }

			return Unary{operator, right}, nil
		}

		expr, err := primary()
        if err != nil {
            return nil, ParserError{}
        }

		return expr, nil
	}

	factor := func() (Expr, error) {
		expr, err := unary()

        if err != nil {
            return nil, ParserError{}
        }

		for match(SLASH, STAR) {
			operator := previous()
			right, err := unary()

            if err != nil {
                return nil, ParserError{}
            }

			expr = Binary{expr, operator, right}
		}

		return expr, nil
	}

	term := func() (Expr, error) {
		expr, err := factor()

        if err != nil {
            return nil, ParserError{}
        }

		for match(MINUS, PLUS) {
			operator := previous()
			right, err := factor()

            if err != nil {
                return nil, ParserError{}
            }

			expr = Binary{expr, operator, right}
		}

		return expr, nil
	}

	comparison := func() (Expr, error) {
		expr, err := term()

        if err != nil {
            return nil, ParserError{}
        }

		for match(GREATER, GREATER_EQUAL, LESS, LESS_EQUAL) {
			operator := previous()
			right, err := term()

            if err != nil {
                return nil, ParserError{}
            }

			expr = Binary{expr, operator, right}
		}
		return expr, nil
	}

	equality := func() (Expr, error) {
		expr, err := comparison()

        if err != nil {
            return nil, ParserError{}
        }

		for match(BANG_EQUAL, EQUAL_EQUAL) {
			operator := previous()
			right, err := comparison()

            if err != nil {
                return nil, ParserError{}
            }

			expr = Binary{expr, operator, right}
		}

		return expr, nil
	}

	expression := func() (Expr, error) {
		return equality()
	}

	primary = func() (Expr, error) {
		if match(FALSE) {
			return Literal{"false", FALSE}, nil
		}
		if match(TRUE) {
			return Literal{"true", TRUE}, nil
		}
		if match(NIL) {
			return Literal{"", NIL}, nil
		}

		if match(NUMBER) {
			return Literal{previous().literal, NUMBER}, nil
		}

		if match(STRING) {
			return Literal{previous().literal, STRING}, nil
		}

		if match(LEFT_PAREN) {
			expr, err := expression()

            if err != nil {
                return nil, ParserError{}
            }

			consume(RIGHT_PAREN, "Expect ')' after expression")
			return Grouping{expr}, nil
		}

		// return token with error
        reportError(peek(), "Expect expression");
		return nil, ParserError{}
	}

    expr, err := expression()
    
    if err != nil {
        return nil, ParserError{}
    }

    return expr, nil
}
