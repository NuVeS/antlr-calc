// Code generated from Calculator.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Calculator

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseCalculatorListener is a complete listener for a parse tree produced by CalculatorParser.
type BaseCalculatorListener struct{}

var _ CalculatorListener = &BaseCalculatorListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCalculatorListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCalculatorListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCalculatorListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCalculatorListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseCalculatorListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseCalculatorListener) ExitStart(ctx *StartContext) {}

// EnterBracet is called when production Bracet is entered.
func (s *BaseCalculatorListener) EnterBracet(ctx *BracetContext) {}

// ExitBracet is called when production Bracet is exited.
func (s *BaseCalculatorListener) ExitBracet(ctx *BracetContext) {}

// EnterNumber is called when production Number is entered.
func (s *BaseCalculatorListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production Number is exited.
func (s *BaseCalculatorListener) ExitNumber(ctx *NumberContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseCalculatorListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseCalculatorListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseCalculatorListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseCalculatorListener) ExitAddSub(ctx *AddSubContext) {}

// EnterPOWER is called when production POWER is entered.
func (s *BaseCalculatorListener) EnterPOWER(ctx *POWERContext) {}

// ExitPOWER is called when production POWER is exited.
func (s *BaseCalculatorListener) ExitPOWER(ctx *POWERContext) {}

// EnterFACTORIAL is called when production FACTORIAL is entered.
func (s *BaseCalculatorListener) EnterFACTORIAL(ctx *FACTORIALContext) {}

// ExitFACTORIAL is called when production FACTORIAL is exited.
func (s *BaseCalculatorListener) ExitFACTORIAL(ctx *FACTORIALContext) {}
