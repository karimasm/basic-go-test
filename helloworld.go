package main

import (
	"fmt"
	"go-basic-test/function"
	"go-basic-test/pengkondisian"
	"go-basic-test/perulangan"
)

func main() {
	// tuliskan kode untuk mencetak 'Hello World' ke console disini
	function.HelloWorld()

	// sebagian fungsi sengaja diimport kesini, tapi punten sisanya males wkwk bisa run langsung di filenya thankies
	for {
		fmt.Println("Pilih fungsi yang ingin dijalankan:")
		fmt.Println("\033[1mSimple Function\033[0m")
		fmt.Println("1. HelloWorld")
		fmt.Println("2. SayHello")
		fmt.Println("3. GetTotal")
		fmt.Println("\033[1mPengkondisian\033[0m")
		fmt.Println("4. PengkondisianIf")
		fmt.Println("\033[1mPerulangan\033[0m")
		fmt.Println("5. Perulangan 10 kali")
		fmt.Println("6. Perulangan-break dengan ketentuan tertentu")
		fmt.Println("7. Perulangan-skip-continue dengan ketentuan tertentu")
		fmt.Println("8. Perulangan kelipatan 7")
		fmt.Println("\033[1m9. Exit\033[0m")

		var choice int
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			function.HelloWorld()
		case 2:
			function.SayHello()
		case 3:
			function.GetTotal()
		case 4:
			pengkondisian.PengkondisianIf()
		case 5:
			perulangan.LoopSepuluhKali()
		case 6:
			perulangan.LoopBreakDiTiga()
		case 7:
			perulangan.LoopSkipKelipatanTiga()
		case 8:
			perulangan.LoopKelipatanTujuh()
		case 9:
			fmt.Println("Exit.")
			return
		default:
			fmt.Println("Duh tidak valid. Hayu coba lagi. Semangat")
		}

		fmt.Println("----------")
	}
}
