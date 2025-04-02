package deferstatement

import "fmt"

func A() {
	fmt.Println("A fonksiyonu çalıştı.")
}

func C() {
	fmt.Println("C fonksiyonu çalıştı.")
}

func D() {
	fmt.Println("D fonksiyonu çalıştı.")
}

func B() {
	defer A()
	defer C()
	defer D()
	fmt.Println("B fonksiyonu çalıştı.")
}
