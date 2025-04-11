package storage

import (
	"3-struct/bins"
	"3-struct/file"
	"encoding/json"
	"github.com/fatih/color"
)

type Storage struct {
	FilePath string
}

func NewStorage(filePath string) *Storage {
	return &Storage{FilePath: filePath}
}

func (storage *Storage) ReadBinList() (*bins.BinList, error) {
	var binList bins.BinList

	jsonFile, err := file.ReadFileWithExt(storage.FilePath, ".json")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(jsonFile, &binList)
	if err != nil {
		return nil, err
	}
	return &binList, nil
}

func (storage *Storage) SaveBinList(bins *bins.BinList) error {
	data, err := bins.ToBytes()
	err = file.WriteFile(storage.FilePath, data)
	if err != nil {
		return err
	}
	color.Green("Saved bins list to file: %s", storage.FilePath)
	return nil
}
