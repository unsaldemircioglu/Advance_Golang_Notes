package conditionals

import "fmt"

func Demo1() {
	var account float64 = 6000
	var takeMoney float64 = 3900

	if takeMoney > account {
		fmt.Println("insufficient funds")
	}

	if takeMoney <= account {
		fmt.Println("Succesful, and preparing your money")
		account = account - takeMoney
	}
	// fmt.Println("Finished. Now inside account active has : " + fmt.Sprintf("%v", account))
	fmt.Printf("Finished. Now inside account active has : %v ", account)
}
