# \CamerasAPI

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppPrintersPrinterUuidCameraPost**](CamerasAPI.md#AppPrintersPrinterUuidCameraPost) | **Post** /app/printers/{printer_uuid}/camera | Register camera to Connect by user.



## AppPrintersPrinterUuidCameraPost

> CameraResponse AppPrintersPrinterUuidCameraPost(ctx, printerUuid).Origin(origin).Execute()

Register camera to Connect by user.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	printerUuid := "printerUuid_example" // string | 
	origin := "origin_example" // string |  (optional) (default to "WEB")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CamerasAPI.AppPrintersPrinterUuidCameraPost(context.Background(), printerUuid).Origin(origin).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CamerasAPI.AppPrintersPrinterUuidCameraPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppPrintersPrinterUuidCameraPost`: CameraResponse
	fmt.Fprintf(os.Stdout, "Response from `CamerasAPI.AppPrintersPrinterUuidCameraPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**printerUuid** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppPrintersPrinterUuidCameraPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **origin** | **string** |  | [default to &quot;WEB&quot;]

### Return type

[**CameraResponse**](CameraResponse.md)

### Authorization

[Cookie](../README.md#Cookie)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

