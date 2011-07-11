package main

import (
//	"fmt"
)

type Ballot []int

func tallyVotes(candidates []string, ballots []Ballot, round int) map[string] int {
	tally := make(map[string] int)
	for _, candidate := range candidates {
		tally[candidate] = 0
	}

	for _, b := range ballots {
		i := b[round - 1]
		cand := candidates[i - 1]
		tally[cand] += 1
	}

	return tally
}

func main() {

}
