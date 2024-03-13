package lox

import "testing"

func TestParseValid(t *testing.T) {
	tests := []struct {
		name     string
		input    []Token
		expected Expr
	}{
		{
			name: "one plus one",
			input: []Token{
				{tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
				{tokenType: PLUS, lexeme: "+", literal: "", line: 1},
				{tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
				{tokenType: EOF, lexeme: "", literal: "", line: 1},
			},
			expected: Binary{
				left:     Literal{value: 1.0},
				operator: Token{tokenType: PLUS, lexeme: "+", literal: "", line: 1},
				right:     Literal{value: 1.0},
			},
		},
		{
            name: "left-associative operations: one plus one minus two",
			input: []Token{
				{tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
				{tokenType: PLUS, lexeme: "+", literal: "", line: 1},
				{tokenType: NUMBER, lexeme: "1.1234", literal: 1.1234, line: 1},
				{tokenType: MINUS, lexeme: "-", literal: "", line: 1},
				{tokenType: NUMBER, lexeme: "2", literal: 2.0, line: 1},
				{tokenType: EOF, lexeme: "", literal: "", line: 1},
			},
			expected: Binary{
				left:     Binary{
                    left: Literal{1.0},
                    operator: Token{tokenType: PLUS, lexeme: "+", literal: "", line: 1},
                    right: Literal{1.1234},
                },
				operator: Token{tokenType: MINUS, lexeme: "-", literal: "", line: 1},
				right:     Literal{value: 2.0},
			},
		},
		{
            name: "complex arithmatic: (1.1 + 2 - 10) * 1.10000001 / 2.24354352",
			input: []Token{
				{tokenType: LEFT_PAREN, lexeme: "(", literal: "", line: 1},
				{tokenType: NUMBER, lexeme: "1.1", literal: 1.1, line: 1},
				{tokenType: PLUS, lexeme: "+", literal: "", line: 1},
				{tokenType: NUMBER, lexeme: "2", literal: 2, line: 1},
				{tokenType: MINUS, lexeme: "-", literal: "", line: 1},
				{tokenType: NUMBER, lexeme: "10", literal: 10, line: 1},
				{tokenType: RIGHT_PAREN, lexeme: ")", literal: "", line: 1},
				{tokenType: STAR, lexeme: "*", literal: "", line: 1},
				{tokenType: NUMBER, lexeme: "1.10000001", literal: 1.10000001, line: 1},
				{tokenType: SLASH, lexeme: "/", literal: "", line: 1},
				{tokenType: NUMBER, lexeme: "2.24354352", literal: 2.24354352, line: 1},
				{tokenType: EOF, lexeme: "", literal: "", line: 1},
			},
			expected: Binary{
                left: Binary{
                    left: Grouping{
                        expression: Binary{
                            left: Binary{
                                left: Literal{1.1},
                                operator: Token{PLUS, "+", "", 1},
                                right: Literal{2},
                            },
                            operator: Token{MINUS, "-", "", 1},
                            right: Literal{10},
                        },
                    },
                    operator: Token{STAR, "*", "", 1},
                    right: Literal{1.10000001},
                },
                operator: Token{SLASH, "/", "", 1},
                right: Literal{2.24354352} ,
            },
		},
		{
			name: "negated decimal multiply grouped expression: -123 * (1 + 1)",
			input: []Token{
				{tokenType: MINUS, lexeme: "-", literal: "", line: 1},
				{tokenType: NUMBER, lexeme: "123", literal: 123.0, line: 1},
				{tokenType: STAR, lexeme: "*", literal: "", line: 1},
				{tokenType: LEFT_PAREN, lexeme: "(", literal: "", line: 1},
				{tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
				{tokenType: PLUS, lexeme: "+", literal: "", line: 1},
				{tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
				{tokenType: RIGHT_PAREN, lexeme: ")", literal: "", line: 1},
				{tokenType: EOF, lexeme: "", literal: "", line: 1},
			},
			expected: Binary{
				left: Unary{
					operator: Token{MINUS, "-", "", 1},
					right:    Literal{123.0},
				},
				operator: Token{STAR, "*", "", 1},
				right: Grouping{
					expression: Binary{
						left:     Literal{value: 1.0},
						operator: Token{tokenType: PLUS, lexeme: "+", literal: "", line: 1},
						right:    Literal{value: 1.0},
					},
				},
			},
		},
		{
			name: "comparision between two groupings: (1 + 1) == (1 + 1)",
			input: []Token{
				{tokenType: LEFT_PAREN, lexeme: "(", literal: "", line: 1},
				{tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
				{tokenType: PLUS, lexeme: "+", literal: "", line: 1},
				{tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
				{tokenType: RIGHT_PAREN, lexeme: ")", literal: "", line: 1},
				{tokenType: EQUAL_EQUAL, lexeme: "==", literal: "", line: 1},
				{tokenType: LEFT_PAREN, lexeme: "(", literal: "", line: 1},
				{tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
				{tokenType: PLUS, lexeme: "+", literal: "", line: 1},
				{tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
				{tokenType: RIGHT_PAREN, lexeme: ")", literal: "", line: 1},
				{tokenType: EOF, lexeme: "", literal: "", line: 1},
			},
			expected: Binary{
				left: Grouping{
					expression: Binary{
						left:     Literal{value: 1.0},
						operator: Token{tokenType: PLUS, lexeme: "+", literal: "", line: 1},
						right:    Literal{value: 1.0},
					},
				},
				operator: Token{tokenType: EQUAL_EQUAL, lexeme: "==", literal: "", line: 1},
				right: Grouping{
					expression: Binary{
						left:     Literal{value: 1.0},
						operator: Token{tokenType: PLUS, lexeme: "+", literal: "", line: 1},
						right:    Literal{value: 1.0},
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			expression, err := Parse(test.input)

			if err != nil {
				t.Error("error occurred\n")
			}

			if expression != test.expected {
				t.Errorf("result was incorrect.\nresult  :%+v\nexpected:%+v\n", expression, test.expected)
			}
		})
	}
}

func TestParseInvalid(t *testing.T) {
	tests := []struct {
		name     string
		input    []Token
		expected Expr
	}{
        {
            name: "invalid expression: 1 + + 1",
            input: []Token{
                {tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
                {tokenType: PLUS, lexeme: "+", literal: "", line: 1},
                {tokenType: PLUS, lexeme: "+", literal: "", line: 1},
                {tokenType: NUMBER, lexeme: "1", literal: 1.0, line: 1},
                {tokenType: EOF, lexeme: "", literal: "", line: 1},
            },
            expected: nil,
        },
        // {
        //     name: "invalid grouping with extra right paren: (1 + 1))",
        //     input: []Token{
        //         {tokenType: LEFT_PAREN, lexeme: "(", literal: "", line: 1},
        //         {tokenType: NUMBER, lexeme: "1", literal: "1", line: 1},
        //         {tokenType: PLUS, lexeme: "+", literal: "", line: 1},
        //         {tokenType: NUMBER, lexeme: "1", literal: "1", line: 1},
        //         {tokenType: RIGHT_PAREN, lexeme: ")", literal: "", line: 1},
        //         {tokenType: RIGHT_PAREN, lexeme: ")", literal: "", line: 1},
        //         {tokenType: EOF, lexeme: "", literal: "", line: 1},
        //     },
        //     expected: nil,
        // },
        // {
        //     name: "invalid grouping with no left paren: 1 + 1)",
        //     input: []Token{
        //         {tokenType: NUMBER, lexeme: "1", literal: "1", line: 1},
        //         {tokenType: PLUS, lexeme: "+", literal: "", line: 1},
        //         {tokenType: NUMBER, lexeme: "1", literal: "1", line: 1},
        //         {tokenType: RIGHT_PAREN, lexeme: ")", literal: "", line: 1},
        //         {tokenType: EOF, lexeme: "", literal: "", line: 1},
        //     },
        //     expected: nil,
        // },
        // {
        //     name: "invalid grouping with no right paren: 1 + 1)",
        //     input: []Token{
        //         {tokenType: LEFT_PAREN, lexeme: "(", literal: "", line: 1},
        //         {tokenType: NUMBER, lexeme: "1", literal: "1", line: 1},
        //         {tokenType: PLUS, lexeme: "+", literal: "", line: 1},
        //         {tokenType: NUMBER, lexeme: "1", literal: "1", line: 1},
        //         {tokenType: EOF, lexeme: "", literal: "", line: 1},
        //     },
        //     expected: nil,
        // },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            expresison, err := Parse(test.input)

            if err == nil && expresison != nil {
                t.Errorf("result was incorrect. result: %+v\n", expresison)
            }
        })
    }
    


}
