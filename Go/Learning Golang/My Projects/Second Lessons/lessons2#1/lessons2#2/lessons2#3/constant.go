package main

import "fmt"

// now learn about Constant
// Constant atau Konstanta adalah keyword di Bahasa Go untuk mengdeklarasikan variabel Constant
// Constant tidak dapat diubah nilainya setelah deklarasinya
// dan tidak bisa di deklarasikan dengan cara :=

// syntax
// const nama = nilai

// example 1

const (
	GFG     = "geeks"
	PI      = 3.14
	Correct = true
)

func main() {
	const NXZ = "Nxozu"
	fmt.Println("Hello", NXZ) // mencetak nilai constant NXZ

	fmt.Println("Happy", PI, "Day") // mencetak nilai constant PI
	const Correct = true
	fmt.Println("Go Rules?", Correct) // mencetak nilai constant Correct

	/* Integer Constant */
	/*
		85             decimal
		0213           octal
		0x4b       	   hexadecimal
		30             int
		30u            unsigned int
		30l            long
		30ul           unsigned long
		212            Legal
		215u           Legal
		0xFeeL         Legal
		078            Illegal: 8 is not an octal digit
		032UU          Illegal: cannot repeat a suffix
	*/

	/* Floating Constant */
	/*
		3.14159        Legal
		314159E-5L     Legal
		510E           Illegal: incomplete exponent
		210f           Illegal: no decimal or exponent
		.e55           Illegal: missing integer or fraction
	*/
	// string Constants

	const A = "MY NIGGAS"
	var b = "my lovely niggas"

	// concat strings
	var hello = A + "" + b
	hello += "!"
	fmt.Println(hello)

	// Compare Strings
	fmt.Println(A == "MY NIGGAS")
	fmt.Println(b < A)

	// contoh lain
	const trueConst = true

	type myBool bool
	var defaultBool = trueConst       // allowed
	var customBool myBool = trueConst // allowed

	// defaultBool = customBool (Not allowed)
	fmt.Println(defaultBool) // true
	fmt.Println(customBool)  // true

	// Move into Third Lessons

}
