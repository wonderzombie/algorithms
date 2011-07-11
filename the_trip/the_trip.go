package main

import (
	"fmt"
)

type ExpenseReport struct {
	expenses []float32
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
		var expense float32
		getInput(&expense)
		report.expenses = append(report.expenses, expense)
	}

	return report
}

func main() {
	report := readExpenses()
	fmt.Printf("Expense report: %v\n", report)
}
