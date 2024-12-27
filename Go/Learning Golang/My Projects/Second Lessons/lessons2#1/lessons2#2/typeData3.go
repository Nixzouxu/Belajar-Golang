// Now we Learn about Complex and Strings
package main

import (
	"fmt"
)

func main() {

	/* Complex Numbers */
	/* few built-in functions in complex numbers:
	complex : membuat angka complex dari 2 floats : Real dan Imaginary (real()) && (imag()) */

	/* Complex Data Type */
	/* Complex64 : complex number which contain float32 as a real and imaginary component
	   Complex128 : complex number which contain float64 as a real and imaginary component */

	// example
	var a complex128 = complex(6, 2)
	var b complex64 = complex(9, 2)
	fmt.Println(a)
	fmt.Println(b)

	fmt.Printf("The type of a is %T and "+"the type of b is %T\n", a, b)
	fmt.Printf("\n\n")

	// another example of complex in build-in functions

	comp1 := complex(10, 11)
	// complex number init syntax
	comp2 := 15 + 33i
	fmt.Println("Complex number 1 is : \n", comp1)
	fmt.Println("Complex number 2 is : \n", comp2)
	// mengambil bagian bilangan Real
	realNum := real(comp1)
	fmt.Println("Real part of Complex number 1 is: \n", realNum)
	// mengambil bagian bilangan imajiner
	imaginary := imag(comp2)
	fmt.Println("Imaginary part of Complex number 2 is: \n", imaginary)
	fmt.Printf("\n")

	/* Lanjut sedikit tentang Boolean */
	// membandingkan 2 variabel

	str1 := "Nxozu"
	str2 := "nxozu"
	str3 := "Nxozu"
	hasil1 := str1 == str2
	hasil2 := str1 == str3

	fmt.Println(hasil1)
	fmt.Println(hasil2)

	// display the type of hasil 1&2
	fmt.Printf("The type of result1 is %T and "+"the type of result2 is %T", hasil1, hasil2)

	// aight rn lets talk about Strings
	// when string is created we cant change that strings cant be change anymore
	// then string can be concantenated using plus (+) operator.

	// example
	str := "My niggas"
	fmt.Printf("length of the string is:%d\n", len(str)) // mencetak panjang string
	fmt.Printf("\nString is: %s", str)                   // mencetak string
	fmt.Printf("\nType of the String is: %T", str)       // mencetak tipe dari str

	// String Contantenation example
	var str4 string = "HELLO_"
	var str5 string = "Nixzouxu"

	fmt.Println("\n\nNew String: ", str4+str5) // output: HELLO_Nixzouxu

	/* EXERCISES BEFORE GO TO NEXT LESSONS

		Since we already work with some variables and data types in Go languange i have some study case that u need to solve it with the variables and data types in Go Languange

		here are the study case :
			1) Buatlah program yang meminta pengguna untuk memasukkan panjang dan lebar dari sebuah persegi panjang, kemudian hitung dan tampilkan luas dan kelilingnya.
			2) Buatlah program yang mengkonversi suhu dari Celsius ke Fahrenheit dan sebaliknya. Minta pengguna untuk memasukkan suhu dan jenis konversi yang diinginkan.
			3) Buatlah program yang menghitung rata-rata dari tiga nilai yang dimasukkan. Simpan nilai-nilai tersebut dalam variabel dan tampilkan hasil rata-ratanya.
			4) Buatlah program yang menyimpan informasi tentang buku, termasuk judul, penulis, dan tahun terbit. Gunakan tipe data yang sesuai untuk setiap informasi.
			5) Anda diminta untuk membuat program sederhana yang menyimpan data mahasiswa. Setiap mahasiswa memiliki nama, umur, dan IPK (Indeks Prestasi Kumulatif).

	 Good luck <3
	*/
}
