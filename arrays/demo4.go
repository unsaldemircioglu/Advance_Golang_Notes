package arrays

import "fmt"

func Demo4() {
	var sayilar [2][3]int

	sayilar[0][0] = 1
	sayilar[1][1] = 2
	sayilar[0][1] = 2

	fmt.Println(sayilar[1][1])

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(sayilar[i][j])
			fmt.Println()
		}
		fmt.Println()
	}

}
