package main

import (
	"fmt"
)

func main() {
	// Setup X dan Y
	var x int
	var y int
	fmt.Println("Masukan angka pertama : ")
	fmt.Scanln(&x)
	fmt.Println("Masukan angka kedua : ")
	fmt.Scanln(&y)

	// Hasil akhirnya
	var jumlah, kurang, kali, bagi = operator(x, y)
	fmt.Println("Hasil Penjumlahan : ", jumlah)
	fmt.Println("Hasil Pengurangan : ", kurang)
	fmt.Println("Hasil Perkalian : ", kali)
	fmt.Println("Hasil Pembagian\t : ", bagi)
}

func operator(x int, y int) (int, int, int, int) {
	// Penjumlahan
	var jumlah = x + y

	// Pengurangan
	var kurang = x - y

	// Perkalian
	var kali = x * y

	// Pembagian
	var bagi = x / y

	return jumlah, kurang, kali, bagi
}
