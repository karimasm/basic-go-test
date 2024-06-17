package main

import (
	"fmt"
)

func main() {
	// 1. buatlah sebuah variable dengan nama 'Age' dengan tipe data Int8
	var Age int8

	// 2. buatlah sebuah variable dengan nama 'Balance' dengan tipe data Int64
	var Balance int64

	// 3. buatlah sebuah variable jokoAge dengan tipe data Int lalu inisiasikan nilainya dengan 32,
	// buat variable susiAge dengan tipe data int lalu inisialisasikan nilainya dengan 24,
	// setelah itu lakukan operasi pengurangan terhadap variable jokoAge dengan susiAge,
	// lalu simpan nilainnya kedalam variable jokoAndSusiAgeGap
	var jokoAge int8 = 32
	var susiAge int8 = 24
	var jokoAndSusiAgeGap int8 = jokoAge - susiAge
	Age = jokoAndSusiAgeGap

	// 4. budi memiliki usaha pembuatan layang-layang, setiap layang layang yang budi buat membutuhkan batang bambu sejumlah 5,
	// hari ini budi berencana untuk membuat layang - layang dengan jumlah batang bambu 43 buah,
	// tuliskan kode untuk mencari tahu berapa jumlah sisa batang bambu yang tersisa setelah proses pembuatan layang-layang setiap hari
	// (jumlah batang bambu dapat berubah setiap harinya, namun batang bambu yang dibutuhkan tetap 5 batang per layang-layang)
	var initialBamboo int64 = 43
	var bambooPerKite int64 = 5
	Balance = int64(initialBamboo) % bambooPerKite

	fmt.Println("Joko's Age:", jokoAge)
	fmt.Println("Susi's Age:", susiAge)
	fmt.Println("Age Gap between Joko and Susi:", Age)
	fmt.Println("Initial sum of bamboo stick:", initialBamboo)
	fmt.Println("Remaining Bamboo Sticks after making kites:", Balance)
}
