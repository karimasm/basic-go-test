package tipeData

import "fmt"

func Array() {
	// 1. buat variable listOfNumber dengan tipe data array of int dengan panjang 10
	// isi datanya dengan nilai dari 1 -10
	//  rubah data dengan index ke 3 dengan nilai 99

	listOfNumber := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	listOfNumber[3] = 99
	fmt.Println(listOfNumber)

	// 2. buat variable sliceOfNumber dengan tipe data slice of int
	// isi datanya dengan nilai dari 1 - 20
	sliceOfNumber := make([]int, 20)
	for i := 0; i < 20; i++ {
		sliceOfNumber[i] = i + 1
	}

	//  rubah data terakhir dengan nilai 99
	sliceOfNumber[len(sliceOfNumber)-1] = 99
	fmt.Println(sliceOfNumber)

}
