package main

import (
	"fmt"
	"strings"
)

type RatioMap map[string]float64

func main() {
	printHelloMessage()
	firstCurrency, secondCurrency, amount := UserInput()
	result := CurrencyConverter(firstCurrency, secondCurrency, amount)
	fmt.Printf("Результат конвертации %s в %s: %.2f", firstCurrency, secondCurrency, result)
}

func printHelloMessage() {
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
			continue
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
			continue
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
	ratios := RatioMap{
		"usd_rub":    86.19,
		"usd_eur":    0.92,
		"eur_to_rub": 93.78,
	}

	directKey := source + "_" + target
	if ratio, exists := ratios[directKey]; exists {
		return ratio * amount
	}

	reverseKey := target + "_" + source
	if ratio, exists := ratios[reverseKey]; exists {
		return amount / ratio
	}

	return 0.00
}
