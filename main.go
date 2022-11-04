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

	for _, adr := range addresses {
		checkUrl(adr)
	}
}

func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "might be down")
		return
	}

	fmt.Println(url, "is up and running!")

}
