package abstractfactory

type Product struct {
	location  string
	inventory int
}

type IProduct interface {
	getLocation() string
	getInventory() int
	setLocation(string)
	setInventory(int)
}

func (p *Product) setLocation(location string) {
	p.location = location
}

func (p *Product) setInventory(inventory int) {
	p.inventory = inventory
}

func (p *Product) getLocation() string {
	return p.location
}

func (p *Product) getInventory() int {
	return p.inventory
}
