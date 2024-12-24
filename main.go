package main

import (
	"context"
	"log"
	"net/url"
	"os"

	"github.com/kylos101/prusa-connect/v2/pkg/openapi"
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

	client := NewClient(baseURL)
	if client == nil {
		log.Fatalf("Error creating client")
	}

	ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
		"hostname": baseURL.Host,
		"basePath": baseURL.Path,
		"scheme":   baseURL.Scheme,
	})

}

func NewClient(baseURL *url.URL) *openapi.APIClient {
	c := openapi.NewAPIClient(&openapi.Configuration{
		Host:   baseURL.String(),
		Scheme: baseURL.Scheme,
	})
	return c
}
