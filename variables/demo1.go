package variables

import "fmt"

func Demo1() {
	var metin string = "Hello Turkey"
	fmt.Print(metin)

	var kdv int = 15
	fmt.Println(kdv)
	fmt.Println(100 + (kdv * 10 / 100))

	var kdv2 float32 = 0.1
	fmt.Print(kdv2)

	kdv3 := 25
	fmt.Print(kdv3)
	fmt.Print("Veri Tipi: %T", kdv3)

}
