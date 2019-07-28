package main

import (
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("http://buildwrangler.com")
	if err != nil {
		fmt.Println(resp, err)
	}

	fmt.Println("Status: ", resp.Status)

	bs := make([]byte, 99999)
	resp.Body.Read(bs)
	fmt.Println(string(bs))
}
