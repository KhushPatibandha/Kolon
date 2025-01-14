package main

import (
	"fmt"
	"os"

	"github.com/KhushPatibandha/Kolon/src/interpreter/evaluator"
	"github.com/KhushPatibandha/Kolon/src/interpreter/object"
	"github.com/KhushPatibandha/Kolon/src/lexer"
	"github.com/KhushPatibandha/Kolon/src/parser"
)

func main() {
	if len(os.Args) < 3 || os.Args[1] != "run:" {
		if os.Args[1] == "--version" {
			fmt.Println("Kolon v0.0.3")
			return
		} else {
			fmt.Println("Usage: kolon run: <path-to-kolon-file>")
			return
		}
	}
	filePath := os.Args[2]
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	tokens := lexer.Tokenizer(string(bytes))
	// for _, token := range tokens {
	// 	token.Help()
	// }

	parser := parser.New(tokens)
	program, err := parser.ParseProgram()
	if err != nil {
		fmt.Println("Error parsing program:", err)
		return
	}
	env := object.NewEnvironment()
	_, _, err = evaluator.Eval(program, env)
	if err != nil {
		fmt.Println("Error evaluating program:", err)
		return
	}
}
