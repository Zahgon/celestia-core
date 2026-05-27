// Package query implements the custom query format used to filter event
// subscriptions in CometBFT.
//
//	abci.invoice.number=22 AND abci.invoice.owner=Ivan
//
// Query expressions can handle attribute values encoding numbers, strings,
// dates, and timestamps.  The complete query grammar is described in the
// query/syntax package.
package query

import (
	"math/big"
	"regexp"
	"strings"
	"time"

	"github.com/cometbft/cometbft/abci/types"
	"github.com/cometbft/cometbft/libs/pubsub/query/syntax"
)

// All is a query that matches all events.
var All *Query

// A Query is the compiled form of a query.
type Query struct {
	ast   syntax.Query
	conds []condition
}

// New parses and compiles the query expression into an executable query.
func New(query string) (*Query, error) { _ = "STUB: not implemented"; return nil, nil }

// MustCompile compiles the query expression into an executable query.
// In case of error, MustCompile will panic.
//
// This is intended for use in program initialization; use query.New if you
// need to check errors.
func MustCompile(query string) *Query { _ = "STUB: not implemented"; return nil }

// Compile compiles the given query AST so it can be used to match events.
func Compile(ast syntax.Query) (*Query, error) { _ = "STUB: not implemented"; return nil, nil }

func ExpandEvents(flattenedEvents map[string][]string) []types.Event {
	_ = "STUB: not implemented"
	return nil
}

//nolint:prealloc

// Matches satisfies part of the pubsub.Query interface.  This implementation
// never reports an error. A nil *Query matches all events.
func (q *Query) Matches(events map[string][]string) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// String matches part of the pubsub.Query interface.
func (q *Query) String() string { _ = "STUB: not implemented"; return "" }

// Syntax returns the syntax tree representation of q.
func (q *Query) Syntax() syntax.Query { _ = "STUB: not implemented"; return *new(syntax.Query) }

// matchesEvents reports whether all the conditions match the given events.
func (q *Query) matchesEvents(events []types.Event) bool { _ = "STUB: not implemented"; return false }

// A condition is a compiled match condition.  A condition matches an event if
// the event has the designated type, contains an attribute with the given
// name, and the match function returns true for the attribute value.
type condition struct {
	tag   string // e.g., "tx.hash"
	match func(s string) bool
}

// findAttr returns a slice of attribute values from event matching the
// condition tag, and reports whether the event type strictly equals the
// condition tag.
func (c condition) findAttr(event types.Event) ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// type does not match tag

// type == tag

// matchesAny reports whether c matches at least one of the given events.
func (c condition) matchesAny(events []types.Event) bool { _ = "STUB: not implemented"; return false }

// matchesEvent reports whether c matches the given event.
func (c condition) matchesEvent(event types.Event) bool { _ = "STUB: not implemented"; return false }

// As a special case, a condition tag that exactly matches the event type
// is matched against an empty string. This allows existence checks to
// work for type-only queries.

// At this point, we have candidate values.

func compileCondition(cond syntax.Condition) (condition, error) {
	_ = "STUB: not implemented"
	return *new(condition), nil
}

// Handle existence checks separately to simplify the logic below for
// comparisons that take arguments.

// All the other operators require an argument.

// Precompile the argument value matcher.

// We use this regex to support queries of the form "8atom", "6.5stake",
// which are actively used in production.
// The regex takes care of removing the non-number suffix.
var extractNum = regexp.MustCompile(`^\d+(\.\d+)?`)

func parseNumber(s string) (*big.Float, error) { _ = "STUB: not implemented"; return nil, nil }

// A map of operator ⇒ argtype ⇒ match-constructor.
// An entry does not exist if the combination is not valid.
//
// Disable the dupl lint for this map. The result isn't even correct.
//
//nolint:dupl
var opTypeMap = map[syntax.Token]map[syntax.Token]func(interface{}) func(string) bool{
	syntax.TContains: {
		syntax.TString: func(v interface{}) func(string) bool {
			return func(s string) bool {
				return strings.Contains(s, v.(string))
			}
		},
	},
	syntax.TEq: {
		syntax.TString: func(v interface{}) func(string) bool {
			return func(s string) bool { return s == v.(string) }
		},
		syntax.TNumber: func(v interface{}) func(string) bool {
			return func(s string) bool {
				w, err := parseNumber(s)
				return err == nil && w.Cmp(v.(*big.Float)) == 0
			}
		},
		syntax.TDate: func(v interface{}) func(string) bool {
			return func(s string) bool {
				ts, err := syntax.ParseDate(s)
				return err == nil && ts.Equal(v.(time.Time))
			}
		},
		syntax.TTime: func(v interface{}) func(string) bool {
			return func(s string) bool {
				ts, err := syntax.ParseTime(s)
				return err == nil && ts.Equal(v.(time.Time))
			}
		},
	},
	syntax.TLt: {
		syntax.TNumber: func(v interface{}) func(string) bool {
			return func(s string) bool {
				w, err := parseNumber(s)
				return err == nil && w.Cmp(v.(*big.Float)) < 0
			}
		},
		syntax.TDate: func(v interface{}) func(string) bool {
			return func(s string) bool {
				ts, err := syntax.ParseDate(s)
				return err == nil && ts.Before(v.(time.Time))
			}
		},
		syntax.TTime: func(v interface{}) func(string) bool {
			return func(s string) bool {
				ts, err := syntax.ParseTime(s)
				return err == nil && ts.Before(v.(time.Time))
			}
		},
	},
	syntax.TLeq: {
		syntax.TNumber: func(v interface{}) func(string) bool {
			return func(s string) bool {
				w, err := parseNumber(s)
				return err == nil && w.Cmp(v.(*big.Float)) <= 0
			}
		},
		syntax.TDate: func(v interface{}) func(string) bool {
			return func(s string) bool {
				ts, err := syntax.ParseDate(s)
				return err == nil && !ts.After(v.(time.Time))
			}
		},
		syntax.TTime: func(v interface{}) func(string) bool {
			return func(s string) bool {
				ts, err := syntax.ParseTime(s)
				return err == nil && !ts.After(v.(time.Time))
			}
		},
	},
	syntax.TGt: {
		syntax.TNumber: func(v interface{}) func(string) bool {
			return func(s string) bool {
				w, err := parseNumber(s)
				return err == nil && w.Cmp(v.(*big.Float)) > 0
			}
		},
		syntax.TDate: func(v interface{}) func(string) bool {
			return func(s string) bool {
				ts, err := syntax.ParseDate(s)
				return err == nil && ts.After(v.(time.Time))
			}
		},
		syntax.TTime: func(v interface{}) func(string) bool {
			return func(s string) bool {
				ts, err := syntax.ParseTime(s)
				return err == nil && ts.After(v.(time.Time))
			}
		},
	},
	syntax.TGeq: {
		syntax.TNumber: func(v interface{}) func(string) bool {
			return func(s string) bool {
				w, err := parseNumber(s)
				return err == nil && w.Cmp(v.(*big.Float)) >= 0
			}
		},
		syntax.TDate: func(v interface{}) func(string) bool {
			return func(s string) bool {
				ts, err := syntax.ParseDate(s)
				return err == nil && !ts.Before(v.(time.Time))
			}
		},
		syntax.TTime: func(v interface{}) func(string) bool {
			return func(s string) bool {
				ts, err := syntax.ParseTime(s)
				return err == nil && !ts.Before(v.(time.Time))
			}
		},
	},
}
