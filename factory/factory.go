package factory

type Product interface {
	Price() int
	Name() string
}

type HotTub struct{}
type WashingMachine struct{}

func (h HotTub) Price() int {
	return 800
}

func (w WashingMachine) Price() int {
	return 100
}

func (h HotTub) Name() string {
	return "Hot Tub"
}

func (w WashingMachine) Name() string {
	return "Washing Machine"
}

func NewProduct(productType string) Product {
	switch productType {
	case "washing machine":
		return WashingMachine{}
	case "hot tub":
		return HotTub{}
	default:
		return nil
	}
}
