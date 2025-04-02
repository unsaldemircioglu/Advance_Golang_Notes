package slices

import "fmt"

func Demo1() {
	// isimler := make([]string, 3)

	// fmt.Println(isimler)

	isimler := make([]string, 3)

	fmt.Println(isimler) // [  ]
	isimler[0] = "Ünsal"
	isimler[1] = "ünsal"
	isimler[2] = "unsal"
	isimler = append(isimler, "ünsal2")
	fmt.Println(len(isimler))
	fmt.Println(isimler) // [Ünsal ünsal unsal]
}
