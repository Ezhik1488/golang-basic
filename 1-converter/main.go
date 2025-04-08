package main

import "fmt"

func main() {
	const (
		USDToEUR = 0.92
		USDToRUB = 86.19
	)

	EURToRUB := USDToRUB / USDToEUR

	fmt.Println(EURToRUB)
}
