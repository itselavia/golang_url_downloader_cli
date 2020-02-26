package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {

	if len(os.Args) > 3 || len(os.Args) < 2 {
		log.Fatalf("Usage: ./file_downloader URL num_threads")
	}

	url := os.Args[1]
	numThreads := os.Args[2]

	fmt.Println(url)
	fmt.Println(numThreads)

	res, err := http.Head(url)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.ContentLength)

}
