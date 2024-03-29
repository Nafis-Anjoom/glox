package lox

import (
	"errors"
	"fmt"
)

func Interpret(expr Expr) (any, error) {
    value, err := evaluate(expr)
    return value, err
}

func evaluate(expr Expr) (any, error) {
    switch expr.(type) {
        case Literal:
            literal, _ := expr.(Literal)
            return literal.value, nil
        case Grouping:
            grouping, _ := expr.(Grouping)
            return evaluate(grouping.expression)
        case Unary:
            unary, _ := expr.(Unary)
            return evaluateUnary(unary)
        case Binary:
            binary, _ := expr.(Binary)
            return evaluateBinary(binary)
    }
    return nil, errors.New("error occurred while evaluating")
}

// If both lhs and rhs are numbers, then all operations are valid
// If sum of two numbers exceed 1.7976931348623157e+308 or recedes -1.7976931348623157e+308, return +inf or -inf
// If both lhs and rhs are strings, then only plus operation is valid
// otherwise, "return not a number"
// If division by zero, return inf (follow ecmaScript)
// TODO: require heavy testing
func evaluateBinary(binary Binary) (any, error) {
    left, err := evaluate(binary.left)
    if err != nil {
        return nil, errors.New("error evaluating the lhs of binary expression")
    }
    right, err := evaluate(binary.right)
    if err != nil {
        return nil, errors.New("error evaluating the rhs of binary expression")
    }

    leftNumber, isLhsFloat := left.(float64)
    rightNumber, isRhsFloat := right.(float64)

    if isLhsFloat && isRhsFloat {
        switch binary.operator.tokenType {
            case MINUS:
                return leftNumber - rightNumber, nil
            case SLASH:
                return leftNumber / rightNumber, nil
            case STAR:
                return leftNumber * rightNumber, nil
            case PLUS:
                return leftNumber + rightNumber, nil
            case GREATER:
                return leftNumber > rightNumber, nil
            case GREATER_EQUAL:
                return leftNumber >= rightNumber, nil
            case LESS:
                return leftNumber < rightNumber, nil
            case LESS_EQUAL:
                return leftNumber <= rightNumber, nil
            case BANG_EQUAL:
                return leftNumber != rightNumber, nil
            case EQUAL_EQUAL:
                return leftNumber == rightNumber, nil
        }
    } else {
        switch binary.operator.tokenType {
            case BANG_EQUAL:
                return isTruthy(left) != isTruthy(right), nil
            case EQUAL_EQUAL:
                return isTruthy(left) == isTruthy(right), nil
            case PLUS:
                return fmt.Sprintf("%v%v", left, right), nil
        }
    }
    fmt.Println("invalid")
    return nil, errors.New("not a number")
}

func evaluateUnary(unary Unary) (any, error) {
    right, err := evaluate(unary.right)

    if err != nil {
        return nil, errors.New("error evaluating unary expression")
    }

    if unary.operator.tokenType == MINUS {
        number, ok := right.(float64)

        if !ok {
            return nil, errors.New("not a number")
        }
        return -number, nil
    } else if unary.operator.tokenType == BANG {
        return !isTruthy(right), nil
    }
    return nil, errors.New("error evaluating unary expression")
}

func isTruthy(expr any) bool {
    if expr == nil {
        return false
    }
    val, ok := expr.(bool)
    if !ok {
        return true
    }
    return val
}
