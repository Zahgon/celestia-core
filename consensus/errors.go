package consensus

type ErrInvalidVote struct {
	Reason string
}

func (e ErrInvalidVote) Error() string { _ = "STUB: not implemented"; return "" }
