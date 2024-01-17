package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
    args := os.Args
    runFile(args[1])
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

