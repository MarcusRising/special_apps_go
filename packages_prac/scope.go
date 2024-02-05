package main

import (
	"fmt"
	planet "projects/go/packages_prac/components"
)

var f = "Hello"

func main() {

	x := 5
	y := x + 4
	var result = y

	fmt.Println(result)
	fmt.Println(f)

	var p1 = person{"Dumbass"}
	p1.sayHello()

	planet.Facts() // gets this from the package folder
}

type person struct {
	first string
}

func (p person) sayHello() {
	fmt.Println("Hi, my name is", p.first)
}
