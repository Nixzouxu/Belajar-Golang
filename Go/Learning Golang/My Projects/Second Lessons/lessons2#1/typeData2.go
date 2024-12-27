// now learn about Boolean
package main

import (
	"fmt"
)

func main() {

	var b1 bool = true // typed declaration with initial value
	var b2 = true      // untyped declaration with initial value
	var b3 bool        // typed declaration without initial value
	b4 := true         // untyped declaration with initial value

	fmt.Println(b1) // returns true
	fmt.Println(b2) // returns true
	fmt.Println(b3) // returns false
	fmt.Println(b4) // returns true

	// learn about integers

	/* Integers Keyword
	- int8: 8-bit signed integer (1 byte) => -128 s/d 127
	- int16: 16-bit signed integer ( 2 byte) => -32768 s/d 32767
	- int32: 32-bit signed integer (4 byte) dst
	- int64: 64-bit signed integer (8 byte)
	*/

	var x int = 500
	var y int = -4500
	fmt.Printf("Type: %T, value: %v\n", x, x)
	fmt.Printf("Type: %T, value: %v\n", y, y)

	// another example
	var x11 int32 = 210
	var x22 int32 = 93
	//Penjumlahan
	fmt.Printf(" Penjumlahan :  %d + %d = %d\n ", x11, x22, x11+x22)
	//pengurangan
	fmt.Printf(" pengurangan :  %d - %d = %d\n ", x11, x22, x11-x22)
	//perkalian
	fmt.Printf(" perkalian :  %d * %d = %d\n ", x11, x22, x11*x22)
	//pembagian
	fmt.Printf(" Pembagian :  %d / %d = %d\n ", x11, x22, x11/x22)
	//modulus(sisa bagi)
	fmt.Printf(" modulus :  %d %% %d = %d\n ", x11, x22, x11%x22) // modulus operator

	/* Unsigned Integers (uint) keywords
	- uint8: 8-bit unsigned integer (1 byte) => 0 s/d 255
	- uint16: 16-bit unsigned integer (2 byte) => 0 s/d 65535
	- uint32: 32-bit unsigned integer (4 byte) => 0 s/d 4294967295
	- uint64: 64-bit unsigned integer (8 byte) => 0 s/d 18446744073709551615
	*/

	// using uint16 & uint32
	var a uint16 = 500
	var b uint32 = 4500
	fmt.Printf("Type: %T, value: %v\n", a, a)
	fmt.Printf("Type: %T, value: %v\n", b, b)

	// learn about Floats

	/* Float data Type
	   float32 => 32 bits (4 bytes) => 1.4e-45 s/d 3.4e+38
	   float64 => 64 bits (8 bytes) => 4.9e-324 s/d 1.8e+
	*/
	var x1 float32 = 143.76
	var x2 float32 = 3.4e+38
	var x3 float64 = 1.7e+308
	fmt.Printf("Type: %T, value: %v\n", x1, x1)
	fmt.Printf("Type: %T, value: %v\n", x2, x2)
	fmt.Printf("Type: %T, value: %v\n", x3, x3)

	// another example of floats
	var j1 float32 = 64.00
	var j2 float32 = 28.56
	fmt.Printf("Penjumlahan :  %f + %f = %f\n ", j1, j2, j1+j2)
	fmt.Printf("pengurangan:  %f - %f = %f\n ", j1, j2, j1-j2)
	fmt.Printf("Perkalian :  %f * %f = %f\n ", j1, j2, j1*j2)
	fmt.Printf("Pembagian :  %f / %f = %f\n ", j1, j2, j1/j2)

}

// NEXT To typeDATA3.go
