package strings

// StringInSlice returns true if a is found the list.
func StringInSlice(a string, list []string) bool { _ = "STUB: not implemented"; return false }

// SplitAndTrim slices s into all subslices separated by sep and returns a
// slice of the string s with all leading and trailing Unicode code points
// contained in cutset removed. If sep is empty, SplitAndTrim splits after each
// UTF-8 sequence. First part is equivalent to strings.SplitN with a count of
// -1.
func SplitAndTrim(s, sep, cutset string) []string { _ = "STUB: not implemented"; return nil }

// SplitAndTrimEmpty slices s into all subslices separated by sep and returns a
// slice of the string s with all leading and trailing Unicode code points
// contained in cutset removed. If sep is empty, SplitAndTrim splits after each
// UTF-8 sequence. First part is equivalent to strings.SplitN with a count of
// -1.  also filter out empty strings, only return non-empty strings.
func SplitAndTrimEmpty(s, sep, cutset string) []string { _ = "STUB: not implemented"; return nil }

// Returns true if s is a non-empty printable non-tab ascii character.
func IsASCIIText(s string) bool { _ = "STUB: not implemented"; return false }

// NOTE: Assumes that s is ASCII as per IsASCIIText(), otherwise panics.
func ASCIITrim(s string) string { _ = "STUB: not implemented"; return "" }

// skip space

// StringSliceEqual checks if string slices a and b are equal
func StringSliceEqual(a, b []string) bool { _ = "STUB: not implemented"; return false }
