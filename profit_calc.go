package main

import (
	"errors"
	"fmt"
	"os"
)

func calculateProfits() {
	revenue, err := getPositiveInput("Enter revenue: ")
	if err != nil {
		fmt.Println("Invalid input for revenue:", err)
		return
	}

	expenses, err := getPositiveInput("Enter expenses: ")
	if err != nil {
		fmt.Println("Invalid input for expenses:", err)
		return
	}

	taxRate, err := getPositiveInput("Enter taxRate: ")
	if err != nil {
		fmt.Println("Invalid input for tax rate:", err)
		return
	}

	earningsBeforeTax, earningsAfterTax, taxRatio := calcProfitValues(revenue, expenses, taxRate)
	earningsBeforeTaxStr := fmt.Sprintf("Earning before tax: %.2f\n", earningsBeforeTax)
	earningsAfterTaxStr := fmt.Sprintf("Earning after tax: %.2f\n", earningsAfterTax)
	taxRatioStr := fmt.Sprintf("tax ratio: %.2f\n", taxRatio)

	fmt.Print(earningsBeforeTaxStr)
	fmt.Print(earningsAfterTaxStr)
	fmt.Print(taxRatioStr)

	os.WriteFile("profit-calculations.txt", []byte(earningsBeforeTaxStr+earningsAfterTaxStr+taxRatioStr), 0644)
}

func getPositiveInput(inputText string) (float64, error) {
	var inputValue float64
	fmt.Print(inputText)
	fmt.Scan(&inputValue)

	if inputValue <= 0 {
		return 0, errors.New("Input must be a positive number")
	}

	return inputValue, nil
}

func calcProfitValues(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningsBeforeTax := revenue - expenses
	earningsAfterTax := earningsBeforeTax * (1 - taxRate/100.0)
	taxRatio := earningsBeforeTax / earningsAfterTax

	return earningsBeforeTax, earningsAfterTax, taxRatio
}
