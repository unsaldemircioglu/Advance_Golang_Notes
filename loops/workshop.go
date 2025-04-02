package loops

import "fmt"

func Demo3() {
	number := 80
	randomNumber := 0

	fmt.Print("Number is: ")
	fmt.Scanln(&randomNumber)
	fmt.Println(randomNumber)

	for randomNumber != number {
		if randomNumber < number {
			fmt.Println("Upper")
			fmt.Scanln(&randomNumber)
		}

		if randomNumber > number {
			fmt.Println("Down")
		}

		if randomNumber == number {
			fmt.Println("Correct")
		}
	}

}
