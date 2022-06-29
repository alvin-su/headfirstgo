package main

import (
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, requst *http.Request) {
	message := []byte("hello,web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", viewHandler)
	err := http.ListenAndServe("localhost:9090", nil)
	log.Fatal(err)
}
