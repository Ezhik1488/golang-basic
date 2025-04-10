package bins

import (
	"errors"
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

func NewBin(id string, name string, private bool) (*Bin, error) {
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

func NewBinList() *BinList {
	return &BinList{}
}
