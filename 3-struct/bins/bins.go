package bins

import (
	"encoding/json"
	"errors"
	"time"
)

type Bin struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
}

type BinList struct {
	Bins []Bin `json:"bins"`
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

func (binList *BinList) ToBytes() ([]byte, error) {
	data, err := json.MarshalIndent(binList, "", "\t")
	return data, err
}
