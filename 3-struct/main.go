package main

import (
	api2 "3-struct/api"
	config2 "3-struct/config"
	"3-struct/file"
	"3-struct/operations"
	"3-struct/storage"
	"fmt"
	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	// Инициализация необходимых компонентов
	binOperation := initApp()

	err := binOperation.CreateBin()
	if err != nil {
		color.Red(err.Error())
	} else {
		color.Green("BinOperation Created Successfully")
	}

	// TODO: Реализация взаимодействия с API JsonBin
	//  1. Получение необходимых флагов, переданных пользователем
	// 	2. Чтение необходимого файла по пути, переданным во флаге --file
	//  3. Выполнение соответствующей операции в зависимости от флага --create, --update, --delete
	// 	   Можно реализовать мапу map[string]func() для хранения функций и map[string]bool для разрешенных операций
	//  4. При create надо сохранять id и name в локальный файл

	// Сохранения Bins в файл

	//// Чтение
	//binListJson, err := storageBin.ReadBinList()
	//if err != nil {
	//	color.Red(err.Error())
	//	return
	//}
	//fmt.Println(binListJson)
}

func initApp() *operations.OperationsBins {
	// Получение переменных окружения
	err := godotenv.Load()
	if err != nil {
		color.Red("Error loading .env file")
	}
	color.Green("Данные из .env успешно загружены!")

	// Инициализация конфига
	config := config2.NewConfig()
	color.Magenta(config.ApiKey)

	// Инициализация ImplStorage
	localDB := file.NewLocalStorage(config.LocalStoragePath, ".json")
	storageBin := storage.NewStorage(localDB)

	// Инициализация JsonBinAPI
	api := api2.NewJsonBinAPI(config)

	// Инициализации операций над Bins
	binOperation := operations.NewOperationsBins(api, storageBin)

	return binOperation
}

// getFilePath - deprecated
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
