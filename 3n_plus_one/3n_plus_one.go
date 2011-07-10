package main

import (
  "fmt"
)

type Pair struct {
  min int
  max int
}

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
    if length > highestLength {
      highestLength = length
    }
  }

  return highestLength
}

func main() {
  boundList := make([]Pair, 0)
  scanning := true

  for scanning {
    bounds := new(Pair)
    count, err := fmt.Scanln(&bounds.min, &bounds.max)

    if count == 0 {
      scanning = false
      continue // TODO(mmontgomery): fix pls
    } else if err != nil {
      panic(fmt.Sprintf("Error: ", err))
    }
    if count != 2 {
      panic("Exactly two arguments required.")
    }
    if bounds.max < bounds.min {
      panic("Max was less than min!")
    }

    boundList = append(boundList, *bounds)
  }

  for _, bounds := range boundList {
    highestLength := calculateLongestCycleLength(bounds.min, bounds.max)
    fmt.Printf("%v %v %v", bounds.min, bounds.max, highestLength)
  }
}
