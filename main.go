package main

import (
	"context"
	"log"
	"net/url"
	"os"
	"os/exec"
	"time"

	"github.com/kylos101/prusa-connect-pi/v2/pkg/openapi"
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

	client := openapi.NewAPIClient(&openapi.Configuration{
		Host:   baseURL.String(),
		Scheme: baseURL.Scheme,
	})
	if client == nil {
		log.Fatalf("Error creating client")
	}

	ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
		"hostname": baseURL.Host,
		"basePath": baseURL.Path,
		"scheme":   baseURL.Scheme,
	})
	cancelCtx, cancel := context.WithCancel(ctx)

	token := os.Getenv("CONNECT_TOKEN")
	if token == "" {
		log.Fatal("No token provided")
	}

	connectInterval := os.Getenv("CONNECT_INTERVAL")
	defaultInterval := 5 * time.Minute

	interval, err := time.ParseDuration(connectInterval)
	if err != nil {
		log.Printf("Error parsing interval: %v, using default of 5m", err)
		interval = defaultInterval
	}

	UploadSnapshot(cancelCtx, client, token)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	defer cancel()
	for range ticker.C {
		UploadSnapshot(cancelCtx, client, token)
	}
}

func UploadSnapshot(ctx context.Context, client *openapi.APIClient, token string) {
	// capture a still
	stillFilename := "output.jpg"
	cmd := exec.Command("rpicam-still", "-n", "-o", stillFilename)
	defer func() {
		log.Printf("Removing output.jpg")
		if err := os.Remove("output.jpg"); err != nil {
			log.Printf("Error removing output.jpg: %v", err)
		}
	}()

	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("Error running rpicam-still: %v", err)
	}
	log.Printf("rpicam-still output: %s/n", out)

	// build a request to upload the still
	stillFile, err := os.Open(stillFilename)
	if err != nil {
		log.Fatalf("Error opening still file: %v", err)
	}

	// prepare the client to auth
	client.GetConfig().AddDefaultHeader("Token", token)
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	client.GetConfig().AddDefaultHeader("Fingerprint", hostname)

	//assemble the request
	snapshotRequest := client.CameraAPI.CSnapshotPut(ctx)
	snapshotRequest.Body(stillFile)

	// upload
	response, err := client.CameraAPI.CSnapshotPutExecute(snapshotRequest)
	if err != nil {
		log.Fatalf("Error uploading snapshot: %v", err)
	}
	if response.StatusCode != 200 {
		log.Printf("Unsuccessful snapshot upload: %v", response)
	}

}
