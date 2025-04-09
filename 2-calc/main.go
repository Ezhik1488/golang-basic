package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Калькулятор")
	operation := choiceOperation()
	numbers := inputNumbers()

	result := Calculate(operation, numbers)
	fmt.Printf("Результат: %.2f", result)

}

func choiceOperation() string {
	var operation string
	for {
		fmt.Print("Введите тип операции(Доступные варианты: AVG, SUM, MED): ")
		_, err := fmt.Scan(&operation)
		if err != nil {
			continue
		}
		operation = strings.ToLower(operation)

		switch operation {
		case "avg":
			return operation
		case "sum":
			return operation
		case "med":
			return operation
		default:
			fmt.Println("Неизвестный тип операции")
		}
	}

}

func inputNumbers() []float64 {
	var numbers []float64
	var rawNumbers string

	for {
		fmt.Print("Введите числа через запятую: ")
		_, err := fmt.Scan(&rawNumbers)
		if err != nil {
			continue
		}

		tempSlice := strings.Split(rawNumbers, ",")
		for _, number := range tempSlice {
			floatNumber, err_ := strconv.ParseFloat(number, 64)
			if err_ != nil {
				fmt.Println("Ошибка при конвертации строки в float64. Повторите ввод. Err:" + err_.Error())
				continue
			}
			numbers = append(numbers, floatNumber)
		}
		return numbers
	}

}

func Calculate(operation string, numbers []float64) float64 {
	switch operation {
	case "avg":
		return Average(numbers)
	case "sum":
		return Sum(numbers)
	case "med":
		return Median(numbers)
	default:
		return 0.00
	}
}

func Average(numbers []float64) float64 {
	length := len(numbers)
	return Sum(numbers) / float64(length)
}

func Sum(numbers []float64) float64 {
	var result float64

	for _, number := range numbers {
		result += number
	}
	return result
}

func Median(numbers []float64) float64 {
	length := len(numbers)

	sort.Float64s(numbers)

	switch {
	case length == 1:
		return numbers[0]
	case length%2 == 0:
		return (numbers[length/2] + numbers[(length/2)-1]) / 2.0
	case length%2 != 0:
		return numbers[length/2]
	default:
		return 0.00
	}
}
