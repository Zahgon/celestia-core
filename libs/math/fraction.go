package math

// Fraction defined in terms of a numerator divided by a denominator in uint64
// format. Fraction must be positive.
type Fraction struct {
	// The portion of the denominator in the faction, e.g. 2 in 2/3.
	Numerator uint64 `json:"numerator"`
	// The value by which the numerator is divided, e.g. 3 in 2/3.
	Denominator uint64 `json:"denominator"`
}

func (fr Fraction) String() string { _ = "STUB: not implemented"; return "" }

// ParseFractions takes the string of a fraction as input i.e "2/3" and converts this
// to the equivalent fraction else returns an error. The format of the string must be
// one number followed by a slash (/) and then the other number.
func ParseFraction(f string) (Fraction, error) {
	_ = "STUB: not implemented"
	return *new(Fraction), nil
}
