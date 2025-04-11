package main

import (
	"3-struct/bins"
	"3-struct/storage"
	"fmt"
	"github.com/fatih/color"
)

func main() {
	// Получение пути, где хранится файл storage_bin
	filePath := getFilePath()

	// Инициализация Storage
	storageBin := storage.NewStorage(filePath)

	binList := TestData()

	// Сохранения Bins в файл
	err := storageBin.SaveBinList(binList)
	if err != nil {
		color.Red(err.Error())
		return
	}

	// Чтение
	binListJson, err := storageBin.ReadBinList()
	if err != nil {
		color.Red(err.Error())
		return
	}
	fmt.Println(binListJson)
}

func getFilePath() string {
	fmt.Print("Укажите путь до файла или его название, если он находится в корне проекта: ")
	var filePath string
	_, err := fmt.Scan(&filePath)
	if err != nil {
		color.Red(err.Error())
		return ""
	}
	return filePath
}

// TestData - функция для тестового добавления данных в JSON файл
func TestData() *bins.BinList {
	// Создание Bin
	bin1, err := bins.NewBin("1", "First", false)
	if err != nil {
		fmt.Println(err)
	}
	bin2, err := bins.NewBin("2", "Second", true)
	if err != nil {
		fmt.Println(err)
	}

	// Создание BinList и добавление в него Bins
	binList := bins.NewBinList()
	binList.Bins = append(binList.Bins, *bin1)
	binList.Bins = append(binList.Bins, *bin2)

	return binList
}
