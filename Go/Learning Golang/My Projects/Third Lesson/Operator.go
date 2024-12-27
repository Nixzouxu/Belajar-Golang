// now we learn about Operator in Golang
package main

import "fmt"

/* Operator Aritmatika */
/* (+ , - , * , / , %) */

// contoh Operator Aritmatika
func main() {
	// Operator Penjumlahan
	var a, b int = 10, 20
	var hasilPenjumlahan int = a + b
	// Operator Pengurangan
	var hasilPengurangan int = a - b
	// Operator Perkalian
	var hasilPerkalian int = a * b
	// Operator Pembagian
	var hasilPembagian int = a / b
	// Operator Modulus
	var hasilModulus int = a % b
	// Output
	fmt.Println("Hasil Penjumlahan : ", hasilPenjumlahan)
	fmt.Println("Hasil Pengurangan : ", hasilPengurangan)
	fmt.Println("Hasil Perkalian : ", hasilPerkalian)
	fmt.Println("Hasil Pembagian : ", hasilPembagian)
	fmt.Println("Hasil Modulus : ", hasilModulus)

	fmt.Printf("\n\n")

	/* ------------------------------------------------------------------------------------------------------------------------------------*/
	/* Operator Relasional */
	/* (== , != , > , < , >= , <=) */
	/* ------------------------------------------------------------------------------------------------------------------------------------*/
	// contoh Operator Relasional
	p := 34
	q := 20

	// '==' (sama dengan)
	hasil1 := p == q
	fmt.Println(hasil1)

	// '!=' (tidak sama dengan)
	hasil2 := p != q
	fmt.Println(hasil2)

	// '<' (kurang dari)
	hasil3 := p < q
	fmt.Println(hasil3)

	// '>' (lebih dari)
	hasil4 := p > q
	fmt.Println(hasil4)

	// '>=' (lebih dari atau sama dengan)
	hasil5 := p >= q
	fmt.Println(hasil5)
	// '<=' (kurang dari atau sama dengan)
	hasil6 := p <= q
	fmt.Println(hasil6)

	fmt.Printf("\n")

	/* ------------------------------------------------------------------------------------------------------------------------------------*/
	/* Operator logika */
	/* (&& , || , !) */
	/* ------------------------------------------------------------------------------------------------------------------------------------*/
	// contoh Operator logika

	var x int = 23
	var y int = 60
	// '&&' (dan)
	if x != y && x <= y {
		fmt.Println("True")
	}
	// '||' (atau)
	if x != y || x <= y {
		fmt.Println("True")
	}
	// '!' (tidak)
	if !(x == y) {
		fmt.Println("True")
	}
	/* ------------------------------------------------------------------------------------------------------------------------------------*/
	// assignment operator
	// (:=, =, +=, -=, *=, /=, %=, <<=, >>=)
	/* ------------------------------------------------------------------------------------------------------------------------------------*/
	var u int = 55
	var i int = 80

	// "=" (simple assignment)
	u = i
	fmt.Println(u)
	// "+=" (add assignment)
	u += i
	fmt.Println(u)
	// "=" (subtract assignment)
	u -= i
	fmt.Println(u)
	// "*=" (Multiply assignment)
	u *= i
	fmt.Println(u)
	// "/=" (division assignment)
	u /= i
	fmt.Println(u)
	// "%=" (modulus assignment)
	u %= i
	fmt.Println(u)

	/* ------------------------------------------------------------------------------------------------------------------------------------*/
	// increment dan decrement operator
	// (++, --)
	/* ------------------------------------------------------------------------------------------------------------------------------------*/
	// contoh increment dan decrement operator
	var l int = 10
	var k int = 20
	// "++" (increment)
	l++ // l = l + 1
	fmt.Println(l)
	// "--" (decrement)
	k-- // k = k - 1
	fmt.Println(k)

}
