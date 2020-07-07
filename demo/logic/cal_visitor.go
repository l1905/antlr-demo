package logic

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"antlr-demo/demo/parser"
	"strconv"
)

type CalVisitor struct {
	parser.BaseCalcVisitor
	Stack []int
}

func (l *CalVisitor) push(i int) {
	l.Stack = append(l.Stack, i)
}

func (l *CalVisitor) Pop() int {
	if len(l.Stack) < 1 {
		panic("Stack is empty unable to Pop")
	}
	// Get the last value from the Stack.
	result := l.Stack[len(l.Stack)-1]

	// Remove the last element from the Stack.
	l.Stack = l.Stack[:len(l.Stack)-1]

	return result
}

func (v *CalVisitor) visitRule(node antlr.RuleNode) interface{} {
	node.Accept(v)
	return nil
}

func (v *CalVisitor) VisitStart(ctx *parser.StartContext) interface{} {
	fmt.Println("VisitStart")
	return v.visitRule(ctx.Expr())
}

func (v *CalVisitor) VisitNumber(ctx *parser.NumberContext) interface{} {
	fmt.Println("VisitNumber")
	i, err := strconv.Atoi(ctx.NUMBER().GetText())
	if err != nil {
		panic(err.Error())
	}

	v.push(i)
	return nil
}

func (v *CalVisitor) VisitMulDiv(ctx *parser.MulDivContext) interface{} {
	fmt.Println("VisitMulDiv")
	//push expression result to Stack
	v.visitRule(ctx.Expr(0))
	v.visitRule(ctx.Expr(1))

	//push result to Stack
	var t antlr.Token = ctx.GetOp()
	right := v.Pop()
	left := v.Pop()
	switch t.GetTokenType() {
	case parser.CalcParserMUL:
		v.push(left * right)
	case parser.CalcParserDIV:
		v.push(left / right)
	default:
		panic("should not happen")

	}

	return nil
}

func (v *CalVisitor) VisitAddSub(ctx *parser.AddSubContext) interface{} {
	fmt.Println("VisitAddSub======")
	//push expression result to Stack
	v.visitRule(ctx.Expr(0))
	v.visitRule(ctx.Expr(1))

	//push result to Stack
	var t antlr.Token = ctx.GetOp()
	right := v.Pop()
	left := v.Pop()
	switch t.GetTokenType() {
	case parser.CalcParserADD:
		v.push(left + right)
	case parser.CalcParserSUB:
		v.push(left - right)
	default:
		panic("should not happen")
	}
	fmt.Println("VisitAddSub======exit")
	return nil

}

// 更偏向于中间递归 ( *, (2, 3) 这种

// 主要参考资料

// 加减乘除 扩充括号

// 跳过除法