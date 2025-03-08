package main

import "fmt"

func main() {
	var revenue float32
	var expenses float32
	var taxRate float32

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Printf("Revenue: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)
}
