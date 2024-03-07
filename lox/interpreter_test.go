package lox

import (
    "fmt"
	"testing"
)

func TestInterpreterValid(t *testing.T) {
    tests := []struct {
        name string
        input Expr
        expected any
    } {
        {
            name: "1.1 + 2.2",
            input: Binary {
                right: Literal{float64(1.1)},
                operator: Token{PLUS, "+", "", 1},
                left: Literal{float64(2.2)},
            },
            expected: float64(1.1) + float64(2.2),
        },
    }

    for _, test := range tests {
        t.Run(test.name, func(t *testing.T) {
            result, err := Interpret(test.input)
            if err != nil {
                t.Errorf("no error expected\n")
            }

            resultNumber, _ := result.(float64)
            // fmt.Printf("difference %v\n", (1.1 + 3.3) - 3.3)
            fmt.Println(float64(1.1) + 2.2)

            if result != test.expected {
                t.Errorf("Incorrect result.\nresult:%f\nexpected:%f\n", resultNumber, test.expected)
            }
        })
    }
}
