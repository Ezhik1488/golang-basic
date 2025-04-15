package file

import (
	"errors"
	"github.com/fatih/color"
	"os"
	"path"
)

type LocalStorage struct {
	Path          string
	FileExtension string
}

func NewLocalStorage(path, fileExtension string) *LocalStorage {
	return &LocalStorage{path, fileExtension}
}

func (storage *LocalStorage) ReadFile() ([]byte, error) {
	// Проверить расширение файла
	if path.Ext(storage.Path) != storage.FileExtension {
		return nil, errors.New("file must have " + storage.FileExtension + " extension")
	}

	file, err := os.ReadFile(storage.Path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (storage *LocalStorage) WriteFile(data []byte) error {
	file, err := os.Create(storage.Path)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			color.Red("Error closing file: " + err.Error())
		}
	}(file)

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	color.Green("Saved bins list to file: %s", storage.Path)
	return nil
}
