package main

import "fmt"

// 1. buat struct Person dengan komposisi sbb :
// - memiliki property Name dengan tipe String
// - memiliki property Age dengan tipe int
// - memiliki property Is married dengan tipe bool
type Person struct {
	Name      string
	Age       int
	IsMarried bool
}

// 2. buat method sayHello untuk struct Person
// dalam method sayHello print "hello "+ nama person ke konsol
func (p Person) sayHello() {
	fmt.Println("Hello " + p.Name)
}

func main() {
	var name string
	var age int
	var isMarried bool

	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&name)

	fmt.Print("Masukkan umur: ")
	fmt.Scanln(&age)

	fmt.Print("Apakah sudah menikah? (true/false): ")
	fmt.Scanln(&isMarried)

	// 3. Buat struct Person baru lalu isi datanya dengan data dari input pengguna
	person := Person{
		Name:      name,
		Age:       age,
		IsMarried: isMarried,
	}

	person.sayHello()
}
