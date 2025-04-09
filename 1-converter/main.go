package main

import "fmt"

func main() {
	const (
		USDToEUR = 0.92
		USDToRUB = 86.19
	)
	var (
		firstCurrency, secondCurrency string
		amount                        float64
	)
	EURToRUB := USDToRUB / USDToEUR

	fmt.Println("___Конвертер валют___")
	firstCurrency, secondCurrency, amount = UserInput()
	fmt.Println(firstCurrency, secondCurrency, amount, EURToRUB)

}

func UserInput() (string, string, float64) {
	var firstCurrency, secondCurrency string
	var amount float64
	fmt.Print("Введите исходную валюту(EUR, USD, RUB): ")
	fmt.Scan(&firstCurrency)
	fmt.Print("Введите целевую валюту(EUR, USD, RUB): ")
	fmt.Scan(&secondCurrency)
	fmt.Print("Введите сумму для конвертации: ")
	fmt.Scan(&amount)

	return firstCurrency, secondCurrency, amount

}

func CurrencyConverter(source, target string, amount float64) float64 {
	return 0.00
}
