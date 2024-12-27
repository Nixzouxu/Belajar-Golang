package main

import (
	"fmt"
)

func main() {
	// learn Variables In Go
	var x float64
	x = 30.0
	y := 42
	fmt.Println(x)                     // 30.0
	fmt.Println(y)                     // 42
	fmt.Printf("x is of type %T\n", x) // mencetak tipe dari si "x" (float64)
	fmt.Printf("y is of type %T\n", y) // mencetak tipe dari si "y" (int)

	var a, b, c = 3, 4, "foo"
	fmt.Println(a)                     // 3
	fmt.Println(b)                     // 4
	fmt.Println(c)                     // foo
	fmt.Printf("a is of type %T\n", a) // the type of a which means integer (3)
	fmt.Printf("b is of type %T\n", b) // the type of b which means integer (4)
	fmt.Printf("c is of type %T\n", c) // the type of c which means strings (foo)

	// Constant In Go

	// constant syntax
	/* const variable type = value; */

	const LENGTH int = 20
	const WIDTH int = 10

	var area int
	area = LENGTH * WIDTH
	fmt.Printf("value of area : %d\n", area)

	// Go Outputs

	var i, j string = "Hello", "Nxozu"

	// using Print
	fmt.Print(i, j) // mencetak argumen dengan format default
	// using Println
	fmt.Println(i, j) // langsung mencetak dengan whitespace(spasi)
	// using Printf
	var k string = "hello"
	var r int = 15
	fmt.Printf("k has value: %v and type: %T\n", k, k)
	fmt.Printf("r has value: %v and type: %T\n", r, r)

}
