package lox

import "testing"

func TestPrint(t *testing.T) {
    expr := Binary{
        left: Unary{
            operator: Token{MINUS, "-", "", 1},
            right: Literal{"123", NUMBER},
        },
        operator: Token{STAR, "*", "", 1},
        right: Grouping{
            expression: Literal{"45.67", NUMBER},
        },
    }

    result := expr.Print()
    expected := "(* (- 123) (group 45.67))"

    if result != expected {
        t.Errorf("Result was incorrect, got #{result}, expected: #{expected}")
    }
 }
