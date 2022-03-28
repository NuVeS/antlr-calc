// example2.go
package main

import (
	"github.com/NuVeS/AntlrCalc/my_listener"
	"github.com/NuVeS/AntlrCalc/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

func calc(input string) int {
	inputStream := antlr.NewInputStream(input)

	lexer := parser.NewCalculatorLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	p := parser.NewCalculatorParser(stream)

	var listener my_listener.CalcListener
	antlr.ParseTreeWalkerDefault.Walk(&listener, p.Start())

	return listener.Pop()
}

func main() {
	println(calc("1 + 4 * 2^(4 - 1)!! / ((3) ^ 3!!)"))
}
