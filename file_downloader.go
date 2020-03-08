package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	if len(os.Args) > 3 || len(os.Args) < 2 {
		log.Fatalf("Usage: ./file_downloader URL num_threads")
	}

	url := os.Args[1]
	numThreads, _ := strconv.Atoi(os.Args[2])

	outputFile, err := os.Create(url[strings.LastIndex(url, "/")+1:])
	if err != nil {
		panic(err)
	}

	res, err := http.Head(url)
	if err != nil {
		panic(err)
	}

	contentLength := int(res.ContentLength)

	bytesPerThread := contentLength / numThreads

	startByte := 0
	endByte := bytesPerThread

	for i := 0; i < numThreads; i++ {

		wg.Add(1)
		go downloadPart(url, startByte, endByte, outputFile, &wg)

		startByte = endByte + 1
		endByte = startByte + bytesPerThread

		if endByte > contentLength {
			endByte = contentLength
		}
	}

	wg.Wait()

}

func downloadPart(url string, startByte int, endByte int, outputFile *os.File, wg *sync.WaitGroup) {

	defer wg.Done()

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

	partBytes, err := ioutil.ReadAll(resp.Body)

	outputFile.WriteAt(partBytes, int64(startByte))

}
