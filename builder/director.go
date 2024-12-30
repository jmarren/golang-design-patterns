package builder

type Director struct {
	builder IBuilder
}

func newDirector(b IBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) setBuilder(b IBuilder) {
	d.builder = b
}

func (d *Director) buildProduct() Product {
	d.builder.setCategory()
	d.builder.setPrice()
	d.builder.setOtherInfo()
	return d.builder.getProduct()
}
