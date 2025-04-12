package storage

import (
	"3-struct/bins"
	"encoding/json"
)

type DB interface {
	ReadFile() ([]byte, error)
	WriteFile(data []byte) error
}

type Storage struct {
	db DB
}

func NewStorage(fileStorage DB) *Storage {
	return &Storage{fileStorage}
}

func (storage *Storage) ReadBinList() (*bins.BinList, error) {
	var binList bins.BinList

	jsonFile, err := storage.db.ReadFile()
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
	if err != nil {
		return err
	}
	err = storage.db.WriteFile(data)
	if err != nil {
		return err
	}

	return nil
}
