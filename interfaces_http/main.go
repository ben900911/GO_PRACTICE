package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logwriter struct{}

func main() {
	resp, err := http.Get("http://google.com")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//fmt.Printf("%+v", resp.Body)
	lw := logwriter{}
	io.Copy(lw, resp.Body)

	/*
		//Read() do not resize the input slice
		//var bs []byte
		bs := make([]byte, 99999)
		num, err := resp.Body.Read(bs)
		if err != nil {
			os.Exit(1)
		}

		fmt.Println(string(bs))
		//fmt.Println("2.", *resp)
	*/
}

func (logwriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
