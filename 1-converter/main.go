package main

import (
	"fmt"
	"strings"
)

const (
	USDToEUR = 0.92
	USDToRUB = 86.19
	EURToRUB = 93.78
)

func main() {
	helloMessage()
	firstCurrency, secondCurrency, amount := UserInput()
	result := CurrencyConverter(firstCurrency, secondCurrency, amount)
	fmt.Printf("Результат конвертации %s в %s: %.2f", firstCurrency, secondCurrency, result)
}

func helloMessage() {
	fmt.Println("\tКонвертер валют")
	fmt.Println("=====================")
	fmt.Println("Доступные валюты для конвертации: USD, EUR, RUB")
	fmt.Print("=====================\n\n")
}

func UserInput() (string, string, float64) {
	firstCurrency := inputStringValue("Введите исходную валюту: ", "")
	secondCurrency := inputStringValue("Введите целевую валюту: ", firstCurrency)
	amount := inputFloatValue("Введите сумму для конвертации: ")

	return firstCurrency, secondCurrency, amount

}

func inputFloatValue(msg string) float64 {
	var result float64

	for result <= 0 {
		fmt.Printf("%s", msg)
		_, err := fmt.Scan(&result)
		if err != nil {
			fmt.Println("Ошибка при получении ввода:" + err.Error())
			return 0.00
		}
		if result <= 0 {
			fmt.Println("[ОШИБКА] Сумма для конвертации должна быть > 0")
		}
	}
	return result
}

func inputStringValue(msg, secondCurr string) string {
	var value string
	for {
		fmt.Printf("%s", msg)
		_, err := fmt.Scan(&value)
		if err != nil {
			fmt.Println("Ошибка при получении ввода:" + err.Error())
			return ""
		}
		value = strings.ToLower(value)

		if secondCurr == value {
			fmt.Println("[ОШИБКА] Выберите разные валюты для конвертации!")
			continue
		}

		switch value {
		case "eur":
			return value
		case "usd":
			return value
		case "rub":
			return value
		default:
			fmt.Println("[ОШИБКА] Данная валюта недоступна для конвертации. Доступные валюты: USD, EUR, RUB")
		}
	}
}

func CurrencyConverter(source, target string, amount float64) float64 {
	switch {
	case source == "usd" && target == "eur":
		return amount * USDToEUR
	case source == "usd" && target == "rub":
		return amount * USDToRUB
	case source == "eur" && target == "rub":
		return amount * EURToRUB
	case source == "eur" && target == "usd":
		return amount / USDToEUR
	case source == "rub" && target == "usd":
		return amount / USDToRUB
	case source == "rub" && target == "eur":
		return amount / EURToRUB
	default:
		return 0.00
	}
}
