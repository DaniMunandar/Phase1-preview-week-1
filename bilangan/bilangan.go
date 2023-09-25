package main

import "fmt"

func main() {
	var number int
	fmt.Print("Please enter a number : ")
	fmt.Scanln(&number)

	if number%2 == 0 {
		fmt.Println("The number is even.")
	} else {
		fmt.Println("The number is odd.")
	}

	fmt.Println("Even numbers upto your input are:")
	for i := 1; i <= number; i++ {
		if i%2 == 0 {
			fmt.Printf("%d \n", i)
		}
	}

}
