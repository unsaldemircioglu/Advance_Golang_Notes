package channels

func CiftSayilar(ciftSayiCn chan int) {
	toplam := 0
	for i := 0; i < 10; i += 2 {
		toplam = toplam + i
	}

	ciftSayiCn <- toplam
}

func TekSayilar(tekSayiCn chan int) {
	toplam := 0
	for i := 1; i < 10; i += 2 {
		toplam = toplam + i
	}

	tekSayiCn <- toplam
}
