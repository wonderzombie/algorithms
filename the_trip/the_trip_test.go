package main

import (
	"sort"
	"testing"
)

func TestMin(t *testing.T) {
	var a, b float64 = 2.0, 1.0

	expected := b
	actual := min(a, b)
	if expected != actual {
		t.Errorf("Incorrect minimum returned. Expected: %v, actual: %v",
				expected, actual);
	}
}

func TestUniq(t *testing.T) {
	var list []float64 = []float64{4.0, 1.0, 2.0, 3.0, 2.00, 2.0, 5.0}
	uniqued := uniq(list)
	if len(uniqued) == len(list) {
		t.Fatalf("Expected uniquified list to be shorter than original list.")
	}
}

func TestNormalizeExpenses(t *testing.T) {
	var list []float64 = []float64{4.0, 1.0, 2.0, 3.0, 2.00, 2.0, 5.0}
	normalized := normalizeExpenses(list)

	if !sort.Float64sAreSorted(normalized) {
		t.Fatalf("Expected normalized list to be sorted.")
	}

	var count float64 = 0
	for _, x := range normalized {
		if x == 2.0 {
			count++
		}
	}

	if count > 1 {
		t.Fatalf("Expected one instance of 2.0. Actual: %v", count)
	}
}

func TestCalculateExpenseCorrection(t *testing.T) {
	report := new(ExpenseReport)
	report.expenses = []float64 {10.00, 20.00, 30.00}
	expected := 10.00

	actual := calculateExpenseCorrection(report)
	if expected != actual {
		t.Fatalf("Expected correction to be %v, was %v.", expected, actual)
	}

	report.expenses = []float64 {3.00, 3.01, 15.00, 15.01}
	expected = 11.99

	actual = calculateExpenseCorrection(report)
	if expected != actual {
		t.Fatalf("Expected correction to be %v, was %v.", expected, actual)
	}
}
