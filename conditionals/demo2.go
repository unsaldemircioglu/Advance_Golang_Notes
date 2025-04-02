package conditionals

import "fmt"

func Demo2() {
	var account float64 = 6000
	var takeMoney float64 = 3900

	if takeMoney > account {
		fmt.Println("insufficient funds")
	} else if account == takeMoney {
		fmt.Println("Your money has been preparing!")
		fmt.Println("Inside bank account money is finish.")
	} else {
		fmt.Println("Your money is preapring-")
	}
}
