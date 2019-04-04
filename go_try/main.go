package main

import "fmt"

func testA() *int {

	var test int
	test = 5

	return &test
}

func main() {
	var test *int

	test = testA()

	fmt.Println(*test)
}
