package main

import "fmt"

func main() {
	revenue := 1000
	expenses := 200
	taxRate := 25

	earningsBT := revenue - expenses
	earningsAT := earningsBT * (1 - (taxRate / 100))

	fmt.Print("earnings before tax: ")
	fmt.Println(earningsBT)
	fmt.Print("earnings after tax: ")
	fmt.Print(earningsAT)
}
