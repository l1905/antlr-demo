package main

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"antlr-demo/demo/logic"
	"antlr-demo/demo/parser"
)

func main() {
	runListener()
}

func runListener()  {
	// Setup the input
	is := antlr.NewInputStream("1 +     2 * (3+1)")

	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)


	listen := logic.CalcListener{}
	// Finally parse the expression
	antlr.ParseTreeWalkerDefault.Walk(&listen, p.Start())
	fmt.Println(listen.Pop())
}

func runVisitor() {
	// Setup the input
	is := antlr.NewInputStream("1 +     2 * 3+1+1")
	// Create the Lexer
	lexer := parser.NewCalcLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewCalcParser(stream)

	visitor := logic.CalVisitor{}
	// Finally parse the expression
	p.Start().Accept(&visitor)

	fmt.Println(visitor.Pop())
}