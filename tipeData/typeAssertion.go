package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 1. rubah tipe data dari variable num ke tipe data int64 ,int8, float64 dan simpan maisng masing hasilnya kedalam variable sendiri
	num := 10
	numInt64 := int64(num)
	numInt8 := int8(num)
	numFloat64 := float64(num)
	//dipake variabelnya
	fmt.Println("numInt64:", numInt64)
	fmt.Println("numInt8:", numInt8)
	fmt.Println("numFloat64:", numFloat64)

	//  2. rubah tipe data nama menjadi string
	var nama interface{}
	nama = "John Doe"
	namaStr := nama.(string)
	fmt.Println("namaStr:", namaStr)

	// 3. rubah tipe data menjadi []byte lalu rubah lagi menjadi string
	var alamat string
	alamat = "Jl Lorem Ipsum"

	alamatBytes := []byte(alamat)
	alamatStr := string(alamatBytes)
	fmt.Println("alamatBytes:", alamatBytes)
	fmt.Println("alamatStr:", alamatStr)

	//iseng ganti alamat wwk
	fmt.Println("Please enter your new house address :3")
	reader := bufio.NewReader(os.Stdin)
	alamat, _ = reader.ReadString('\n')
	alamat = strings.TrimSpace(alamat) // i just learnt that Scanln only could read one whole word

	alamatBytes = []byte(alamat)
	alamatStr = string(alamatBytes)
	fmt.Println("Updated alamatBytes:", alamatBytes)
	fmt.Println("Updated alamatStr:", alamatStr)
}
