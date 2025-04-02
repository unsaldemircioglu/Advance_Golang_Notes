package errorhandling

import (
	"errors"
	"fmt"
)

func TahminEt(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir sayı giriniz,")
	}
	return "Bildiniz", nil
}

func Demo2() {
	mesaj, _ := TahminEt(50)
	fmt.Println(mesaj)

}
