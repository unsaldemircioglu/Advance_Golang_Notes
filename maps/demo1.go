package maps

import "fmt"

func Demo1() {
	// key-value
	sozluk := make(map[string]string)
	sozluk["table"] = "Masa"
	sozluk["book"] = "Kitap"
	sozluk["pencil"] = "Kalem"

	fmt.Println(sozluk["book"])
	fmt.Println("Eleman Sayısı: ", len(sozluk))
	delete(sozluk, "book")
	fmt.Println("Eleman Syısı: ", len(sozluk))
	delete(sozluk, "book")
	deger := sozluk["table"]
	fmt.Println(deger)

	deger, varMi := sozluk["pencil"]
	fmt.Println(deger)
	fmt.Println("Listede Olma Durumu: ", varMi)

	sozluk2 := map[string]string{"glass": "bardak", "microphone": "Mikrafon"}
	fmt.Println(sozluk2)

}
