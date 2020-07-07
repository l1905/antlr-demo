package logic

import (
	"fmt"
	"antlr-demo/demo/parser"
	"strconv"
)

type CalcListener struct {
	*parser.BaseCalcListener

	stack []int
}

func (l *CalcListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *CalcListener) Pop() int {
	if len(l.stack) < 1 {
		panic("Stack is empty unable to Pop")
	}

	// Get the last value from the Stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the Stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *CalcListener) ExitMulDiv(c *parser.MulDivContext) {
	fmt.Println("ExitMulDiv")
	right, left := l.Pop(), l.Pop()
	switch c.GetOp().GetTokenType() {
	case parser.CalcParserMUL:
		l.push(left * right)
	case parser.CalcParserDIV:
		l.push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *CalcListener) ExitAddSub(c *parser.AddSubContext) {
	fmt.Println("ExitAddSub=====")
	right, left := l.Pop(), l.Pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalcParserADD:
		l.push(left + right)
	case parser.CalcParserSUB:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *CalcListener) ExitNumber(c *parser.NumberContext) {
	fmt.Println("ExitNumber")
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}

	l.push(i)
}

func (l *CalcListener) ExitParenthesis(ctx *parser.ParenthesisContext) {

}

// 左子树 （2,3)+
// 中间子树 (+ (2,3)
// 1 + 2 * 3