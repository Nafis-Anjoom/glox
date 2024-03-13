package lox

import (
	"math"
	"testing"
)

func TestInterpreterArithmetic(t *testing.T) {
    tests := []struct {
        name string
        input Expr
        expected any
    } {
        {
            name: "low-precision flaoting point addition: 1.1 + 2.2",
            input: Binary {
                right: Literal{float64(1.1)},
                operator: Token{PLUS, "+", "", 1},
                left: Literal{float64(2.2)},
            },
            expected: float64(1.1) + float64(2.2),
        },
        {
            name: "high-precision floating point addition: 1.10000001 + 2.24354352",
            input: Binary {
                right: Literal{float64(1.10000001)},
                operator: Token{PLUS, "+", "", 1},
                left: Literal{float64(2.24354352)},
            },
            expected: float64(1.10000001) + float64(2.24354352),
        },
        {
            name: "low-precision flaoting point positive difference: 2.2 - 1.1",
            input: Binary {
                left: Literal{float64(2.2)},
                operator: Token{MINUS, "-", "", 1},
                right: Literal{float64(1.1)},
            },
            expected: float64(2.2) - float64(1.1),
        },
        {
            name: "high-precision floating point positive difference: 2.24354352 - 1.10000001",
            input: Binary {
                left: Literal{float64(2.24354352)},
                operator: Token{MINUS, "-", "", 1},
                right: Literal{float64(1.10000001)},
            },
            expected: float64(2.24354352) - float64(1.10000001),
        },
        {
            name: "low-precision flaoting point negative difference: 1.1 - 2.2",
            input: Binary {
                right: Literal{float64(2.2)},
                operator: Token{MINUS, "-", "", 1},
                left: Literal{float64(1.1)},
            },
            expected: float64(1.1) - float64(2.2),
        },
        {
            name: "high-precision floating point negative difference: 1.10000001 - 2.24354352",
            input: Binary {
                right: Literal{float64(2.24354352)},
                operator: Token{MINUS, "-", "", 1},
                left: Literal{float64(1.10000001)},
            },
            expected: float64(1.10000001) - float64(2.24354352),
        },
        {
            name: "low-precision flaoting point multiplication: 3.3 * 2.2",
            input: Binary {
                left: Literal{float64(3.3)},
                operator: Token{STAR, "*", "", 1},
                right: Literal{float64(2.2)},
            },
            expected: float64(3.3) * float64(2.2),
        },
        {
            name: "high-precision floating point multiplication: 1.10000001 * 2.24354352",
            input: Binary {
                right: Literal{float64(1.10000001)},
                operator: Token{STAR, "*", "", 1},
                left: Literal{float64(2.24354352)},
            },
            expected: float64(1.10000001) * float64(2.24354352),
        },
        {
            name: "multiplication overflows to +Inf: 1.7976931348623157e+308 * 1.5",
            input: Binary {
                right: Literal{float64(1.7976931348623157e+308)},
                operator: Token{STAR, "*", "", 1},
                left: Literal{float64(1.5)},
            },
            expected: math.Inf(1),
        },
        {
            name: "multiplication underflows to -Inf: 1.7976931348623157e+308 * -2",
            input: Binary {
                left: Literal{float64(1.7976931348623157e+308)},
                operator: Token{STAR, "*", "", 1},
                right: Unary{
                    operator: Token{MINUS, "-", "", 1},
                    right: Literal{float64(2)},
                },
            },
            expected: math.Inf(-1),
        },
        {
            name: "low-precision flaoting point division: 3.3 / 2.2",
            input: Binary {
                left: Literal{float64(3.3)},
                operator: Token{SLASH, "/", "", 1},
                right: Literal{float64(2.2)},
            },
            expected: float64(3.3) / float64(2.2),
        },
        {
            name: "high-precision floating point division: 1.10000001 / 2.24354352",
            input: Binary {
                left: Literal{float64(1.10000001)},
                operator: Token{SLASH, "/", "", 1},
                right: Literal{float64(2.24354352)},
            },
            expected: float64(1.10000001) / float64(2.24354352),
        },
        {
            name: "grouped expression: (1.1 + 2 - 10)",
            input: Grouping{
                expression: Binary{
                    left: Binary{
                        left: Literal{1.1},
                        operator: Token{PLUS, "+", "", 1},
                        right: Literal{2.0},
                    },
                    operator: Token{MINUS, "-", "", 1},
                    right: Literal{10.0},
                },
            },
            expected: (1.1 + 2 - 10),
        },
        {
            name: "grouped expression binary: (1.1 + 2)",
            input: Grouping{
                expression: Binary{
                    left: Literal{1.1},
                    operator: Token{PLUS, "+", "", 1},
                    right: Literal{2.0},
                },
            },
            expected: (1.1 + 2.0),
        },
        {
            name: "low-precision floating point grouped division: (1.1 + 2 - 10) * 1.11 / 2.242",
            input: Binary{
                left: Binary{
                    left: Grouping{
                        expression: Binary{
                            left: Binary{
                                left: Literal{1.1},
                                operator: Token{PLUS, "+", "", 1},
                                right: Literal{2.0},
                            },
                            operator: Token{MINUS, "-", "", 1},
                            right: Literal{10.0},
                        },
                    },
                    operator: Token{STAR, "*", "", 1},
                    right: Literal{1.1},
                },
                operator: Token{SLASH, "/", "", 1},
                right: Literal{2.242} ,
            },
            expected: (1.1 + 2.0 - 10.0) * 1.1 / 2.242,
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result, err := Interpret(test.input)
            if err != nil {
                t.Errorf("no error expected\n")
            }

            resultNumber, _ := result.(float64)

            if result != test.expected {
                t.Errorf("Incorrect result.\nresult  :%.25f\nexpected:%.25f\n", resultNumber, test.expected)
            }
        })
    }
}
