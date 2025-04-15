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

type Operation map[string]func() error

func main() {
	// Инициализация необходимых компонентов
	binOperation := initApp()

	allowOperation := Operation{
		"create": binOperation.CreateBin,
		"delete": binOperation.DeleteBin,
		"update": binOperation.UpdateBin,
		"get":    binOperation.GetBin,
		"list":   binOperation.PrintBinsList,
	}

	flagValue := binOperation.ConvertFlagToMap()

	for key, value := range flagValue {
		if value {
			action := allowOperation[key]
			err := action()
			if err != nil {
				return
			}
		}
	}
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
