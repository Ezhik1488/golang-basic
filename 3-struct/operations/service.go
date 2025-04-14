package operations

import (
	api2 "3-struct/api"
	"3-struct/bins"
	"3-struct/file"
	"3-struct/storage"
	"flag"
	"github.com/fatih/color"
	"time"
)

type OperationsBins struct {
	Api       *api2.JsonBinAPI
	Storage   *storage.Storage
	flags     *Flags
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
		flags:     flags,
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

func (op *OperationsBins) CreateBin() error {
	dataFile, err := op.localFile.ReadFile()
	if err != nil {
		panic(err)
	}
	response, err := op.Api.Create(dataFile)
	if err != nil {
		color.Red(err.Error())
		return err
	}
	binID := response.MetaData.ID
	newBin, err := bins.NewBin(binID, op.flags.BinName, false)
	if err != nil {
		color.Red(err.Error())
		return err
	}

	op.binList.Bins = append(op.binList.Bins, *newBin)
	err = op.Storage.SaveBinList(op.binList)
	if err != nil {
		color.Red(err.Error())
		return err
	}
	return nil
}

func (op *OperationsBins) UpdateBin() string { return "" }

func (op *OperationsBins) DeleteBin() string { return "" }

func (op *OperationsBins) GetBin() string {

	return ""
}

func (op *OperationsBins) PrintBinsList() string { return "" }

func getFlags() *Flags {
	createBin := flag.Bool("create", true, "Добавить Bin")
	updateBin := flag.Bool("update", false, "Обновить Bin")
	deleteBin := flag.Bool("delete", false, "Удалить Bin")
	getBin := flag.Bool("get", false, "Получить Bin ")
	listBins := flag.Bool("list", false, "Получить список ID и Name из локального файла")
	dataFile := flag.String("file", "data.json", "Путь файла для загрузки/обновления")
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
