package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://gogle.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, link := range links {
		//block call
		//checkLink(link)
		//go + func  : run the func with new goroutine
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down!")
		return
	}

	fmt.Println(link, " is up!")
}
