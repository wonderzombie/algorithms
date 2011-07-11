package main

import (
//	"fmt"
	"testing"
)

func TestTallyVotes(t *testing.T) {
	candidates := []string { "Alice", "Bob", "Chuck" }
	ballots := []Ballot {
		Ballot { 1, 2, 3 },
		Ballot { 2, 3, 1 },
		Ballot { 1, 2, 3 },
		Ballot { 1, 3, 2 },
	}

	tally := tallyVotes(candidates, ballots, 1)

	if tally["Alice"] != 3 || tally["Bob"] != 1 || tally["Chuck"] != 0 {
		t.Fatalf("Expected Alice: 3, Bob: 1, and Chuck: 0. Actual: %v", tally)
	}
}
