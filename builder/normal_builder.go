package builder

type NormalBuilder struct {
	price    int
	category string
}

func newNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setPrice() {
	b.price = 100
}

func (b *NormalBuilder) setCategory() {
	b.category = "miscellaneous"
}

func (b *NormalBuilder) getPrice() int {
	return b.price
}

func (b *NormalBuilder) getCategory() string {
	return b.category
}

func (b *NormalBuilder) setOtherInfo() {
	return
}

func (b *NormalBuilder) getProduct() Product {
	return Product{
		price:          b.price,
		category:       b.category,
		salePercentOff: 0,
		warrantyLength: 0,
	}
}
