package main

import "fmt"

func main()  {

	if (1 > 2) {
		fmt.Println("1 > 2")
	} else {
		fmt.Println("1 <= 2")
	}

	var grade string = "A"

	switch grade {
	case "A":
		fmt.Println("Very Good!")
	case "B":
		fmt.Println("Good!")
	case "C":
		fmt.Println("Keep Moving!")
	case "D":
		fmt.Println("You should Work Harder!")
	}


}
