package main

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
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
		log.Fatalf("Error parsing API URL: %v\n", err)
	}

	token := os.Getenv("CONNECT_TOKEN")
	if token == "" {
		log.Fatal("No token provided")
	}

	ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
		"hostname": baseURL.Host,
		"basePath": baseURL.Path,
		"scheme":   baseURL.Scheme,
	})
	cancelCtx, cancel := context.WithCancel(ctx)

	connectInterval := os.Getenv("CONNECT_INTERVAL")
	defaultInterval := 5 * time.Minute

	interval, err := time.ParseDuration(connectInterval)
	if err != nil {
		log.Printf("Unable to parse connect interval. Error: %v. Using default internval of 5m", err)
		interval = defaultInterval
	}

	client := newClient(baseURL, token)
	UploadSnapshot(cancelCtx, client)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	defer cancel()
	for range ticker.C {
		UploadSnapshot(cancelCtx, client)
	}
}

func newClient(baseURL *url.URL, token string) *openapi.APIClient {
	defaultHeaders := map[string]string{
		"Token":       token,
		"Fingerprint": getFingerprint(),
	}

	client := openapi.NewAPIClient(&openapi.Configuration{
		Host:          baseURL.String(),
		Scheme:        baseURL.Scheme,
		DefaultHeader: defaultHeaders,
	})
	if client == nil {
		log.Fatal("Error creating client")
	}
	return client
}

func UploadSnapshot(ctx context.Context, client *openapi.APIClient) {
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
		log.Fatalf("Error running rpicam-still: %v\n", err)
	}
	log.Printf("rpicam-still output: %s\n", out)

	// build a request to upload the still
	stillFile, err := os.Open(stillFilename)
	if err != nil {
		log.Fatalf("Error opening still file: %v\n", err)
	}

	//assemble the request
	snapshotRequest := client.CameraAPI.CSnapshotPut(ctx)
	snapshotRequest.Body(stillFile)

	// upload
	response, err := client.CameraAPI.CSnapshotPutExecute(snapshotRequest)
	if err != nil {
		log.Fatalf("Error uploading snapshot. Error: %v,\n", err)
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
