package file

import (
	"errors"
	"os"
	"path"
)

func ReadFileWithExt(filePath string, ext string) ([]byte, error) {
	// Проверить расширение файла
	if path.Ext(filePath) != ext {
		return nil, errors.New("file must have " + ext + " extension")
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func WriteFile(filePath string, data []byte) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}
