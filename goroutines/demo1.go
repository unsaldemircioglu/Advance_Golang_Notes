package goroutines

import "fmt"

func CiftSayilar() {
	for i := 0; i < 10; i += 2 {
		fmt.Println("Çift Sayı: ", i)
	}
}

func TekSayilar() {
	for i := 1; i < 10; i += 2 {
		fmt.Println("Tek Sayı: ", i)
	}
}
