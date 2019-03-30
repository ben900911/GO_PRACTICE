package main

import (
	"io"
	"os"
)

func main() {
	input := os.Args

	//fmt.Print(input[1])

	file, err := os.OpenFile(input[1], os.O_RDWR|os.O_CREATE, 0755)

	if err != nil {
		//fmt.Println(err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, file)
}
