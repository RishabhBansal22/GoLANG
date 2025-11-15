package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.gettyimage.me")
	if err != nil {
		// handle error and exit
		fmt.Println("error fetching:", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error reading body
		fmt.Println("error reading response body:", err)
		return
	}
	fmt.Println(string(body))
}
