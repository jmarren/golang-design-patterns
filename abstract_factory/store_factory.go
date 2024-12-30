package abstractfactory

import (
	"fmt"
)

type IStore interface {
	makeEmployee() IEmployee
	makeProduct() IProduct
}

type ChicagoStore struct{}

type NewYorkStore struct{}

func (c *ChicagoStore) makeEmployee() IEmployee {
	return &Employee{
		location: "Chicago",
	}
}

func (c *ChicagoStore) makeProduct() IProduct {
	return &Product{
		location: "Chicago",
	}
}

func (n *NewYorkStore) makeProduct() IProduct {
	return &Product{
		location: "New York",
	}
}

func (n *NewYorkStore) makeEmployee() IEmployee {
	return &Employee{
		location: "New York",
	}
}

func GetStoreFactory(location string) (IStore, error) {
	if location == "Chicago" {
		return &ChicagoStore{}, nil
	}

	if location == "New York" {
		return &NewYorkStore{}, nil
	}
	return nil, fmt.Errorf("store does not exist for location %s", location)
}
