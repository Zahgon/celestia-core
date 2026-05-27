package main

import (
	"math/rand"
)

// combinations takes input in the form of a map of item lists, and returns a
// list of all combinations of each item for each key. E.g.:
//
// {"foo": [1, 2, 3], "bar": [4, 5, 6]}
//
// Will return the following maps:
//
// {"foo": 1, "bar": 4}
// {"foo": 1, "bar": 5}
// {"foo": 1, "bar": 6}
// {"foo": 2, "bar": 4}
// {"foo": 2, "bar": 5}
// {"foo": 2, "bar": 6}
// {"foo": 3, "bar": 4}
// {"foo": 3, "bar": 5}
// {"foo": 3, "bar": 6}
func combinations(items map[string][]interface{}) []map[string]interface{} {
	_ = "STUB: not implemented"
	//nolint:prealloc
	return nil
}

// combiner is a utility function for combinations.
func combiner(head map[string]interface{}, pending []string, items map[string][]interface{}) []map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

// uniformChoice chooses a single random item from the argument list, uniformly weighted.
type uniformChoice []interface{}

func (uc uniformChoice) Choose(r *rand.Rand) interface{} { _ = "STUB: not implemented"; return nil }

// probSetChoice picks a set of strings based on each string's probability (0-1).
type probSetChoice map[string]float64

func (pc probSetChoice) Choose(r *rand.Rand) []string { _ = "STUB: not implemented"; return nil }

// uniformSetChoice picks a set of strings with uniform probability, picking at least one.
type uniformSetChoice []string

func (usc uniformSetChoice) Choose(r *rand.Rand) []string {
	_ = "STUB: not implemented"
	//nolint:prealloc
	return nil
}

// weightedChoice chooses a single random key from a map of keys and weights.
type weightedChoice map[interface{}]uint

func (wc weightedChoice) Choose(r *rand.Rand) interface{} { _ = "STUB: not implemented"; return nil }
