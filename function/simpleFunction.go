package function

import "fmt"

//1. buat function helloWorld yang akan mencetak "hello world" ke konsol

func HelloWorld() {
	fmt.Println("Hello, World!")
}

// 2. buat function sayHello yang akan menerima parameter berupa string, lalu cetak "hello " + nama ke konsol
func SayHello() {
	var name string
	fmt.Print("Masukkan nama Anda: ")
	fmt.Scanln(&name)
	fmt.Println("Hello " + name)
}

// 3. buat function getTotal yyang akan menerima 3 buah parameter bertipe int, lalu kembalikan nilai dari (a x b) : c) * a + (a + b + c)
func GetTotal() {
	var a, b, c int

	fmt.Println("Enter value for a:")
	fmt.Scan(&a)

	fmt.Println("Enter value for b:")
	fmt.Scan(&b)

	fmt.Println("Enter value for c:")
	fmt.Scan(&c)

	if c == 0 {
		fmt.Println("Error: Division by zero is not allowed.")
		return
	}

	result := ((a*b)/c)*a + (a + b + c)
	fmt.Printf("The result of (a x b) : c) * a + (a + b + c) is: %d\n", result)
}
