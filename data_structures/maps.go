package main

import "fmt"

func MapStuff() {
	someMap := map[string]int{"1": 9, "22": 99, "333": 999}
	fmt.Println(someMap)

	for key, value := range someMap {
		fmt.Printf("key: %v; value: %v \n", key, value)
	}
}
