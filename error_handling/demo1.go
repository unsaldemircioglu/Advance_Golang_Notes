package errorhandling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("demo1.txt")
	//nil
	if err != nil {
		fmt.Println("Dosya Bulunamadı")
	} else {
		fmt.Println(f.Name())
	}
}
