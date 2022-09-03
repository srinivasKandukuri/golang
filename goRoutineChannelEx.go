package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	links := []string{
		"https://github.com/fabpot",
		"https://github.com/andrew",
		"https://github.com/taylorotwell",
		"https://github.com/egoist",
		"https://github.com/HugoGiraudel",
		"https://github.com/srinivasKandukuri",
		"http://google.com",
	}

	checkUrls(links)
}

func checkUrls(links []string) {
	c := make(chan string)
	var wg sync.WaitGroup

	for _, link := range links {
		wg.Add(1)
		go checkUrl(link, c , &wg)
	}

	go func() {
		wg.Wait()	// this blocks the goroutine until WaitGroup counter is zero
		close(c)    // Channels need to be closed, otherwise the below loop will go on forever
	}()  

	for msg := range c {
		fmt.Println(msg)
	}
}

func checkUrl(url string, c chan string, wg *sync.WaitGroup) {
	defer (*wg).Done()
	_, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Println("We could not reach:", url)
	} else {
		fmt.Println("Success reaching the website:", url)
	}
}
