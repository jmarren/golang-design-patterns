package main

import (
	"fmt"

	"github.com/jmarren/golang-design-patterns/abstract_factory"
	"github.com/jmarren/golang-design-patterns/builder"
	"github.com/jmarren/golang-design-patterns/prototype"
	"github.com/jmarren/golang-design-patterns/singleton"
)

func main() {
	fmt.Println("running!")
	abstractfactory.Run()
	builder.Run()
	prototype.Run()
	singleton.Run()
}
