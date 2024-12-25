package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
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
		apiUrl = "https://connect.prusa3d.com/app/"
		log.Printf("CONNECT_API_URL not specified. Using default API URL: %s\n", apiUrl)
	}
	baseURL, err := url.Parse(apiUrl)
	if err != nil {
		log.Fatalf("Error parsing API URL: %v\n", err)
	}

	token := os.Getenv("CONNECT_TOKEN")
	if token == "" {
		log.Fatal("CONNECT_TOKEN not provided!")
	}
	fingerPrint := getFingerprint()

	// set some server variables for openapi
	serverVariableContext := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
		"hostname": baseURL.Host,
		"basePath": baseURL.Path,
		"scheme":   baseURL.Scheme,
	})
	// define which api keys we'd like to use for openapi
	apiKeysContext := context.WithValue(serverVariableContext, openapi.ContextAPIKeys, map[string]string{
		"Token":       token,
		"Fingerprint": fingerPrint,
	})

	cancelCtx, cancel := context.WithCancel(apiKeysContext)

	connectInterval := os.Getenv("CONNECT_INTERVAL")
	interval := 5 * time.Minute
	if connectInterval == "" {
		log.Printf("CONNECT_INTERVAL not specified. Using default internval of %v.", interval)
	} else {
		customInterval, err := time.ParseDuration(connectInterval)
		if err != nil {
			log.Printf("CONNECT_INTERVAL parsing error: %v. Using default internval of 5m.", err)
		} else {
			interval = customInterval
		}
	}

	client := newClient(baseURL)
	defer cancel()
	UploadSnapshot(cancelCtx, client)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		UploadSnapshot(cancelCtx, client)
	}
}

func newClient(baseURL *url.URL) *openapi.APIClient {
	client := openapi.NewAPIClient(&openapi.Configuration{
		Host:   baseURL.String(),
		Scheme: baseURL.Scheme,
	})
	if client == nil {
		log.Fatal("Error creating client")
	}
	return client
}

func UploadSnapshot(ctx context.Context, client *openapi.APIClient) {
	log.Print("Capturing still")
	stillFilename := "output.jpg"
	cmd := exec.Command("rpicam-still", "-n", "-o", stillFilename)
	var err error
	defer func() {
		if err != nil {
			log.Printf("Unexpected error: %v\n", err)
		}
		log.Print("Removing output.jpg")
		if err := os.Remove("output.jpg"); err != nil {
			log.Printf("Error removing output.jpg: %v", err)
		}
	}()
	out, err := cmd.Output()
	if err != nil {
		err = fmt.Errorf("error running rpicam-still: %v", err)
		return
	}
	log.Printf("rpicam-still output: %v\n", out)

	log.Print("Assemble a request to upload snapshot")
	stillFile, err := os.Open(stillFilename)
	if err != nil {
		err = fmt.Errorf("error opening still file: %v", err)
		return
	}
	snapshotRequest := client.CameraAPI.CSnapshotPut(ctx).Body(stillFile)

	log.Print("Uploading snapshot")
	response, err := client.CameraAPI.CSnapshotPutExecute(snapshotRequest)
	if err != nil {
		err = fmt.Errorf("error uploading snapshot: %+v", err)
		return
	}
	if response.StatusCode != 200 {
		log.Printf("Unsuccessful snapshot upload: %v. Status code and status: %d-%v\n", response, response.StatusCode, response.Status)
	}
}

func getFingerprint() string {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}
	hash := sha256.Sum256([]byte(hostname))
	hashStr := hex.EncodeToString(hash[:])
	return hashStr[:16]

}
