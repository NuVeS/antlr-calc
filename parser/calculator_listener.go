// Code generated from Calculator.g4 by ANTLR 4.9. DO NOT EDIT.

package parser // Calculator

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CalculatorListener is a complete listener for a parse tree produced by CalculatorParser.
type CalculatorListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterBracet is called when entering the Bracet production.
	EnterBracet(c *BracetContext)

	// EnterNumber is called when entering the Number production.
	EnterNumber(c *NumberContext)

	// EnterMulDiv is called when entering the MulDiv production.
	EnterMulDiv(c *MulDivContext)

	// EnterAddSub is called when entering the AddSub production.
	EnterAddSub(c *AddSubContext)

	// EnterPOWER is called when entering the POWER production.
	EnterPOWER(c *POWERContext)

	// EnterFACTORIAL is called when entering the FACTORIAL production.
	EnterFACTORIAL(c *FACTORIALContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitBracet is called when exiting the Bracet production.
	ExitBracet(c *BracetContext)

	// ExitNumber is called when exiting the Number production.
	ExitNumber(c *NumberContext)

	// ExitMulDiv is called when exiting the MulDiv production.
	ExitMulDiv(c *MulDivContext)

	// ExitAddSub is called when exiting the AddSub production.
	ExitAddSub(c *AddSubContext)

	// ExitPOWER is called when exiting the POWER production.
	ExitPOWER(c *POWERContext)

	// ExitFACTORIAL is called when exiting the FACTORIAL production.
	ExitFACTORIAL(c *FACTORIALContext)
}
