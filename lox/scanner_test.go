package lox

import (
    "testing"
)

type test struct {
    name string
    input string
    expected []Token
}

func TestScan(t *testing.T) {
    tests := []test {
        {
            name: "decimal add decimal",
            input: "1 + 1",
            expected: []Token{
                {tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
                {tokenType: PLUS, lexeme: "+", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "decimal minus decimal",
            input: "10 - 10",
            expected: []Token{
                {tokenType: NUMBER, lexeme: "10", literal: 10.0, line: 1},
                {tokenType: MINUS, lexeme: "-", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "10", literal: 10.0, line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "decimal multiply decimal",
            input: "2 * 20",
            expected: []Token{
                {tokenType: NUMBER, lexeme: "2", literal: 2.0, line: 1},
                {tokenType: STAR, lexeme: "*", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "20", literal: 20.0, line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "decimal divide decimal",
            input: "13 / 10",
            expected: []Token{
                {tokenType: NUMBER, lexeme: "13", literal: 13.0, line: 1},
                {tokenType: SLASH, lexeme: "/", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "10", literal: 10.0, line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "flaot divide decimal",
            input: "13.0001 / 10",
            expected: []Token{
                {tokenType: NUMBER, lexeme: "13.0001", literal: 13.0001, line: 1},
                {tokenType: SLASH, lexeme: "/", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "10", literal: 10.0, line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "grouped addition on decimals divide decimal",
            input: "(1 + 1) / 10",
            expected: []Token{
                {tokenType: LEFT_PAREN, lexeme: "(", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
                {tokenType: PLUS, lexeme: "+", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
                {tokenType: RIGHT_PAREN, lexeme: ")", literal: "", line: 1},
                {tokenType: SLASH, lexeme: "/", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "10", literal: 10.0, line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "is decimal equal decimal",
            input: "1 == 1",
            expected: []Token{
                {tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
                {tokenType: EQUAL_EQUAL, lexeme: "==", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "is decimal not equal decimal",
            input: "1 != 1",
            expected: []Token{
                {tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
                {tokenType: BANG_EQUAL, lexeme: "!=", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "normal string expression",
            input: `"brown fox"`,
            expected: []Token{
                {tokenType: STRING, lexeme:  `"brown fox"`, literal: "brown fox", line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "string expression with comma within",
            input: `"brown, fox"`,
            expected: []Token{
                {tokenType: STRING, lexeme:  `"brown, fox"`, literal: "brown, fox", line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "decimal plus decimal in quotation",
            input: `"1 + 1"`,
            expected: []Token{
                {tokenType: STRING, lexeme:  `"1 + 1"`, literal: "1 + 1", line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "two greater than one",
            input: "2 > 1",
            expected: []Token{
                {tokenType: NUMBER, lexeme:  "2", literal: 2.0, line: 1},
                {tokenType: GREATER, lexeme:  ">", literal: "", line: 1},
                {tokenType: NUMBER, lexeme:  "1", literal: 1.0, line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "two greater than or equal to one",
            input: "2 >= 1",
            expected: []Token{
                {tokenType: NUMBER, lexeme:  "2", literal: 2.0, line: 1},
                {tokenType: GREATER_EQUAL, lexeme:  ">=", literal: "", line: 1},
                {tokenType: NUMBER, lexeme:  "1", literal: 1.0, line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "one less two",
            input: "1 < 2",
            expected: []Token{
                {tokenType: NUMBER, lexeme:  "1", literal: 1.0, line: 1},
                {tokenType: LESS, lexeme:  "<", literal: "", line: 1},
                {tokenType: NUMBER, lexeme:  "2", literal: 2.0, line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "one less than or equal to two",
            input: "1 <= 2",
            expected: []Token{
                {tokenType: NUMBER, lexeme:  "1", literal: 1.0, line: 1},
                {tokenType: LESS_EQUAL, lexeme:  "<=", literal: "", line: 1},
                {tokenType: NUMBER, lexeme:  "2", literal: 2.0, line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "number assigned to a variable called foo",
            input: "foo = 1234",
            expected: []Token{
                {tokenType: IDENTIFIER, lexeme: "foo", literal: "", line: 1},
                {tokenType: EQUAL, lexeme: "=", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "1234", literal: 1234.0, line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "negate equality expression",
            input: "!(1 == 1)",
            expected: []Token{
                {tokenType: BANG, lexeme: "!", literal: "", line: 1},
                {tokenType: LEFT_PAREN, lexeme: "(", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
                {tokenType: EQUAL_EQUAL, lexeme: "==", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
                {tokenType: RIGHT_PAREN, lexeme: ")", literal: "", line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "negate number",
            input: "-1",
            expected: []Token{
                {tokenType: MINUS, lexeme: "-", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "expression end in semicolon",
            input: "1 + 1;",
            expected: []Token{
                {tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
                {tokenType: PLUS, lexeme: "+", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
                {tokenType: SEMICOLON, lexeme: ";", literal: "", line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "comma seperated numbers",
            input: "1, 2, 3",
            expected: []Token{
                {tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
                {tokenType: COMMA, lexeme: ",", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "2", literal: 2.0, line: 1},
                {tokenType: COMMA, lexeme: ",", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "3", literal: 3.0, line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "scoped expression in braces in single line",
            input: "{1 + 1}",
            expected: []Token{
                {tokenType: LEFT_BRACE, lexeme: "{", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
                {tokenType: PLUS, lexeme: "+", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
                {tokenType: RIGHT_BRACE, lexeme: "}", literal: "", line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "access object properties using dot operator",
            input: "foo.bar",
            expected: []Token{
                {tokenType: IDENTIFIER, lexeme: "foo", literal: "", line: 1},
                {tokenType: DOT, lexeme: ".", literal: "", line: 1},
                {tokenType: IDENTIFIER, lexeme: "bar", literal: "", line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
        },
        {
            name: "expressions in new line",
            input: "foo = 1234;\nbar = 4321;",
            expected: []Token{
                {tokenType: IDENTIFIER, lexeme: "foo", literal: "", line: 1},
                {tokenType: EQUAL, lexeme: "=", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "1234", literal: 1234.0, line: 1},
                {tokenType: SEMICOLON, lexeme: ";", literal: "", line: 1},
                {tokenType: IDENTIFIER, lexeme: "bar", literal: "", line: 2},
                {tokenType: EQUAL, lexeme: "=", literal: "", line: 2},
                {tokenType: NUMBER, lexeme: "4321", literal: 4321.0, line: 2},
                {tokenType: SEMICOLON, lexeme: ";", literal: "", line: 2},
                {tokenType: EOF, lexeme: "", literal: "", line: 2},
            },
        },
    }

    for _, testCase := range tests {
        t.Run(testCase.name, func(t *testing.T) {
            result, err := Scan(testCase.input)
            if err != nil {
                t.Errorf("scanning failed during test: '%s'\n", testCase.name)
            }

            isEqual := areTokenArrsEqual(result, testCase.expected)
            if !isEqual {
                t.Errorf("Result was incorrect,\n got result:\n %s,\n expected:\n%s\n", result, testCase.expected)
            }
        })
    }
 }

func areTokenArrsEqual(tokenArr1, tokenArr2 []Token) bool {
    if len(tokenArr1) != len(tokenArr2) {
        return false
    }

    for i := 0; i < len(tokenArr1); i++ {
        if !areTokenEqual(tokenArr1[i], tokenArr2[i]) {
            return false
        }
    }
    return true
}

func areTokenEqual(token1, token2 Token) bool {
    if token1.tokenType == token2.tokenType && token1.lexeme == token2.lexeme &&
        token1.literal == token2.literal && token1.line == token2.line {
        return true
    }
    return false
}
