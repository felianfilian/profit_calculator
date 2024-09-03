package main

import (
	"fmt"
	"math"
)

func main() {
	revenue := userInput("Earnings: ")
	expenses := userInput("Expenses: ")
	taxRate := 25.0

	earningsBT, earningsAT := getEarnings(revenue, expenses, taxRate)

	fmt.Printf("\nearnings before tax: %.2f Eur", earningsBT)
	fmt.Printf("\nearnings after tax: %.2f Eur", earningsAT)
}

func getEarnings(rev float64, exp float64, tax float64) (float64, float64) {
	ebt := rev - exp
	eat := ebt * (1 - (tax / 100))
	ebt = math.Round(ebt)
	eat = math.Round(eat)
	return ebt, eat
}

func userInput(text string) float64 {
	var userInputNumber float64
	fmt.Print(text)
	fmt.Scan(&userInputNumber)
	return userInputNumber
}
