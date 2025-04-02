package main

import (
	"golesson/project"
)

func main() {
	// variables.Demo1()
	// conditionals.Demo1()
	// functions.SelamVer()
	// var sonuc = functions.Topla(2, 3)
	// fmt.Println(sonuc * 10)

	// sonuc1, sonuc2, sonuc3, sonuc4 := functions.DortIslem(10, 6)

	// fmt.Println("Toplam : ", sonuc1)
	// fmt.Println("Çıkartma : ", sonuc2)
	// fmt.Println("Çarpma : ", sonuc3)
	// fmt.Println("Bölme : ", sonuc4)

	// sonuc1, sonuc2, sonuc3, _ := functions.DortIslem(10, 6)

	// fmt.Println("Toplam : ", sonuc1)
	// fmt.Println("Çıkartma : ", sonuc2)
	// fmt.Println("Çarpma : ", sonuc3)
	// // fmt.Println("Bölme : ", sonuc4)

	// fmt.Println(functions.ToplaVariadic(1, 4, 6, 3, 10))
	// fmt.Println(functions.ToplaVariadic(1, 4))
	// fmt.Println(functions.ToplaVariadic(1, 4, 6, 3, 10))

	// sayilar := []int{4, 6, 7, 0, 10}

	// fmt.Println(functions.ToplaVariadic(sayilar...))

	// sayi := 20
	// pointers.Demo1(&sayi)
	// fmt.Println("Main'deki Sayı: ", sayi)

	// sayi := 20
	// pointers.Demo1(&sayi)
	// fmt.Println("Maindeki Sayı: ", sayi)

	//-------------------------------------------
	// has a error :(
	// sayilar := []int{1, 2, 3}
	// pointers.Demo2(sayilar)
	// fmt.Println("Maindeki Sayı: ", sayilar[0])
	//---------------------------------------------

	// go goroutines.CiftSayilar()
	// go goroutines.TekSayilar()
	// time.Sleep(5 * time.Second) // five second
	// fmt.Println("Main Bitti")

	//---------------------------------------------

	// ciftSayiCn := make(chan int)
	// tekSayiCn := make(chan int)
	// go channels.CiftSayilar(ciftSayiCn)
	// go channels.TekSayilar(tekSayiCn)

	// time.Sleep(5 * time.Second)
	// fmt.Println("Main Out!")
	//---------------------------------------------

	// interfaces.Demo1()

	// resful.Demo2()

	project.GetAllProducts()
}
