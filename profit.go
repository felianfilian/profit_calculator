package main

import (
	"fmt"
	"math"
)

func main() {
	revenue := 1000.0
	expenses := 200.0
	taxRate := 25.0

	fmt.Print("Earnings: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	earningsBT := getEarnings(revenue, expenses)
	earningsAT := earningsBT * (1 - (taxRate / 100))
	earningsBT = math.Round(earningsBT)

	fmt.Printf("\nearnings before tax: %.1f\n", earningsBT)
	//sfmt.Println(earningsBT)
	fmt.Print("earnings after tax: ")
	fmt.Print(earningsAT)
}

func getEarnings(rev float64, exp float64) float64 {
	return rev - exp
}
