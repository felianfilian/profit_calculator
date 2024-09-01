package main

import "fmt"

func main() {
	revenue := 1000.0
	expenses := 200.0
	taxRate := 25.0

	fmt.Print("Earnings: ")
	fmt.Scan(&revenue)
	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	earningsBT := revenue - expenses
	earningsAT := earningsBT * (1 - (taxRate / 100))

	fmt.Print("earnings before tax: ")
	fmt.Println(earningsBT)
	fmt.Print("earnings after tax: ")
	fmt.Print(earningsAT)
}
