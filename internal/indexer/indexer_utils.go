package indexer

import (
	"math/big"

	"github.com/cometbft/cometbft/state/indexer"
)

// If the actual event value is a float, we get the condition and parse it as a float
// to compare against
func compareFloat(op1 *big.Float, op2 interface{}) (int, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

// If the event value we compare against the condition (op2) is an integer
// we convert the int to float with a precision equal to the number of bits
// needed to represent the integer to avoid rounding issues with floats
// where 100 would equal to 100.2 because 100.2 is rounded to 100, while 100.7
// would be rounded to 101.
func compareInt(op1 *big.Int, op2 interface{}) (int, bool, error) {
	_ = "STUB: not implemented"
	return 0, false, nil
}

func CheckBounds(ranges indexer.QueryRange, v interface{}) (bool, error) {
	_ = "STUB: not implemented"
	// These functions fetch the lower and upper bounds of the query
	// It is expected that for x > 5, the value of lowerBound is 6.
	// This is achieved by adding one to the actual lower bound.
	// For a query of x < 5, the value of upper bound is 4.
	// This is achieved by subtracting one from the actual upper bound.
	return false, nil
}

// For integers this behavior will work. However, for floats, we cannot simply add/sub 1.
// Query :x < 5.5 ; x = 5 should match the query. If we subtracted one as for integers,
// the upperBound would be 4.5 and x would not match.  Thus we do not subtract anything for
// floating point bounds.

// We can rewrite these functions to not add/sub 1 but the function handles also time arguments.
// To be sure we are not breaking existing queries that compare time, and as we are planning to replace
// the indexer in the future, we adapt the code here to handle floats as a special case.

// *Explanation for the isFloat condition below.*
// In LowerBoundValue(), for floating points, we cannot simply add 1 due to the reasons explained in
// in the comment at the beginning. The same is true for subtracting one for UpperBoundValue().
// That means that for integers, if the condition is >=, cmp will be either 0 or 1
// ( cmp == -1 should always be false).
// 	But if the lowerBound is a float, we have not subtracted one, so returning a 0
// is correct only if ranges.IncludeLowerBound is true.
// example int: x < 100; upperBound = 99; if x.Cmp(99) == 0 the condition holds
// example float: x < 100.0; upperBound = 100.0; if x.Cmp(100) ==0 then returning x
// would be wrong.
