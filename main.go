package main

import (
	"fmt"
	"net/http"
)

func main() {

	addresses := []string{"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, adr := range addresses {
		go checkUrl(adr, c)
	}

	for i := 0; i < len(addresses); i++ {
		fmt.Println(<-c)
	}

}

func checkUrl(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down")
		c <- "Migh be down"
		return
	}

	fmt.Println(url, "is up and running!")
	c <- "is up and running!"

}
