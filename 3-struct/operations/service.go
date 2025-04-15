package operations

import (
	api2 "3-struct/api"
	"3-struct/bins"
	"3-struct/file"
	"3-struct/storage"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"time"
)

type OperationsBins struct {
	Api       *api2.JsonBinAPI
	Storage   *storage.Storage
	Flags     *Flags
	localFile *file.LocalStorage
	binList   *bins.BinList
}

func NewOperationsBins(api *api2.JsonBinAPI, storage *storage.Storage) *OperationsBins {
	flags := getFlags()
	localFile := file.NewLocalStorage(flags.File, ".json")
	binList, err := storage.ReadBinList()
	if err != nil {
		binList = &bins.BinList{}
	}

	return &OperationsBins{
		Api:       api,
		Storage:   storage,
		Flags:     flags,
		localFile: localFile,
		binList:   binList,
	}
}

type Flags struct {
	Create  bool
	Update  bool
	Delete  bool
	Get     bool
	List    bool
	File    string
	BinID   string
	BinName string
}

var (
	ErrAddToJsonBin  = errors.New("ошибка при добавлении Bins в JSON BIN")
	ErrCreateBin     = errors.New("ошибка создания нового Bin")
	ErrSaveToStorage = errors.New("ошибка сохранения в хранилище")
	ErrReadFile      = errors.New("ошибка при чтении файла")
	ErrUpdate        = errors.New("ошибка при обновлении Bin")
	ErrDelete        = errors.New("ошибка при удалении Bin с JSON BIN")
	ErrGetBin        = errors.New("ошибка получения Bin с JSON BIN")
)

func (op *OperationsBins) CreateBin() error {
	dataFile, err := op.localFile.ReadFile()
	if err != nil {
		return ErrReadFile
	}
	response, err := op.Api.Create(dataFile)
	if err != nil {
		color.Red(err.Error())
		return ErrAddToJsonBin
	}
	binID := response.MetaData.ID
	newBin, err := bins.NewBin(binID, op.Flags.BinName, false)
	if err != nil {
		color.Red(err.Error())
		return ErrCreateBin
	}

	op.binList.Bins = append(op.binList.Bins, *newBin)
	err = op.Storage.SaveBinList(op.binList)
	if err != nil {
		color.Red(err.Error())
		return ErrSaveToStorage
	}
	return nil
}

func (op *OperationsBins) UpdateBin() error {
	dataFile, err := op.localFile.ReadFile()
	if err != nil {
		color.Red("Ошибка при чтении файла %s", op.Flags.File)
		return ErrReadFile
	}
	err = op.Api.Update(dataFile, op.Flags.BinID)
	if err != nil {
		color.Red(err.Error())
		return ErrUpdate
	}
	color.Green("Updated bin successfully: %s", op.Flags.BinID)
	return nil
}

func (op *OperationsBins) DeleteBin() error {
	err := op.Api.Delete(op.Flags.BinID)
	if err != nil {
		color.Red(err.Error())
		return ErrDelete
	}
	for i, bin := range op.binList.Bins {
		if bin.ID == op.Flags.BinID {
			op.binList.Bins = append(op.binList.Bins[:i], op.binList.Bins[i+1:]...)
		}
	}
	err = op.Storage.SaveBinList(op.binList)
	if err != nil {
		color.Red(err.Error())
		return ErrSaveToStorage
	}

	color.Green("Deleted bin successfully: %s", op.Flags.BinID)
	return nil
}

func (op *OperationsBins) GetBin() error {
	result, err := op.Api.Get(op.Flags.BinID)
	if err != nil {
		color.Red(err.Error())
		return ErrGetBin
	}
	jsonData, err := json.MarshalIndent(result, "", "\t")
	if err != nil {
		color.Red("Ошибка при машралинге %s", err.Error())
	}
	fmt.Println(string(jsonData))
	return nil
}

func (op *OperationsBins) PrintBinsList() error {
	for _, bin := range op.binList.Bins {
		jsonData, err := json.MarshalIndent(bin, "", " ")
		if err != nil {
			color.Red("Ошибка при машралинге %s", err.Error())
		}
		fmt.Println(string(jsonData))
	}
	return nil
}

func getFlags() *Flags {
	createBin := flag.Bool("create", false, "Добавить Bin")
	updateBin := flag.Bool("update", false, "Обновить Bin")
	deleteBin := flag.Bool("delete", false, "Удалить Bin")
	getBin := flag.Bool("get", false, "Получить Bin ")
	listBins := flag.Bool("list", false, "Получить список ID и Name из локального файла")
	dataFile := flag.String("file", "", "Путь файла для загрузки/обновления")
	binID := flag.String("id", "", "BinID для получения/обновления/удаления")
	binName := flag.String("name", time.Now().String(), "Имя Bin в локальном файле")

	flag.Parse()
	return &Flags{
		Create:  *createBin,
		Update:  *updateBin,
		Delete:  *deleteBin,
		Get:     *getBin,
		List:    *listBins,
		File:    *dataFile,
		BinID:   *binID,
		BinName: *binName,
	}
}

func (op *OperationsBins) ConvertFlagToMap() map[string]bool {
	return map[string]bool{
		"create": op.Flags.Create,
		"update": op.Flags.Update,
		"delete": op.Flags.Delete,
		"get":    op.Flags.Get,
		"list":   op.Flags.List,
	}
}
