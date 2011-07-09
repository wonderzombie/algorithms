package main

import (
  "flag"
  "fmt"
)

func isEven(x int) bool {
  return (x % 2 == 0)
}

func nextCycle(i int) int {
  if isEven(i) {
    i /= 2
  } else {
    i = 3 * i + 1;
  }
  return i
}

func findInArray(needle int, haystack []int) (int, bool) {
  for i := 0; i < len(haystack); i++ {
    if haystack[i] == needle {
      return i, true
    }
  }
  return -1, false
}

func calculateCycleLengthWithMemo(i int, memo []int) (int, []int) {
  length := 1
  newMemo := make([]int, 0)

  for i != 1 {
    i = nextCycle(i)
    length += 1
    newMemo = append(newMemo, i)
  }

  return length, newMemo
}

func calculateCycleLength(i int) int {
  length := 1

  for i != 1 {
    i = nextCycle(i)
    length += 1
  }

  return length
}

func calculateLongestCycleLength(min int, max int) int {
  highestLength := 0

  for i := min; i <= max; i++ {
    length := calculateCycleLength(i)
    fmt.Printf("Length for %v was %v.\n", i, length)
    if length > highestLength {
      highestLength = length
    }
  }

  return highestLength
}

func main() {
  var max, min int
  flag.IntVar(&min, "min", 100, "integer to start with")
  flag.IntVar(&max, "max", 200, "integer to end with")
  flag.Parse()

  if max < min {
    panic(fmt.Sprintf("Max was less than min!"))
  }

  highestLength := calculateLongestCycleLength(min, max)

  fmt.Printf("highest length for %v - %v was: %v\n",
             min, max, highestLength)
}
