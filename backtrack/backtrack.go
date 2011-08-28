package main

import (
	"fmt"
)

const MAX_CANDIDATES = 2
var finished = false

func isSolution(a []int, k int, n int) bool {
	return k == n
}


func processSolution(a []int, k int) {
	fmt.Printf("{")
	for i := 0; i <= k; i++ {
		if (a[i] == 1) {
			fmt.Printf(" %d", i)
		}
	}
	fmt.Printf(" }\n")
}


func constructCandidates(a []int, k int, n int, c *[]int,
		ncandidates *int) {
	(*c)[0] = 1
	(*c)[1] = 0

	*ncandidates = 2
}


func backtrack(a []int, k int, n int) {
	var c = make([]int, MAX_CANDIDATES)
	var ncandidates int


	if isSolution(a, k, n) {
		fmt.Printf("Solution: %v\n", a)
		processSolution(a, k)
	} else {
		k += 1
		fmt.Printf("a is not a solution: %v\n", a)
		fmt.Printf("Backtracking with k == %v\n", k)
		constructCandidates(a, k, n, &c, &ncandidates)
		for i := 0; i < ncandidates; i++ {
			fmt.Printf("Iterating. i is now %v\n", i)
			fmt.Printf("a before: %v\n", a)
			a[k] = c[i]
			fmt.Printf("a after: %v\n", a)
			backtrack(a, k, n)
			if finished { return }
		}
	}
}


func main() {
	var NMAX = 4
	backtrack(make([]int, NMAX), 0, 3)
}
