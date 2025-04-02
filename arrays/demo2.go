package arrays

import (
	"fmt"
)

func Demo2() {
	var citys [5]string
	citys[0] = "Ankara"
	citys[1] = "Kastamonu"
	citys[2] = "Istanbul"
	citys[3] = "Erzurum"
	citys[4] = "Şanlıurfa"
	fmt.Println(citys)

	for i := 0; i < 5; i++ {
		fmt.Println(citys[i])
	}

}
