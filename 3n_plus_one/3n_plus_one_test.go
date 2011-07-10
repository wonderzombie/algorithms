package main

import (
  "fmt"
  "testing"
)

///
// findInArray
//

func TestFindInArray(t *testing.T) {
  arr := []int { 1, 2, 3 }
  expectedIndex := 1

  actualIndex, found := findInArray(2, arr)
  if !found {
    t.Fatal("Expected 2 to be found in array %v.", arr)
  }
  if expectedIndex != actualIndex {
    t.Fatal("Expected 2 at index %v, actual: %v", 
            expectedIndex, actualIndex)
  }
}

func TestFindInEmptyArray(t *testing.T) {
  arr := []int {}

  index, found := findInArray(2, arr)
  if found {
    t.Fatal("Did not expect to find anything in empty array.")
  }
  if index != -1 {
    t.Fatal("Expected index to be -1, actual %v", index)
  }
}


///
// nextCycle
//

func TestNextCycleOdd(t *testing.T) {
  i := 5
  expected := i * 3 + 1
  actual := nextCycle(i)

  if expected != actual {
    t.Fatal("Expected %v, actual %v", expected, actual)
  }
}

func TestNextCycleEven(t *testing.T) {
  i := 4
  expected := i / 2
  actual := nextCycle(i)

  if expected != actual {
    t.Fatal("Expected %v, actual %v", expected, actual)
  }
}

func TestCalculateCycleLength(t *testing.T) {
  i := 22
  expected := 16

  actual := calculateCycleLength(i)
  if expected != actual {
    t.Fatal("For %v, expected cycle length: %v, actual: %v",
            i, expected, actual)
  }
}

func TestCalculateCycleLengthWithMemo(t *testing.T) {
  i := 22
  expected := 16
  memo := make([]int, 0)

  actual, newMemo := calculateCycleLengthWithMemo(i, memo)
  if newMemo == nil {
    t.Fatal("Expected returned memo not to be nil.")
  } else {
    fmt.Printf("Returned memo: %+v\n", newMemo)
  }

  if expected != actual {
    t.Fatalf("For %v, expected cycle length: %v, actual: %v",
            i, expected, actual)
  }

}


///
// calculateLongestCycleLength
//

func TestCalculateLongestCycleLength(t *testing.T) {
  min, max := 1, 10
  expected := 20
  actual := calculateLongestCycleLength(min, max)
  if expected != actual {
    t.Fatal("Expected length for %v - %v: %v, actual: %v",
            min, max, expected, actual)
  }
}
