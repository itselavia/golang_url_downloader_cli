# golang_url_downloader_cli
A simple golang utility to download file from a URL

### Usage
1. Clone the repo
2. cd into the directory (Make sure you have Golang installed)
3. go run file_downloader.go URL NUM_THREADS<br/>
   Example: go run file_downloader.go http://ipv4.download.thinkbroadband.com/100MB.zip 100
   
### Script Design and Approach
This Go utility is implemented using Goroutines and WaitGroups to manage synchronization among the Goroutines. <br/>
It makes an HTTP HEAD request to the URL to fetch the content length of the file.<br/>
Then, it divides the number of bytes to be downloaded by the number of threads.<br/>
Goroutines are launched with appropriate start and end byte ranges. Also, these are added to the WaitGroup when they are launched.<br/>
The WaitGroup waits for all the Goroutines to finish before the script terminates.<br/>
```go
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
```

Each Goroutine makes an HTTP GET request to the URL with a Range header specifying the byte ranges to be downloaded.<br/>
The downloaded content is written to the output file using the WriteAt function to write the bytes at a specific offset in the file.<br/>
```go
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
```
