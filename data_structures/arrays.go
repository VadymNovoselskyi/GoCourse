package main

import (
	"fmt"
	"slices"
)

func ArrayStuff() {
	prices := []int{1, 2, 3, 4, 5, 6}
	fmt.Println((prices))

	featuredPrices := prices[:3]
	fmt.Println((featuredPrices))

	// append but override the original
	featuredPrices = append(featuredPrices, 99)
	fmt.Println((prices))
	fmt.Println((featuredPrices))

	// shift but also grab the elements outside of slice's reach
	featuredPrices = featuredPrices[1:cap(featuredPrices)]
	fmt.Println((featuredPrices))

	slices.Sort(featuredPrices)
	fmt.Println((featuredPrices))

	for i, price := range featuredPrices {
		fmt.Printf("Element with idx %d is %d \n", i, price)
	}
}
