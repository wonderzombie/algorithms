package main

import (
  "testing"
)

func assert(t bool, s string) {
	if !t {
		panic(s)
	}
}


func TestConstructCandidates(t *testing.T) {
	var (
		answer []int = make([]int, 1)
		c []int = make([]int, 2)
		ncandidates = 0
	)

	constructCandidates(answer, 0, 0, &c, &ncandidates)

	assert(ncandidates == 2, "Expected two ncandidates.")
	assert(c[0] == 1, "Expected c[0] to be 1.")
	assert(c[1] == 0, "Expected c[1] to be 0.")
}


func TestIsSolution(t *testing.T) {
	assert(isSolution(make([]int, 1), 1, 1),
			"Expected 1 == 1 to be valid solution.")
}
