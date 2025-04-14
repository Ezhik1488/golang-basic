package operations

import (
	api2 "3-struct/api"
	"3-struct/storage"
	"flag"
)

type OperationsBins struct {
	Api     *api2.JsonBinAPI
	Storage *storage.Storage
	flags   *Flags
}

func NewOperationsBins(api *api2.JsonBinAPI, storage *storage.Storage) *OperationsBins {
	return &OperationsBins{Api: api, Storage: storage, flags: getFlags()}
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

func (op *OperationsBins) CreateBin() string { return "" }

func (op *OperationsBins) UpdateBin() string { return "" }

func (op *OperationsBins) DeleteBin() string { return "" }

func (op *OperationsBins) GetBin() string {

	return ""
}

func (op *OperationsBins) PrintBinsList() string { return "" }

func getFlags() *Flags {
	createBin := flag.Bool("create", false, "Добавить Bin")
	updateBin := flag.Bool("update", false, "Обновить Bin")
	deleteBin := flag.Bool("delete", false, "Удалить Bin")
	getBin := flag.Bool("get", false, "Получить Bin ")
	listBins := flag.Bool("list", false, "Получить список ID и Name из локального файла")
	dataFile := flag.String("file", "", "Путь файла для загрузки/обновления")
	binID := flag.String("id", "", "BinID для получения/обновления/удаления")
	binName := flag.String("name", "", "Имя Bin в локальном файле")

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
