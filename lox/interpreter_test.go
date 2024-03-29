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
            name: "division by zero: 11.0 / 0",
            input: Binary {
                left: Literal{float64(11.0)},
                operator: Token{SLASH, "/", "", 1},
                right: Literal{float64(0.0)},
            },
            expected: math.Inf(1),
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

            if resultNumber != test.expected {
                t.Errorf("Incorrect result.\nresult  :%.25f\nexpected:%.25f\n", resultNumber, test.expected)
            }
        })
    }
}

func TestStringConcat(t *testing.T) {
    tests := []struct {
        name string
        input Expr
        expected any
    } {
        {
            name: `"hello" + ", world!"`,
            input: Binary{
                left: Literal{"hello"},
                operator: Token{PLUS, "+", "", 1},
                right: Literal{", world!"},
            },
            expected: "hello, world!",
        },
        {
            name: `"" + ", world!"`,
            input: Binary{
                left: Literal{""},
                operator: Token{PLUS, "+", "", 1},
                right: Literal{", world!"},
            },
            expected: ", world!",
        },
        {
            name: `"hello" + ""`,
            input: Binary{
                left: Literal{"hello"},
                operator: Token{PLUS, "+", "", 1},
                right: Literal{""},
            },
            expected: "hello",
        },
        {
            name: `"" + ""`,
            input: Binary{
                left: Literal{""},
                operator: Token{PLUS, "+", "", 1},
                right: Literal{""},
            },
            expected: "",
        },
        {
            name: `"" + 1`,
            input: Binary{
                left: Literal{""},
                operator: Token{PLUS, "+", "", 1},
                right: Literal{1.0},
            },
            expected: "1",
        },
        {
            name: `1 + ""`,
            input: Binary{
                left: Literal{1.0},
                operator: Token{PLUS, "+", "", 1},
                right: Literal{""},
            },
            expected: "1",
        }, 
        {
            name: `1 + 1 + "1"`,
            input: Binary{
                left: Binary{
                    left: Literal{1.0},
                    operator: Token{PLUS, "+", "", 1},
                    right: Literal{1.0},
                },
                operator: Token{PLUS, "+", "", 1},
                right: Literal{"1"},
            },
            expected: "21",
        }, 
        {
            name: `"1" + 1 + 1`,
            input: Binary{
                left: Binary{
                    left: Literal{"1"},
                    operator: Token{PLUS, "+", "", 1},
                    right: Literal{1.0},
                },
                operator: Token{PLUS, "+", "", 1},
                right: Literal{1.0},
            },
            expected: "111",
        }, 
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result, err := Interpret(test.input)

            if err != nil {
                t.Errorf("no error expected\n")
            }

            resultString, _ := result.(string)

            if resultString != test.expected {
                t.Errorf("Incorrect result.\nresult:  %v\nexpected:%v\n", resultString, test.expected)
            }
        })
    }
}

func TestLogicOperator(t *testing.T) {
    tests := []struct {
        name string
        input Expr
        expected any
    } {
        {
            name: "true == false",
            input: Binary{
                left: Literal{true},
                operator: Token{EQUAL_EQUAL, "==", "", 1},
                right: Literal{false},
            },
            expected: false,
        },
        {
            name: "true == true",
            input: Binary{
                left: Literal{true},
                operator: Token{EQUAL_EQUAL, "==", "", 1},
                right: Literal{true},
            },
            expected: true,
        },
        {
            name: "false == false",
            input: Binary{
                left: Literal{false},
                operator: Token{EQUAL_EQUAL, "==", "", 1},
                right: Literal{false},
            },
            expected: true,
        },
        {
            name: "false != false",
            input: Binary{
                left: Literal{false},
                operator: Token{BANG_EQUAL, "!=", "", 1},
                right: Literal{false},
            },
            expected: false,
        },
        {
            name: "true != false",
            input: Binary{
                left: Literal{true},
                operator: Token{BANG_EQUAL, "!=", "", 1},
                right: Literal{false},
            },
            expected: true,
        },
        {
            name: "!false == true",
            input: Binary{
                left: Unary{
                    operator: Token{BANG, "!", "", 1},
                    right: Literal{false},

                },
                operator: Token{EQUAL_EQUAL, "==", "", 1},
                right: Literal{true},
            },
            expected: true,
        },
        {
            name: "9.5 == 9.5",
            input: Binary{
                left: Literal{9.5},
                operator: Token{EQUAL_EQUAL, "==", "", 1},
                right: Literal{9.5},
            },
            expected: true,
        },
        {
            name: "1 == 2",
            input: Binary{
                left: Literal{1.0},
                operator: Token{EQUAL_EQUAL, "==", "", 1},
                right: Literal{2.0},
            },
            expected: false,
        },
        {
            name: "1 < 2",
            input: Binary{
                left: Literal{1.0},
                operator: Token{LESS, "<", "", 1},
                right: Literal{2.0},
            },
            expected: true,
        },
        {
            name: "1 > 2",
            input: Binary{
                left: Literal{1.0},
                operator: Token{GREATER, ">", "", 1},
                right: Literal{2.0},
            },
            expected: false,
        },
        {
            name: "1 <= 2",
            input: Binary{
                left: Literal{1.0},
                operator: Token{LESS_EQUAL, "<=", "", 1},
                right: Literal{2.0},
            },
            expected: true,
        },
        {
            name: "1 >= 2",
            input: Binary{
                left: Literal{1.0},
                operator: Token{GREATER_EQUAL, ">=", "", 1},
                right: Literal{2.0},
            },
            expected: false,
        },
        {
            name: "2 <= 2",
            input: Binary{
                left: Literal{2.0},
                operator: Token{LESS_EQUAL, "<=", "", 1},
                right: Literal{2.0},
            },
            expected: true,
        },
        {
            name: "2 >= 2",
            input: Binary{
                left: Literal{2.0},
                operator: Token{GREATER_EQUAL, ">=", "", 1},
                right: Literal{2.0},
            },
            expected: true,
        },
        {
            name: `"string" == true`,
            input: Binary{
                left: Literal{"string"},
                operator: Token{EQUAL_EQUAL, "==", "", 1},
                right: Literal{true},
            },
            expected: true,
        },


    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result, err := Interpret(test.input)
            if err != nil {
                t.Errorf("no error expected\n")
            }

            if result != test.expected {
                t.Errorf("Incorrect result.\nresult  :%v\nexpected:%v\n", result, test.expected)
            }
        })
    }
}
