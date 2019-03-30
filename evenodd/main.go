package main

import "fmt"

func main() {
	ten := []int{}

	for i := 0; i <= 10; i++ {
		ten = append(ten, i)

		if i%2 == 0 {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
}
