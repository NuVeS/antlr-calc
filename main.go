// example2.go
package main

import (
	"github.com/NuVeS/AntlrCalc/my_listener"
	"github.com/NuVeS/AntlrCalc/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func calc(input string) int {
	// Setup the input
	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewCalculatorLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalculatorParser(stream)

	// Finally parse the expression (by walking the tree)
	var listener my_listener.CalcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	return listener.Pop()
}

func main() {
	println(calc("1 + 4 * 2^(4 - 1)!! / ((3) ^ 3!!)"))
}
