package perulangan

import "fmt"

// biar ga pusing baca outputnya jadi ku pisah fungsinya
// 1. buat perulangan sebanyak 10 kali dengan setiap perulangannya melakukan print ke konsol berupa "ini merupakan perulangan ke (jumlah perulangan)"
func LoopSepuluhKali() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("ini merupakan perulangan ke- %d\n", i)
	}
	fmt.Println()
}

//  2. buat perulangan sebanyak 10 kali dengan setiap perulangannyaa melakukan print ke konsol berupa "ini merupakan perulangan ke (jumlah perulangan)"
//     apabila mencapai perulangan ke 3 hentikan perulangan tersebut
func LoopBreakDiTiga() {
	for i := 1; i <= 10; i++ {
		if i == 3 {
			break
		}
		fmt.Printf("ini merupakan perulangan ke- %d\n", i)
	}
	fmt.Println()
}

//  2. buat perulangan sebanyak 10 kali dengan setiap perulangannya melakukan print ke konsol berupa "ini merupakan perulangan ke (jumlah perulangan)"
//     apabila mencapai perulangan yang merupakan kelipatan 3 skip perulangan tersebut dan langsung ke perulangan berikutnya
func LoopSkipKelipatanTiga() {
	for i := 1; i <= 10; i++ {
		if i%3 == 0 {
			continue
		}
		fmt.Printf("ini merupakan perulangan ke- %d\n", i)
	}
	fmt.Println()
}

//  2. buat perulangan sebanyak 30 kali dengan setiap perulangannyaa melakukan print ke konsol berupa "ini merupakan perulangan ke (jumlah perulangan)"
//
// apabila mencapai perulangan yang merupakan kelipatan 7 print "ini adalah kelipatan 7"
func LoopKelipatanTujuh() {
	multipleCounter := 0

	for i := 1; i <= 30; i++ {
		if i%7 == 0 {
			multipleCounter++
			fmt.Printf("ini adalah kelipatan 7 yang ke-%d\n", multipleCounter)
		} else {
			fmt.Printf("ini merupakan perulangan ke %d\n", i)
		}
	}
}
