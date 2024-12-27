// learn about konversi Tipe data di Golang
// yaitu mengubah tipe data yang awal menjadi tipe data lainnya

// syntax
// TipeData(Value)

package main

import "fmt"

func main() {
	var totalJumlah int = 846
	var num int = 19
	var rataRata float32

	// konversi tipe eksplisit
	rataRata = float32(totalJumlah) / float32(num)
	fmt.Printf("Rata-Rata = %f\n", rataRata)

	/* konversi tipe implisit */
	// rataRata = float32(totalJumlah) / num
	// fmt.Printf("Rata-Rata = %f\n", rataRata)
	/* -----------------------------------------------------------------------------------------------------------------------------------------*/
	/* var geek1 int64 = 875

	 it will give compile time error as we
	 are performing an assignment between
	 mixed types i.e. int64 as int type
	var geek2 int = geek1

	var geek3 int = 100

	 it gives compile time error
	 as this is invalid operation
	 because types are mix i.e.
	 int64 and int
	var addition = geek1 + geek3 */

}
