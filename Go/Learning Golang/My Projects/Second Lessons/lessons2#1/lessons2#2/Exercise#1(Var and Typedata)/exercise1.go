// Exercise 1 //

/*
	first Program

Buatlah program yang meminta pengguna untuk memasukkan panjang dan lebar dari sebuah persegi panjang, kemudian hitung dan tampilkan luas dan kelilingnya.
*/
package main

import "fmt"

func main() {

	var panjang float64 = 20.0
	var lebar float64 = 8.0
	var luas float64
	var keliling float64

	luas = panjang * lebar
	keliling = 2 * (panjang * lebar)

	fmt.Printf("Luas dari Persegi panjang tersebut adalah: %2.f\n", luas)
	fmt.Printf("Keliling dari persegi panjang tersebut adalah: %2.f\n", keliling)
	fmt.Printf("\n\n")
	// Second Project
	/* Buatlah program yang mengkonversi suhu dari Celsius ke Fahrenheit dan sebaliknya. Minta pengguna untuk memasukkan suhu dan jenis konversi yang diinginkan. */
	var celsius float64 = 35.8
	var fahrenheit float64 = 42.8

	// converting Celcius to Fahrenheit
	fahrenheit = (celsius * 9 / 5) + 32
	fmt.Printf("Suhu dalam Celcius adalah : %2.f\n", celsius)
	fmt.Printf("Suhu dalam Fahrenheit adalah : %2.f\n", fahrenheit)
	fmt.Printf("\n\n")
	// Third Project
	/* Buatlah program yang menghitung rata-rata dari tiga nilai yang dimasukkan. Simpan nilai-nilai tersebut dalam variabel dan tampilkan hasil rata-ratanya. */

	var x1, x2, x3 float64
	var rataRata float64

	fmt.Print("Masukkan Nilai x1: ")
	fmt.Scanln(&x1)
	fmt.Print("Masukkan Nilai x2: ")
	fmt.Scanln(&x2)
	fmt.Print("Masukkan Nilai x3: ")
	fmt.Scanln(&x3)

	rataRata = (x1 + x2 + x3) / 3
	fmt.Printf("Rata-rata nilai: %.2f\n", rataRata)
	fmt.Printf("\n\n")

	// dan lanjutkan dengan program lainnya

	// ==> Move into the next lessons

}
