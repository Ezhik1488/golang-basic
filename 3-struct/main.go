package main

import (
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	ID        string
	Name      string
	Private   bool
	CreatedAt time.Time
}

type BinList struct {
	Bins []Bin
}

func newBin(id string, name string, private bool) (*Bin, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	if name == "" {
		return nil, errors.New("name is required")
	}

	return &Bin{
		ID:        id,
		Name:      name,
		Private:   private,
		CreatedAt: time.Now(),
	}, nil
}

func newBinList() *BinList {
	return &BinList{}
}

func main() {
	bin1, err := newBin("1", "First", false)
	if err != nil {
		fmt.Println(err)
	}
	binList := newBinList()
	binList.Bins = append(binList.Bins, *bin1)

	fmt.Println(bin1)
	fmt.Println(binList)

}
