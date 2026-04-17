package main

import "fmt"

func calculateProfits() {
	revenue := getInput("Enter revenue: ")
	expenses := getInput("Enter expenses: ")
	taxRate := getInput("Enter taxRate: ")

	earningsBeforeTax, earningsAfterTax, taxRatio := _calcProfitValues(revenue, expenses, taxRate)

	fmt.Printf("Earning before tax: %.2f\n", earningsBeforeTax)
	fmt.Printf("Earning after tax: %.2f\n", earningsAfterTax)
	fmt.Printf("tax ratio: %.2f\n", taxRatio)
}

func getInput(inputText string) (variable float64) {
	fmt.Print(inputText)
	fmt.Scan(&variable)
	return
}

func _calcProfitValues(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100.0)
	taxRatio := earningsBeforeTax / earningsAfterTax

	return earningsBeforeTax, earningsAfterTax, taxRatio
}
