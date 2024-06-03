package main

import (
	"fmt"
	"net/http"
	"time"
)

func checkWebSite(url string, c chan string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(url, "is down")
		c <- url
		return
	}
	fmt.Println(url, "is up")
	c <- url
}

func main() {
	links := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://www.baidu.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkWebSite(link, c)
	}

	for l := range c {
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkWebSite(link, c)
		}(l)
	}
}
