package lox

import "testing"

func TestString(t *testing.T) {
    tests := []struct {
        name string
        input Token
        expected string
    } {
        {
            name: "identifier literal",
            input: Token {tokenType: IDENTIFIER, lexeme: "foo", literal: "foo", line: 1},
            expected: "tokenType: IDENTIFIER, lexeme: foo, literal: foo, line: 1\n",
        },
        {
            name: "string literal",
            input: Token {tokenType: STRING, lexeme: "\"temp string\"", literal: "temp string", line: 1},
            expected: "tokenType: STRING, lexeme: \"temp string\", literal: temp string, line: 1\n",
        },
        {
            name: "number literal",
            input: Token {tokenType: NUMBER, lexeme: "1234.4321", literal: "1234.4321", line: 1},
            expected: "tokenType: NUMBER, lexeme: 1234.4321, literal: 1234.4321, line: 1\n",
        },
        {
            name: "non-literal",
            input: Token {tokenType: COMMA, lexeme: ",", literal: "", line: 1},
            expected: "tokenType: COMMA, lexeme: ,, literal: , line: 1\n",
        },
    }

    for _, testCase := range tests {
        t.Run(testCase.name, func(t *testing.T) {
            result := testCase.input.String()
            if result != testCase.expected {
                t.Errorf("Incorrect Result.\nresult : %s\nexpcted: %s\n", result, testCase.expected)
            }
        })
    }
 }
