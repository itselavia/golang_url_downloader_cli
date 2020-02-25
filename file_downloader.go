package main

import (
	"fmt"
	"net/http"
)

func main() {

	res, err := http.Head("https://www.bluedata.com/wp-content/uploads/dlm_uploads/2019/04/BlueData-EPIC-Software-Architecture-technical-white-paper.pdf")
	if err != nil {
		panic(err)
	}
	fmt.Println(res.ContentLength)

}
