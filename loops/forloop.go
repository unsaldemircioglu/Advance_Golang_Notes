package loops

import "fmt"

func Demo1() {
	var word string = "Text"
	for i := 0; i < 5; i++ {
		fmt.Println(word)
	}

	// infinity loop
	i := 0
	for i <= 5 {
		fmt.Println(word)
		i = i + 1
	}
	fmt.Println("Finished")
}
