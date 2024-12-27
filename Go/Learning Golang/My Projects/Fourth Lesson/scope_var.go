// Scope Variable in Golang
// mendefinisikan bahwa bagian dari program yang variabel nya dapat di akses
// oleh semua fungsi yang ada di dalamnya
package main

import "fmt"

// terdapat 2 jenis scope yaitu global scope dan local scope
// global artinya variabel dapat di akses oleh semua fungsi
// local artinya variabel hanya dapat di akses oleh fungsi yang membuatnya

// syntax
// var variableName type = value

// global variable declaration

var myVariable int = 200

func main() {
	// local variabel di dalam fungsi main
	var localVar int = 300                                          // local variabel
	fmt.Printf("Di dalam main - Global Variabel: %d\n", myVariable) // akses disini
	fmt.Printf("Di dalam main - Global Variabel: %d\n", localVar)   // akses disini
}
func display() {
	fmt.Printf("Di dalam Display -  Global Variabel: %d\n", myVariable) // akses disini
}
