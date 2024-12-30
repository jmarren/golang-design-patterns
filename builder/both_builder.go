package builder

type BothBuilder struct {
	NormalBuilder
	WarrantyBuilder
	OnSaleBuilder
}

func newBothBuilder() *BothBuilder {
	return &BothBuilder{}
}

func (b *BothBuilder) setOtherInfo() {
	b.setSalePercentOff()
	b.setWarranty()
}

func (b *BothBuilder) getWarrantyLength() int {
	return b.warrantyLength
}

func (b *BothBuilder) getProduct() Product {
	return Product{
		price:          b.NormalBuilder.price,
		category:       b.NormalBuilder.category,
		salePercentOff: b.salePercentOff,
		warrantyLength: b.warrantyLength,
	}
}
