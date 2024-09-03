package main

import (
	"fmt"
	"math"
)

func main() {
	revenue := userInput("Earnings")
	expenses := userInput("Expenses: ")
	taxRate := 25.0

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

func userInput(text string) float64 {
	var userInputNumber float64
	fmt.Print(text)
	fmt.Scan(&userInputNumber)
	return userInputNumber
}
