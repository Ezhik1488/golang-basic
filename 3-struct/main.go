package main

import (
	"3-struct/bins"
	"fmt"
)

func main() {
	bin1, err := bins.NewBin("1", "First", false)
	if err != nil {
		fmt.Println(err)
	}
	binList := bins.NewBinList()
	binList.Bins = append(binList.Bins, *bin1)

	fmt.Println(bin1)
	fmt.Println(binList)

}
