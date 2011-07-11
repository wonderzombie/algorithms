package main

import (
	"fmt"
	"sort"
)

type ExpenseReport struct {
	expenses []float64
}

func getInput(a ...interface{}) {
	_, err := fmt.Scanln(a...)

	if err != nil {
		panic(err)
	}
}

func readExpenses() *ExpenseReport {
	var students int
	// count, err := fmt.Scanln(&students)
	getInput(&students)

	report := new(ExpenseReport)
	for i := 0; i < students; i++ {
		var expense float64
		getInput(&expense)
		report.expenses = append(report.expenses, expense)
	}

	return report
}

// See http://golang.org/src/cmd/godoc/index.go?h=unique#L957.
func uniq(list []float64) []float64 {
	var last float64
	i := 0
	for _, x := range list {
		if i == 0 || x != last {
			last = x
			list[i] = x
			i++
		}
	}
	return list[0:i]
}

func normalizeExpenses(e []float64) []float64 {
	sort.SortFloat64s(e)
	e = uniq(e)
	return e
}

func isLenEven(list []float64) bool {
	return len(list) % 2 == 0
}

func min(a float64, b float64) float64 {
	if b < a {
		return b
	}
	return a
}

func calculateExpenseCorrection(report *ExpenseReport) float64 {
	e := report.expenses
	e = normalizeExpenses(e)

	var (
		diff1 float64
		diff2 float64
		correction float64
	)
	var m int = (len(e) - 1) / 2

	if isLenEven(e) {
	  diff1 = e[m + 1] - e[m]
		diff2 = diff1
	} else {
		diff1 = e[m] - e[m - 1]
		diff2 = e[m + 1] - e[m]
	}
	correction = min(diff1, diff2)

	return correction
}

func main() {
	report := readExpenses()
	fmt.Printf("Expense report: %v\n", report)
	correction := calculateExpenseCorrection(report)
	fmt.Println(correction)
}
