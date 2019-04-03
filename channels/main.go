package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://gogle.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		//block call
		//checkLink(link)
		//go + func  : run the func with new goroutine
		go checkLink(link, c)
	}
	//wait for the channel
	for L := range c {

		//go checkLink(L, c)
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(L)
	}
}

func checkLink(link string, c chan string) {
	//time.Sleep(5 * time.Second)
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, " might be down!")
		c <- link
		return
	}
	fmt.Println(link, " is up!")
	c <- link
}
