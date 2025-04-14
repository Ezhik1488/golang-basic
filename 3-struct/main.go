package main

import (
	api2 "3-struct/api"
	"3-struct/bins"
	config2 "3-struct/config"
	"3-struct/file"
	"3-struct/storage"
	"fmt"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	// Получение переменных окружения
	err := godotenv.Load()
	if err != nil {
		color.Red("Error loading .env file")
	}
	color.Green("Данные из .env успешно загружены!")

	// Инициализация конфига
	config := config2.NewConfig()
	color.Magenta(config.ApiKey)

	// Получение пути, где хранится файл storage_bin
	filePath := getFilePath()

	// Инициализация ImplStorage
	localDB := file.NewLocalStorage(filePath, ".json")
	storageBin := storage.NewStorage(localDB)

	binList := TestData()

	// Инициализация JsonBinAPI
	api := api2.NewJsonBinAPI(config)
	result, err := api.Get("67fd00008a456b7966891650")
	if err != nil {
		fmt.Println("Ooops... some error")
	}
	fmt.Println(result)
	// TODO: Реализация взаимодействия с API JsonBin
	//  1. Получение необходимых флагов, переданных пользователем
	// 	2. Чтение необходимого файла по пути, переданным во флаге --file
	//  3. Выполнение соответствующей операции в зависимости от флага --create, --update, --delete
	// 	   Можно реализовать мапу map[string]func() для хранения функций и map[string]bool для разрешенных операций
	//  4. При create надо сохранять id и name в локальный файл

	// Сохранения Bins в файл
	err = storageBin.SaveBinList(binList)
	if err != nil {
		color.Red(err.Error())
		return
	}

	//// Чтение
	//binListJson, err := storageBin.ReadBinList()
	//if err != nil {
	//	color.Red(err.Error())
	//	return
	//}
	//fmt.Println(binListJson)
}

func getFilePath() string {
	fmt.Print("Укажите путь до файла в котором хранятся данные о Bins: ")
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
