package main

import "fmt"

func main() {
	// Membuat slice menggunakan literal
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice awal:", numbers)

	// Menambahkan elemen ke slice
	numbers = append(numbers, 6, 7)
	fmt.Println("Slice setelah append:", numbers)

	// Iterasi menggunakan loop for
	fmt.Println("Iterasi menggunakan loop for:")
	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	// Iterasi menggunakan loop range
	fmt.Println("Iterasi menggunakan loop range:")
	for i, num := range numbers {
		fmt.Printf("Indeks: %d, Nilai: %d\n", i, num)
	}

	// Membuat sub-slice
	subSlice := numbers[2:5]
	fmt.Println("Sub-slice:", subSlice)

	// Mengubah sub-slice mempengaruhi slice asli
	subSlice[0] = 10
	fmt.Println("Sub-slice yang dimodifikasi:", subSlice)
	fmt.Println("Slice asli setelah sub-slice dimodifikasi:", numbers)
}
