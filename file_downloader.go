package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	if len(os.Args) > 3 || len(os.Args) < 2 {
		log.Fatalf("Usage: ./file_downloader URL num_threads")
	}

	url := os.Args[1]
	numThreads, _ := strconv.Atoi(os.Args[2])

	fmt.Println(url)
	fmt.Println(numThreads)

	res, err := http.Head(url)
	if err != nil {
		panic(err)
	}

	contentLength := int(res.ContentLength)

	fmt.Println(contentLength)

	bytesPerThread := contentLength / numThreads
	fmt.Println(bytesPerThread)

	startByte := 0
	endByte := bytesPerThread

	for i := 0; i < numThreads; i++ {

		wg.Add(1)
		go downloadPart(url, startByte, endByte, &wg)

		startByte = endByte + 1
		endByte = startByte + bytesPerThread

		if endByte > contentLength {
			endByte = contentLength
		}
	}

	wg.Wait()

}

func downloadPart(url string, startByte int, endByte int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Println("start: ", startByte)
	fmt.Println("end: ", endByte)

	client := new(http.Client)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Range", fmt.Sprintf("bytes=%d-%d", startByte, endByte))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.ContentLength)

}
