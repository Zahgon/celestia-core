package math

import (
	"errors"
)

var ErrOverflowInt32 = errors.New("int32 overflow")
var ErrOverflowUint8 = errors.New("uint8 overflow")
var ErrOverflowInt8 = errors.New("int8 overflow")

// SafeAddInt32 adds two int32 integers
// If there is an overflow this will panic
func SafeAddInt32(a, b int32) int32 { _ = "STUB: not implemented"; return 0 }

// SafeSubInt32 subtracts two int32 integers
// If there is an overflow this will panic
func SafeSubInt32(a, b int32) int32 { _ = "STUB: not implemented"; return 0 }

// SafeConvertInt32 takes a int and checks if it overflows
// If there is an overflow this will panic
func SafeConvertInt32(a int64) int32 { _ = "STUB: not implemented"; return 0 }

// SafeConvertUint8 takes an int64 and checks if it overflows
// If there is an overflow it returns an error
func SafeConvertUint8(a int64) (uint8, error) { _ = "STUB: not implemented"; return 0, nil }

// SafeConvertInt8 takes an int64 and checks if it overflows
// If there is an overflow it returns an error
func SafeConvertInt8(a int64) (int8, error) { _ = "STUB: not implemented"; return 0, nil }
