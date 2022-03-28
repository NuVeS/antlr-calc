package my_listener

import (
	"fmt"
	"math"
	"strconv"

	"github.com/NuVeS/AntlrCalc/parser"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

type CalcListener struct {
	*parser.BaseCalculatorListener

	stack []int
}

func (l *CalcListener) push(i int) {
	l.stack = append(l.stack, i)
}

func (l *CalcListener) Pop() int {
	if len(l.stack) < 1 {
		panic("stack is empty unable to pop")
	}

	// Get the last value from the stack.
	result := l.stack[len(l.stack)-1]

	// Remove the last element from the stack.
	l.stack = l.stack[:len(l.stack)-1]

	return result
}

func (l *CalcListener) ExitMulDiv(c *parser.MulDivContext) {
	right, left := l.Pop(), l.Pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalculatorParserMUL:
		l.push(left * right)
	case parser.CalculatorParserDIV:
		l.push(left / right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *CalcListener) ExitAddSub(c *parser.AddSubContext) {
	right, left := l.Pop(), l.Pop()

	switch c.GetOp().GetTokenType() {
	case parser.CalculatorParserADD:
		l.push(left + right)
	case parser.CalculatorParserSUB:
		l.push(left - right)
	default:
		panic(fmt.Sprintf("unexpected op: %s", c.GetOp().GetText()))
	}
}

func (l *CalcListener) ExitNumber(c *parser.NumberContext) {
	i, err := strconv.Atoi(c.GetText())
	if err != nil {
		panic(err.Error())
	}

	l.push(i)
}

func (l *CalcListener) ExitPOWER(ctx *parser.POWERContext) {
	pow, num := l.Pop(), l.Pop()
	l.push(int(math.Pow(float64(num), float64(pow))))
}

func (l *CalcListener) ExitFACTORIAL(ctx *parser.FACTORIALContext) {
	num := l.Pop()
	res := 1
	for i := 1; i <= num; i++ {
		res *= i
	}
	l.push(res)
}

//

// VisitTerminal is called when a terminal node is visited.
func (l *CalcListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (l *CalcListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (l *CalcListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (l *CalcListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (l *CalcListener) EnterStart(ctx *parser.StartContext) {}

// ExitStart is called when production start is exited.
func (l *CalcListener) ExitStart(ctx *parser.StartContext) {}

// EnterBracet is called when production Bracet is entered.
func (l *CalcListener) EnterBracet(ctx *parser.BracetContext) {}

// ExitBracet is called when production Bracet is exited.
func (l *CalcListener) ExitBracet(ctx *parser.BracetContext) {}

// EnterNumber is called when production Number is entered.
func (l *CalcListener) EnterNumber(ctx *parser.NumberContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (l *CalcListener) EnterMulDiv(ctx *parser.MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (l *CalcListener) EnterAddSub(ctx *parser.AddSubContext) {}

// EnterPOWER is called when production POWER is entered.
func (l *CalcListener) EnterPOWER(ctx *parser.POWERContext) {}

// EnterFACTORIAL is called when production FACTORIAL is entered.
func (l *CalcListener) EnterFACTORIAL(ctx *parser.FACTORIALContext) {}
