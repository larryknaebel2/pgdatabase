package client

import (
	"examples/pgdatabase/internal/clnt"
	"fmt"

	"github.com/jinzhu/copier"
)

type Supply struct {
	Id           int64
	Name         string
	Description  string
	Manufacturer string
	Color        string
	Inventory    int64
}

type S struct {
	Id   int64
	Name string
}

func GetAllSupplies() ([]Supply, error) {
	s, err := clnt.GetAllSupplies()
	if err != nil {
		return nil, fmt.Errorf("%v", err)
	}
	var supplies []Supply
	copier.Copy(&supplies, &s)

	var partOfSupply []S
	copier.Copy(&partOfSupply, &s) // just checking if can copy just part of struct - YES
	//copy(supplies,s)
	//y := *(*Y)(unsafe.Pointer(&z))
	// y, err := deepcopy.Anything(s)
	// if err != nil {
	// 	fmt.Print(fmt.Errorf("Deepcopy Error: %v", err))
	// }
	// fmt.Printf("%v", y)
	fmt.Printf("Part: %v\n", partOfSupply)
	return supplies, nil
}

func GetSupplyById(id int64) (Supply, error) {
	supply, err := clnt.GetSupplyById(id)
	if err != nil {
		return Supply{}, fmt.Errorf("%v", err)
	}
	var asupply Supply
	copier.Copy(&asupply, &supply)
	//y := *(*Supplies)(unsafe.Pointer(&supply))
	return asupply, nil
}
