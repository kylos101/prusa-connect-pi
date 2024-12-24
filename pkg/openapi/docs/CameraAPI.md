# \CameraAPI

All URIs are relative to *http://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CInfoPut**](CameraAPI.md#CInfoPut) | **Put** /c/info | Update camera attributes
[**CSnapshotPut**](CameraAPI.md#CSnapshotPut) | **Put** /c/snapshot | Upload snapshot to Connect



## CInfoPut

> CameraResponse CInfoPut(ctx).CameraRequest(cameraRequest).Execute()

Update camera attributes

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
	cameraRequest := *openapiclient.NewCameraRequest(*openapiclient.NewCameraConfig("Olomouc")) // CameraRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CameraAPI.CInfoPut(context.Background()).CameraRequest(cameraRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.CInfoPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CInfoPut`: CameraResponse
	fmt.Fprintf(os.Stdout, "Response from `CameraAPI.CInfoPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCInfoPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cameraRequest** | [**CameraRequest**](CameraRequest.md) |  | 

### Return type

[**CameraResponse**](CameraResponse.md)

### Authorization

[Fingerprint](../README.md#Fingerprint), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CSnapshotPut

> CSnapshotPut(ctx).Body(body).Execute()

Upload snapshot to Connect

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
	body := os.NewFile(1234, "some_file") // *os.File | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.CameraAPI.CSnapshotPut(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CameraAPI.CSnapshotPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCSnapshotPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

[Fingerprint](../README.md#Fingerprint), [Token](../README.md#Token)

### HTTP request headers

- **Content-Type**: image/jpg
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

