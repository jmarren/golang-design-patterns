package builder

type OnSaleBuilder struct {
	NormalBuilder
	salePercentOff int
}

func newOnSaleBuilder() *OnSaleBuilder {
	return &OnSaleBuilder{}
}

func (b *OnSaleBuilder) setSalePercentOff() {
	b.salePercentOff = 10
}

func (b *OnSaleBuilder) setOtherInfo() {
	b.setSalePercentOff()
}

func (b *OnSaleBuilder) getSalePercentOff() int {
	return b.salePercentOff
}

func (b *OnSaleBuilder) getProduct() Product {
	return Product{
		price:          b.price,
		category:       b.category,
		salePercentOff: b.salePercentOff,
	}
}
