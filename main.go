package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Starting application!")

	url := "https://itunes.apple.com/lookup?id=1736401501"
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error making HTTP request: %s", err)
	}
	defer res.Body.Close()

	fmt.Printf("Got status code: %s", res.Status)

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Printf("Error reading response body: %s", err)
	}

	fmt.Printf("HTTP Response: %s", string(body))

}
