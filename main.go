package main

import (
	"fmt"
	"log"
	"os"
    "glox/lox"
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


    // fmt.Println(expr.Print())
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
    tokens, err := lox.Scan(source)
    if err != nil {
        log.Fatal("Error while scanning")
    }

    for _, token := range tokens {
        fmt.Println(token)
    }
}

