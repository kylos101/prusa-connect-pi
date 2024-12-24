package main

import (
	"log"
	"net/url"
	"os"
)

func main() {
	apiUrl := os.Getenv("CONNECT_API_URL")
	if apiUrl == "" {
		log.Printf("Using default API URL")
		apiUrl = "https://connect.prusa3d.com/app/"
	}
	baseURL, err := url.Parse(apiUrl)
	if err != nil {
		log.Fatalf("Error parsing API URL: %v", err)
	}

	_, err = NewClient(baseURL)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

}

func NewClient(baseURL *url.URL) (any, error) {
	panic("unimplemented")
}
