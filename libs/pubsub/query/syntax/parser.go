package syntax

import (
	"io"
	"math/big"
	"time"
)

// Parse parses the specified query string. It is shorthand for constructing a
// parser for s and calling its Parse method.
func Parse(s string) (Query, error) { _ = "STUB: not implemented"; return *new(Query), nil }

// Query is the root of the parse tree for a query.  A query is the conjunction
// of one or more conditions.
type Query []Condition

func (q Query) String() string { _ = "STUB: not implemented"; return "" }

// A Condition is a single conditional expression, consisting of a tag, a
// comparison operator, and an optional argument. The type of the argument
// depends on the operator.
type Condition struct {
	Tag string
	Op  Token
	Arg *Arg

	opText string
}

func (c Condition) String() string { _ = "STUB: not implemented"; return "" }

// An Arg is the argument of a comparison operator.
type Arg struct {
	Type Token
	text string
}

func (a *Arg) String() string { _ = "STUB: not implemented"; return "" }

// Number returns the value of the argument text as a number, or nil if the
// text does not encode a valid number value.
func (a *Arg) Number() *big.Float { _ = "STUB: not implemented"; return nil }

// If it is indeed a big integer, we make sure to convert it to a float with enough precision
// to represent all the bits

// Time returns the value of the argument text as a time, or the zero value if
// the text does not encode a timestamp or datestamp.
func (a *Arg) Time() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Value returns the value of the argument text as a string, or "".
func (a *Arg) Value() string { _ = "STUB: not implemented"; return "" }

// Parser is a query expression parser. The grammar for query expressions is
// defined in the syntax package documentation.
type Parser struct {
	scanner *Scanner
}

// NewParser constructs a new parser that reads the input from r.
func NewParser(r io.Reader) *Parser { _ = "STUB: not implemented"; return nil }

// Parse parses the complete input and returns the resulting query.
func (p *Parser) Parse() (Query, error) { _ = "STUB: not implemented"; return *new(Query), nil }

// parseCond parses a conditional expression: tag OP value.
func (p *Parser) parseCond() (Condition, error) {
	_ = "STUB: not implemented"
	return *new(Condition), nil
}

// no argument

// require advances the scanner and requires that the resulting token is one of
// the specified token types.
func (p *Parser) require(tokens ...Token) error { _ = "STUB: not implemented"; return nil }

// tokLabel makes a human-readable summary string for the given token types.
func tokLabel(tokens []Token) string { _ = "STUB: not implemented"; return "" }

// ParseDate parses s as a date string in the format used by DATE values.
func ParseDate(s string) (time.Time, error) { _ = "STUB: not implemented"; return *new(time.Time), nil }

// ParseTime parses s as a timestamp in the format used by TIME values.
func ParseTime(s string) (time.Time, error) { _ = "STUB: not implemented"; return *new(time.Time), nil }
