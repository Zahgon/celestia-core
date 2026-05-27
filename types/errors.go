package types

type (
	// ErrInvalidCommitHeight is returned when we encounter a commit with an
	// unexpected height.
	ErrInvalidCommitHeight struct {
		Expected int64
		Actual   int64
	}

	// ErrInvalidCommitSignatures is returned when we encounter a commit where
	// the number of signatures doesn't match the number of validators.
	ErrInvalidCommitSignatures struct {
		Expected int
		Actual   int
	}
)

func NewErrInvalidCommitHeight(expected, actual int64) ErrInvalidCommitHeight {
	_ = "STUB: not implemented"
	return *new(ErrInvalidCommitHeight)
}

func (e ErrInvalidCommitHeight) Error() string { _ = "STUB: not implemented"; return "" }

func NewErrInvalidCommitSignatures(expected, actual int) ErrInvalidCommitSignatures {
	_ = "STUB: not implemented"
	return *new(ErrInvalidCommitSignatures)
}

func (e ErrInvalidCommitSignatures) Error() string { _ = "STUB: not implemented"; return "" }
