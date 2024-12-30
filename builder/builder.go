package builder

import (
	"fmt"
)

func Run() {
	fmt.Println("----------- Builder Pattern ------------")
	normalBuilder := getBuilder("normal")
	OnSaleBuilder := getBuilder("onsale")
	warrantyBuilder := getBuilder("warranty")
	bothBuilder := getBuilder("both")

	director := newDirector(normalBuilder)
	normalProduct := director.buildProduct()
	fmt.Printf("normalProduct: %v\n", normalProduct)

	director = newDirector(OnSaleBuilder)
	onSaleProduct := director.buildProduct()
	fmt.Printf("normalProduct: %v\n", onSaleProduct)

	director = newDirector(warrantyBuilder)
	warrantyProduct := director.buildProduct()
	fmt.Printf("warrantyProduct: %v\n", warrantyProduct)

	director = newDirector(bothBuilder)
	bothProduct := director.buildProduct()
	fmt.Printf("bothProduct: %v\n", bothProduct)
}

type IBuilder interface {
	setCategory()
	setPrice()
	setOtherInfo()
	getProduct() Product
}

type Product struct {
	category       string
	price          int
	salePercentOff int
	warrantyLength int
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return newNormalBuilder()
	}
	if builderType == "onsale" {
		return newOnSaleBuilder()
	}
	if builderType == "warranty" {
		return newWarrantyBuilder()
	}

	if builderType == "both" {
		return newBothBuilder()
	}
	return nil
}
