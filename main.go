package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kylos101/prusa-connect/v2/connect"
)

func main() {
	// Replace with your actual base URL
	baseURL := "http://localhost:8000"

	client, err := connect.NewClient(baseURL)
	if err != nil {
		log.Fatalf("Error creating client: %v", err)
	}

	// Set authentication if needed
	client.SetAuth(os.Getenv("CONNECT_TOKEN"), os.Getenv("CONNECT_FINGERPRINT"), os.Getenv("CONNECT_COOKIE"))

	// Example: Register a camera
	printerUUID := "cfed5dce-86f4-4d7c-a198-9a81b176369f" // Replace with a valid UUID
	origin := "OTHER"

	cameraResponse, err := client.RegisterCamera(printerUUID, origin)
	if err != nil {
		log.Printf("Error registering camera: %v", err)
		// If you want to see the raw response for debugging:
		// fmt.Println(string(err.(*connect.APIError).Body))
		os.Exit(1)
	}

	fmt.Printf("Registered Camera ID: %d\n", cameraResponse.ID)
	fmt.Printf("Registered Camera Token: %s\n", cameraResponse.Token)

	// Example: Upload a snapshot (replace with actual image data)
	snapshotData, err := os.ReadFile("path/to/your/snapshot.jpg") // Replace with the path to your image
	if err != nil {
		log.Fatalf("Error reading snapshot file: %v", err)
	}

	err = client.UploadSnapshot(snapshotData)
	if err != nil {
		log.Fatalf("Error uploading snapshot: %v", err)
	}

	fmt.Println("Snapshot uploaded successfully")

	// Example: Update Camera Info
	config := connect.Config{
		Name:   "Test Camera",
		Driver: "V4L2",
		Path:   "/dev/video1",
	}
	options := connect.Options{}
	capabilities := []string{"resolution"}
	cameraRequest := connect.CameraRequest{
		Config:       &config,
		Options:      &options,
		Capabilities: capabilities,
	}

	updatedCamera, err := client.UpdateCamera(cameraRequest)
	if err != nil {
		log.Fatalf("Error updating camera: %v", err)
	}

	fmt.Printf("Updated Camera Name: %s\n", updatedCamera.Name)

}
