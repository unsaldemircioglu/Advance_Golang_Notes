package deferstatement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func Log() {
	fmt.Println("Loglandı!")
}

func (p product) save() {
	fmt.Println("Ürün kaydedildi!")
	defer Log()
}

func Demo3() {
	p := product{productName: "Laptop", unitPrice: 5000}
	p.save()
	p = product{productName: "Mouse", unitPrice: 45}

	fmt.Println("İşlem Başarılı!")
	fmt.Println(p.productName)
	defer p.save()
}
