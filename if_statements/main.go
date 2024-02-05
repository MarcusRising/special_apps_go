package main

import (
	"fmt"
	"math/rand"
)

/*
	Comparison operators
	Comparison operators compare two operands and yield an untyped boolean value.

	==    equal
	!=    not equal
	<     less
	<=    less or equal
	>     greater
	>=    greater or equal
*/
// https://go.dev/ref/spec#Comparison_operators

func init() {
	fmt.Printf("Start of if/then statements\n")
}

func someStatements() {
	x := 50
	y := 0
	if x > y {
		fmt.Printf("X: %v is > then Y: %v\n", x, y)
	} else if x == y {
		fmt.Printf("They are the same X: %v and Y: %v\n", x, y)
	} else if x < y {
		fmt.Printf("X: %v is < then Y: %v\n", x, y)
	} else {
		fmt.Printf("error")
	}
}

//////////////////////////////////////////////////
/*
	Logical operators
	Logical operators apply to boolean values
	and yield a result of the same type as the operands.
	The right operand is evaluated conditionally.

	&&    conditional AND    p && q  is  "if p then q else false"
	||    conditional OR     p || q  is  "if p then true else q"
	!     NOT                !p      is  "not p"
*/

func logicalStatements() {
	x := 40
	y := 5
	if x == 42 || y == 42 {
		fmt.Printf("Meaning of life found!\n")
	} else if x < 42 && y < 42 {
		fmt.Printf("Both are less than the meaning of life\n")
	} else if x > 30 || x < 42 {
		fmt.Printf("x is getting close to the meaning of life\n")
	}
	if x != 42 {
		fmt.Printf("X is not 42\n")
	}
	if y != 42 {
		fmt.Printf("Y is not 42\n")
	}
}

//////////////////////////////////////////////////
/*
	The expression [evaluated in an if statement ]may be preceded by a simple statement,
	which executes before the expression is evaluated.
*/
// https://go.dev/ref/spec#If_statements
/*
	The comma ok idiom is also like this
*/
// https://go.dev/play/p/OXGzjxVkag0

func statementStatementId() {
	x := 40
	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is GREATER THAN OR EQUAL x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is LESS THAN x which is %v\n", z, x)
	}
}

func main() {
	//someStatements()
	//logicalStatements()
	statementStatementId()
}
