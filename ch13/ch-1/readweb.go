package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func responseSize(url string, channel chan Page) {
	//fmt.Println("Getting:", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(len(body))
	page := Page{url: url, size: len(body)}
	//channel <- len(body)
	channel <- page
}

type Page struct {
	url  string
	size int
}

func main() {
	//sizes := make(chan int)

	//urls := []string{"https://example.com", "https://www.baidu.com", "https://platform.innostic.com"}

	/*go responseSize("https://example.com", sizes)
	go responseSize("https://www.baidu.com", sizes)
	go responseSize("https://platform.innostic.com", sizes)*/

	//time.Sleep(time.Second * 5)
	/*fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)*/

	/*	for _, url := range urls {
			go responseSize(url, sizes)
			//fmt.Println(<-sizes)
		}
		for i := 0; i < len(urls); i++ {
			fmt.Println(<-sizes)
		}*/
	pages := make(chan Page)
	urls := []string{"https://example.com", "https://www.baidu.com", "https://platform.innostic.com"}
	for _, url := range urls {
		go responseSize(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s:%d\n", page.url, page.size)
	}
}
