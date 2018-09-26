package main

import "fmt"

func main() {
	mySlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range mySlice {
		if value % 2 == 0 {
			fmt.Printf("%d is even\n", value)
		} else {
			fmt.Println(value, " is odd")
		}
	}
}
