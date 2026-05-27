package time

import (
	"time"
)

// Now returns the current time in UTC with no monotonic component.
func Now() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

// Canonical returns UTC time with no monotonic component.
// Stripping the monotonic component is for time equality.
// See https://github.com/tendermint/tendermint/pull/2203#discussion_r215064334
func Canonical(t time.Time) time.Time {
	_ = "STUB: not implemented"
	return *

	// WeightedTime for computing a median.
	new(time.Time)
}

type WeightedTime struct {
	Time   time.Time
	Weight int64
}

// NewWeightedTime with time and weight.
func NewWeightedTime(time time.Time, weight int64) *WeightedTime {
	_ = "STUB: not implemented"
	return nil
}

// WeightedMedian computes weighted median time for a given array of WeightedTime and the total voting power.
func WeightedMedian(weightedTimes []*WeightedTime, totalVotingPower int64) (res time.Time) {
	_ = "STUB: not implemented"
	return *new(time.Time)
}
