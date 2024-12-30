package abstractfactory

import (
	"fmt"
)

func Run() {
	fmt.Println("------------ Abstract Factory ---------------")
	chicagoFactory, _ := GetStoreFactory("Chicago")
	newYorkFactory, _ := GetStoreFactory("New York")

	chicagoEmployee := chicagoFactory.makeEmployee()
	chicagoEmployee.setSalary(100)
	newYorkProduct := newYorkFactory.makeProduct()
	newYorkProduct.setInventory(300)
	printEmployee(chicagoEmployee)
	printProduct(newYorkProduct)

}

func printEmployee(e IEmployee) {
	fmt.Printf("employee works at %s and received a salary of %d\n", e.getLocation(), e.getSalary())
}

func printProduct(p IProduct) {
	fmt.Printf("the product is at %s and there are %d at this location\n", p.getLocation(), p.getInventory())
}
