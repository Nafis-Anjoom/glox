package lox

import "log"

func Scan(text string) ([]Token, error) {
    runes := []rune(text)

    start := 0
    current := 0
    line := 1

    tokens := make([]Token, 0)

    advance := func() rune {
        current++
        return runes[current - 1]
    }

    addToken := func(tokenType TokenType) {
        tokens = append(tokens, Token{ tokenType, string(runes[start:current]), "", line })
    }

    isAtEnd := func() bool {
        return current >= len(runes)
    }

    match := func(expected rune) bool {
        if isAtEnd() || runes[current] != expected {
            return false
        }

        current++
        
        return true
    }

    peek := func() rune {
        if isAtEnd() {
            return '\x00'
        }
        return runes[current]
    }

    tokenizeString := func() {
        for peek() != '"' && !isAtEnd() {
            if peek() == '\n' {
                line++
            }
            advance()
        }

        if isAtEnd() {
            log.Fatal("Unterminated string")
            return
        }

        advance()

        lexeme := string(runes[start: current])
        value := string(runes[start + 1: current - 1])
        tokens = append(tokens, Token{ STRING, lexeme, value, line })
    }

    isDigit := func(c rune) bool {
        return c >= '0' && c <= '9'
    }

    peekNext := func() rune {
        if current + 1 >= len(runes) {
            return '\x00'
        }
        return runes[current + 1]
    }

    number := func() {
        for isDigit(peek()) {
            advance()
        }

        if peek() == '.' && isDigit(peekNext()) {
            advance()
            for isDigit(peek()) {
                advance()
            }
        }

        number := string(runes[start: current])
        tokens = append(tokens, Token{ NUMBER, number, number, line })

    }

    isAlpha := func(c rune) bool {
        return (c >= 'a' && c <= 'z') ||
               (c >= 'A' && c <= 'Z') ||
               (c == '_');
    }

    isAlphaNumeric := func(c rune) bool {
        return isAlpha(c) || isDigit(c)
    }

    identifier := func() {
        for isAlphaNumeric(peek()) {
            advance()
        }

        text := string(runes[start: current])
        tokenType, isFound := keywords[text]

        if !isFound {
            tokenType = IDENTIFIER
        }
        tokens = append(tokens, Token{ tokenType, text, "", line })
    }

    scanToken := func() error {
        c := advance()

        switch c {
        case '(':
            {
                addToken(LEFT_PAREN)
                break
            }
        case ')':
            {
                addToken(RIGHT_PAREN)
                break
            }
        case '{':
            {
                addToken(LEFT_BRACE)
                break
            }
        case '}':
            {
                addToken(RIGHT_BRACE)
                break
            }
        case ',':
            {
                addToken(COMMA)
                break
            }
        case '.':
            {
                addToken(DOT)
                break
            }
        case '-':
            {
                addToken(MINUS)
                break
            }
        case '+':
            {
                addToken(PLUS)
                break
            }
        case ';':
            {
                addToken(SEMICOLON)
                break
            }
        case '*':
            {
                addToken(STAR)
                break
            }
        case '!':
            {
                if match('=') {
                    addToken(BANG_EQUAL)
                } else {
                    addToken(BANG)
                }
                break
            }
        case '=':
            {
                if match('=') {
                    addToken(EQUAL_EQUAL)
                } else {
                    addToken(EQUAL)
                }
                break
            }
        case '<':
            {
                if match('=') {
                    addToken(LESS_EQUAL)
                } else {
                    addToken(LESS)
                }
                break
            }
        case '>':
            {
                if match('=') {
                    addToken(GREATER_EQUAL)
                } else {
                    addToken(GREATER)
                }
                break
            }
        case '/':
            {
                if match('/') {
                    for ;peek() != '\n' && !isAtEnd(); {
                        advance()
                    }
                } else {
                    addToken(SLASH)
                }
                break
            }
        case ' ':
        case '\r':
        case '\t':
            {
                break
            }
        case '\n':
            {
                line++
                break
            }
        case '"':
            {
                tokenizeString()
                break
            }
        default:
            {
                if isDigit(c) {
                    number()
                } else if isAlpha(c) {
                    identifier()
                } else {
                    log.Fatal("Unexpected Character")
                }
                break
            }

    }

        return nil;
    }

    for !isAtEnd() {
        start = current
        if err := scanToken(); err != nil {
            return nil, err
        }
    }

    tokens = append(tokens, Token{ EOF, "", "", line })
    
    return tokens, nil
}

