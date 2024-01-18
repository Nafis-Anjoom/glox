package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
    // args := os.Args
    // runFile(args[1])
    // types := []string{
    //     "Binary   : Expr left, Token operator, Expr right",
    //     "Grouping : Expr expression",
    //     "Literal  : Object value",
    //     "Unary    : Token operator, Expr right",
    // }
    // tools.DefineAst("./", "expr", types)

    expr := Binary{
        left: &Unary{
            operator: Token{MINUS, "-", "", 1},
            right: &Literal{"123"},
        },
        operator: Token{STAR, "*", "", 1},
        right: &Grouping{
            expression: &Literal{"45.67"},
        },
    }

    fmt.Println(expr.Print())
}

func runFile(path string) {
    fmt.Println("running file")
    bytes, err := os.ReadFile(path)
    if err != nil {
        log.Fatal(err)
    }

    run(string(bytes))
}

func run(source string) {
    scanner := Scanner { source }

    tokens, err := scanner.Scan()
    if err != nil {
        log.Fatal("Error while scanning")
    }

    for _, token := range tokens {
        fmt.Println(token)
    }

}

