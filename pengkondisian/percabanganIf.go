package pengkondisian

import "fmt"

func PengkondisianIf() {
	const umur = 27
	var status string
	//  buat pengkondisian dengan ketentuan sebagai berikut

	// 1.  status=  anak-anak jika umur kurang dari <= 12 tahun
	// 1.  status=  remaja jika umur berada diantara 13-20 tahun
	// 1.  status=  dewasa jika umur berada diantara 21-50 tahun
	// 1.  status=  manula jika umur  >= 51 tahun
	if umur <= 12 {
		status = "anak-anak"
	} else if umur >= 13 && umur <= 20 {
		status = "remaja"
	} else if umur >= 21 && umur <= 50 {
		status = "dewasa"
	} else if umur >= 51 {
		status = "manula"
	} else {
		status = "umur tidak valid"
	}

	fmt.Printf("DenganStatus: %s\n", status)
}
