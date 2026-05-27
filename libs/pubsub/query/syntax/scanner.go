package syntax

import (
	"bufio"
	"bytes"
	"io"
	"time"
)

// Token is the type of a lexical token in the query grammar.
type Token byte

const (
	TInvalid  = iota // invalid or unknown token
	TTag             // field tag: x.y
	TString          // string value: 'foo bar'
	TNumber          // number: 0, 15.5, 100
	TTime            // timestamp: TIME yyyy-mm-ddThh:mm:ss([-+]hh:mm|Z)
	TDate            // datestamp: DATE yyyy-mm-dd
	TAnd             // operator: AND
	TContains        // operator: CONTAINS
	TExists          // operator: EXISTS
	TEq              // operator: =
	TLt              // operator: <
	TLeq             // operator: <=
	TGt              // operator: >
	TGeq             // operator: >=

	// Do not reorder these values without updating the scanner code.
)

var tString = [...]string{
	TInvalid:  "invalid token",
	TTag:      "tag",
	TString:   "string",
	TNumber:   "number",
	TTime:     "timestamp",
	TDate:     "datestamp",
	TAnd:      "AND operator",
	TContains: "CONTAINS operator",
	TExists:   "EXISTS operator",
	TEq:       "= operator",
	TLt:       "< operator",
	TLeq:      "<= operator",
	TGt:       "> operator",
	TGeq:      ">= operator",
}

func (t Token) String() string { _ = "STUB: not implemented"; return "" }

const (
	// TimeFormat is the format string used for timestamp values.
	TimeFormat = time.RFC3339

	// DateFormat is the format string used for datestamp values.
	DateFormat = "2006-01-02"
)

// Scanner reads lexical tokens of the query language from an input stream.
// Each call to Next advances the scanner to the next token, or reports an
// error.
type Scanner struct {
	r   *bufio.Reader
	buf bytes.Buffer
	tok Token
	err error

	pos, last, end int
}

// NewScanner constructs a new scanner that reads from r.
func NewScanner(r io.Reader) *Scanner { _ = "STUB: not implemented"; return nil }

// Next advances s to the next token in the input, or reports an error.  At the
// end of input, Next returns io.EOF.
func (s *Scanner) Next() error { _ = "STUB: not implemented"; return nil }

// skip whitespace

// Token returns the type of the current input token.
func (s *Scanner) Token() Token {
	_ = "STUB: not implemented"

	// Text returns the text of the current input token.
	return *new(Token)
}

func (s *Scanner) Text() string { _ = "STUB: not implemented"; return "" }

// Pos returns the start offset of the current token in the input.
func (s *Scanner) Pos() int {
	_ = "STUB: not implemented"

	// Err returns the last error reported by Next, if any.
	return 0
}

func (s *Scanner) Err() error {
	_ = "STUB: not implemented"

	// scanNumber scans for numbers with optional fractional parts.
	// Examples: 0, 1, 3.14
	return nil
}

func (s *Scanner) scanNumber(first rune) error { _ = "STUB: not implemented"; return nil }

func (s *Scanner) scanString(first rune) error {
	_ = "STUB: not implemented"
	// discard opening quote
	return nil
}

// discard closing quote

func (s *Scanner) scanCompare(first rune) error { _ = "STUB: not implemented"; return nil }

// the assigned token is correct

// depends on token order

func (s *Scanner) scanTagLike(first rune) error { _ = "STUB: not implemented"; return nil }

// to check for TIME, DATE

func (s *Scanner) scanTimestamp() error {
	_ = "STUB: not implemented"
	// discard "TIME" label
	return nil
}

func (s *Scanner) scanDatestamp() error {
	_ = "STUB: not implemented"
	// discard "DATE" label
	return nil
}

func (s *Scanner) scanWhile(ok func(rune) bool) error { _ = "STUB: not implemented"; return nil }

func (s *Scanner) rune() (rune, error) { _ = "STUB: not implemented"; return 0, nil }

func (s *Scanner) unrune() { _ = "STUB: not implemented"; return }

func (s *Scanner) fail(err error) error { _ = "STUB: not implemented"; return nil }

func (s *Scanner) invalid(ch rune) error { _ = "STUB: not implemented"; return nil }

func isDigit(r rune) bool { _ = "STUB: not implemented"; return false }

func isTagRune(r rune) bool { _ = "STUB: not implemented"; return false }

func isFirstTagRune(r rune) bool { _ = "STUB: not implemented"; return false }

func isTimeRune(r rune) bool { _ = "STUB: not implemented"; return false }

func isDateRune(r rune) bool { _ = "STUB: not implemented"; return false }
