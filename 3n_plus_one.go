package main

import (
  "flag"
  "fmt"
)

type CycleLookup []int

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

func calculateCycleLengthWithMemo(i int, memo []int) (int, []int) {
  length := 0

  for i != 1 {
    // fmt.Printf("%d ", i)
    i = nextCycle(i)
    length += 1
  }

  return length, memo
}

func calculateCycleLength(i int) int {
  var length = 0

  for i != 1 {
    // fmt.Printf("%d ", i)
    i = nextCycle(i)
    length += 1
  }

  return length
}

func main() {
  var max, min int
  flag.IntVar(&min, "min", 100, "integer to start with")
  flag.IntVar(&max, "max", 200, "integer to end with")
  flag.Parse()

  if max < min {
    panic(fmt.Sprintf("Max was less than min!"))
  }

  fmt.Printf("Max: %d, Min: %d\n", max, min)

  highestLength := 0

  for i := min; i <= max; i++ {
    length := calculateCycleLength(i)
    fmt.Printf("Length for %v was %v.\n", i, length)
    if length > highestLength {
      highestLength = length
    }
  }

  fmt.Printf("highest length for %v - %v was: %v\n",
             min, max, highestLength)
}
