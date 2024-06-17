package main

import "fmt"

func main() {
	// 1. Buat variable firstName dengan tipe data string lalu isi dengan nama depan anda
	firstName := "Karima"

	// 2. Buat variable middleName dengan tipe data string lalu isi dengan nama tengah anda
	middleName := "Siti"

	// 3. Buat variable lastName dengan tipe data string lalu isi dengan nama belakang anda
	lastName := "Masna"

	// 4. Gabungkan variable firstName, middleName, lastName kedalam variable baru fullName lalu cetak ke konsol variable fullName
	fullName := firstName + " " + middleName + " " + lastName
	fmt.Println(fullName)
}
