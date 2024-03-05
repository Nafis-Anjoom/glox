package lox

import "testing"

func TestPrint(t *testing.T) {
    expr := Binary{
        left: Unary{
            operator: Token{MINUS, "-", "", 1},
            right: Literal{123},
        },
        operator: Token{STAR, "*", "", 1},
        right: Grouping{
            expression: Literal{45.67},
        },
    }

    result := expr.Print()
    expected := "(* (- 123) (group 45.67))"

    if result != expected {
        t.Errorf("incorrect result.\nresult: %v\nexpected: %v\n", result, expected)
    }
 }
