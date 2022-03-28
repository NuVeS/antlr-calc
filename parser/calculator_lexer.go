// Code generated from Calculator.g4 by ANTLR 4.9. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 12, 51, 8,
	1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9,
	7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 3, 2, 3, 2, 3, 3,
	3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3, 6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9,
	3, 9, 3, 10, 6, 10, 41, 10, 10, 13, 10, 14, 10, 42, 3, 11, 6, 11, 46, 10,
	11, 13, 11, 14, 11, 47, 3, 11, 3, 11, 2, 2, 12, 3, 3, 5, 4, 7, 5, 9, 6,
	11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 3, 2, 4, 3, 2, 50, 59, 5,
	2, 11, 12, 15, 15, 34, 34, 2, 52, 2, 3, 3, 2, 2, 2, 2, 5, 3, 2, 2, 2, 2,
	7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3, 2, 2, 2, 2, 13, 3, 2, 2, 2,
	2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19, 3, 2, 2, 2, 2, 21, 3, 2, 2,
	2, 3, 23, 3, 2, 2, 2, 5, 25, 3, 2, 2, 2, 7, 27, 3, 2, 2, 2, 9, 29, 3, 2,
	2, 2, 11, 31, 3, 2, 2, 2, 13, 33, 3, 2, 2, 2, 15, 35, 3, 2, 2, 2, 17, 37,
	3, 2, 2, 2, 19, 40, 3, 2, 2, 2, 21, 45, 3, 2, 2, 2, 23, 24, 7, 44, 2, 2,
	24, 4, 3, 2, 2, 2, 25, 26, 7, 49, 2, 2, 26, 6, 3, 2, 2, 2, 27, 28, 7, 45,
	2, 2, 28, 8, 3, 2, 2, 2, 29, 30, 7, 47, 2, 2, 30, 10, 3, 2, 2, 2, 31, 32,
	7, 35, 2, 2, 32, 12, 3, 2, 2, 2, 33, 34, 7, 42, 2, 2, 34, 14, 3, 2, 2,
	2, 35, 36, 7, 43, 2, 2, 36, 16, 3, 2, 2, 2, 37, 38, 7, 96, 2, 2, 38, 18,
	3, 2, 2, 2, 39, 41, 9, 2, 2, 2, 40, 39, 3, 2, 2, 2, 41, 42, 3, 2, 2, 2,
	42, 40, 3, 2, 2, 2, 42, 43, 3, 2, 2, 2, 43, 20, 3, 2, 2, 2, 44, 46, 9,
	3, 2, 2, 45, 44, 3, 2, 2, 2, 46, 47, 3, 2, 2, 2, 47, 45, 3, 2, 2, 2, 47,
	48, 3, 2, 2, 2, 48, 49, 3, 2, 2, 2, 49, 50, 8, 11, 2, 2, 50, 22, 3, 2,
	2, 2, 5, 2, 42, 47, 3, 8, 2, 2,
}

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "'*'", "'/'", "'+'", "'-'", "'!'", "'('", "')'", "'^'",
}

var lexerSymbolicNames = []string{
	"", "MUL", "DIV", "ADD", "SUB", "FACT", "OP_BRACETS", "CL_BRACETS", "POW",
	"NUMBER", "WHITESPACE",
}

var lexerRuleNames = []string{
	"MUL", "DIV", "ADD", "SUB", "FACT", "OP_BRACETS", "CL_BRACETS", "POW",
	"NUMBER", "WHITESPACE",
}

type CalculatorLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

// NewCalculatorLexer produces a new lexer instance for the optional input antlr.CharStream.
//
// The *CalculatorLexer instance produced may be reused by calling the SetInputStream method.
// The initial lexer configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewCalculatorLexer(input antlr.CharStream) *CalculatorLexer {
	l := new(CalculatorLexer)
	lexerDeserializer := antlr.NewATNDeserializer(nil)
	lexerAtn := lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)
	lexerDecisionToDFA := make([]*antlr.DFA, len(lexerAtn.DecisionToState))
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Calculator.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// CalculatorLexer tokens.
const (
	CalculatorLexerMUL        = 1
	CalculatorLexerDIV        = 2
	CalculatorLexerADD        = 3
	CalculatorLexerSUB        = 4
	CalculatorLexerFACT       = 5
	CalculatorLexerOP_BRACETS = 6
	CalculatorLexerCL_BRACETS = 7
	CalculatorLexerPOW        = 8
	CalculatorLexerNUMBER     = 9
	CalculatorLexerWHITESPACE = 10
)
