package builder

type WarrantyBuilder struct {
	NormalBuilder
	warrantyLength int
}

func newWarrantyBuilder() *WarrantyBuilder {
	return &WarrantyBuilder{}
}

func (b *WarrantyBuilder) setWarranty() {
	b.warrantyLength = 10
}

func (b *WarrantyBuilder) setOtherInfo() {
	b.setWarranty()
}

func (b *WarrantyBuilder) getWarrantyLength() int {
	return b.warrantyLength
}

func (b *WarrantyBuilder) getProduct() Product {
	return Product{
		price:          b.price,
		category:       b.category,
		salePercentOff: 0,
		warrantyLength: b.warrantyLength,
	}
}
