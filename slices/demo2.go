package slices

import "fmt"

func Demo2() {
	sehirler := []string{"Anakra", "İstanbul", "İzmir"}
	fmt.Println(sehirler)

	sehirlerKopya := make([]string, len(sehirler))
	fmt.Println(sehirlerKopya)
	copy(sehirlerKopya, sehirler)
	fmt.Println(sehirlerKopya)

	sehirler = append(sehirler, "Adana")
	fmt.Println(sehirler)
	fmt.Println(sehirlerKopya)

	fmt.Println(sehirler[1:3])
	fmt.Println(sehirler[1:4])
	fmt.Println(sehirler[1:])
	fmt.Println(sehirler[:2])
}
