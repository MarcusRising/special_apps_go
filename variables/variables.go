package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var bigNum = 99

func variables() {
	a := 42 // short declaration operator
	fmt.Println(a)

	var g int = 89
	fmt.Println(g)
	g = 43
	fmt.Println(g)

	b, c, d, _, f := 0, 1, 2, 3, "Happy"
	fmt.Println(b, c, d, f)
}

func hexDex() {

	adam := 40
	fmt.Printf("40 as Binary: %#b\n", adam)
	fmt.Printf("40 as octal : %#o\n", adam)
	fmt.Printf("40 as .. : %T\n", adam)
	//fmt.Printf("40 as octal : %p\n", adam)
	fmt.Printf("40 as Hex: %X\n", adam)
	//print in hex and binary as %x %#b
	a, b, c, d, e, f := 1, 2, 3, 4, 5, 6
	//fmt.Printf("A: %x , %#b. \nB: %x , %#b. \nC: %x , %#b. \nD: %x , %#b. \nE: %x , %#b. \nF: %x , %#b. \n", a, a, b, b, c, c, d, d, e, e, f, f)
	fmt.Printf("%v \t %b \t %#x\n", a, a, a)
	fmt.Printf("%v \t %b \t %#x\n", b, b, b)
	fmt.Printf("%v \t %b \t %#x\n", c, c, c)
	fmt.Printf("%v \t %b \t %#x\n", d, d, d)
	fmt.Printf("%v \t %b \t %#x\n", e, e, e)
	fmt.Printf("%v \t %b \t %#x\n", f, f, f)
}

func conversionFun() {
	y := 42
	x := 42.0
	fmt.Printf("%v of type %T\n", y, y)
	fmt.Printf("%v of type %T\n", x, x)

	var m float32 = 43.742
	fmt.Printf("%v of type %T\n", m, m)

	var z = float64(m)
	fmt.Printf("Z: %v of type %T\n", z, z)

}

func timeTest() {
	fmt.Println("The time is ", time.Now())
	fmt.Println("The time is ", time.DateTime)

}

func mathRand() {
	fmt.Println("My favorite number is:", rand.Intn(11))
	fmt.Println(math.Pi) // Pi
}

func mathAdd(x, y int) int {
	return x + y

}

func swap(x, y string) (string, string) {
	return y, x
}

func namedReturns(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return

}

func zeroValues() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	// 	0 for numeric types,
	// false for the boolean type, and
	// "" (the empty string) for strings.
}

func typeConversions() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	var i int = 42
	var l float64 = float64(i)
	fmt.Printf("the type of %t", l)
}

func shiftingPlace() {
	fmt.Printf("%d \t %b\n", 1, 1)
	fmt.Printf("%d \t %b\n", 1<<1, 1<<1)
	fmt.Printf("%d \t %b\n", 1<<2, 1<<2)
	fmt.Printf("%d \t %b\n", 1<<3, 1<<3)
	fmt.Printf("%d \t %b\n", 1<<4, 1<<4)
	fmt.Printf("%d \t %b\n", 1<<5, 1<<5)
	fmt.Printf("%d \t %b\n", 1<<6, 1<<6)
}

func main() {
	//hexDex()
	//variables()
	//conversionFun()
	//timeTest()
	//mathRand()

	//fmt.Println(mathAdd(4, 5))  //9
	/*--------------------------------------------------*/
	// a, b := swap("Hello", "world")
	// fmt.Println(a, b)
	/*--------------------------------------------------*/

	//fmt.Println(namedReturns(10))
	/*--------------------------------------------------*/

	//zeroValues()
	/*--------------------------------------------------*/
	//typeConversions()
	/*--------------------------------------------------*/

	//shiftingPlace()
	/*--------------------------------------------------*/
	fmt.Printf("You can do it Mark!")
}
