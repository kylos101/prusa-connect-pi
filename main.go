package main

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/pb33f/libopenapi"
	"github.com/pb33f/libopenapi/datamodel"
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
	config := datamodel.DocumentConfiguration{
		AllowFileReferences:   true,
		AllowRemoteReferences: true,
		BaseURL:               baseURL,
	}

	prusaConnect, err := os.ReadFile("./specs/prusaconnect.0.22.0.yaml")
	if err != nil {
		log.Fatalf("Error reading OpenAPI spec file: %v", err)
	}
	log.Print("Read OpenAPI spec file, creating new openapi document")
	document, err := libopenapi.NewDocumentWithConfiguration(prusaConnect, &config)
	if err != nil {
		log.Fatalf("Error creating new OpenAPI document from spec file: %v", err)
	}
	log.Print("Created new openapi document, building v3 model")
	docModel, errors := document.BuildV3Model()
	if len(errors) > 0 {
		for i := range errors {
			log.Printf("error: %e\n", errors[i])
		}
		log.Fatalf("cannot create v3 model from document: %d errors reported", len(errors))
	}

	paths := docModel.Model.Paths.PathItems.Len()
	schemas := docModel.Model.Components.Schemas.Len()

	// print the number of paths and schemas in the document
	fmt.Printf("There are %d paths and %d schemas "+
		"in the document", paths, schemas)

}
