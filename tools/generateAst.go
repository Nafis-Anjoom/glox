package tools

import (
	"log"
	"os"
	"strings"
)


func DefineAst(outputDir string, baseName string, types []string) {
    path := outputDir + "/" + baseName + ".go"

    file, err := os.Create(path)

    if err != nil {
        log.Fatal("error creating a file")
    }

    defer file.Close()

    var output strings.Builder

    output.WriteString("package main\n\n")

    for i := 0; i < len(types); i++ {
        typeSplits := strings.Split(types[i], ":")
        structName := strings.Trim(typeSplits[0], " ")
        fields := strings.Split(strings.Trim(typeSplits[1], " "), ", ")

        output.WriteString("type " + structName + " struct{\n")
        
        for _, field := range fields {
            temp := strings.Split(field, " ")
            output.WriteString("\t" + temp[1] + " " + temp[0] + "\n")
        }
        output.WriteString("}\n\n")
        
    }

    file.WriteString(output.String())
}
