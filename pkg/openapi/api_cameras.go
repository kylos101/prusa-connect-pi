/*
connect-api

# API Rules * Null values are not sent / received. * Object IDs are **always** sent in list methods, but are **ignored** in create / update methods. * All request and response objects are at the root of the returned structure, if they contain only one entity. * Response of multiple entities is returned as an object that contains the list of entities and a structure `pager`, if necessary. ### Additional documentation: * [Camera registration](../camera_registration/) * [Camera communication](../camera_communication/) ### HTTP Status * 200 - OK, response contains data (usually the entire entity) * 201 - OK, entry created; response contains data as required * 204 - OK, no response * 304 - Response has not been modified * 400 - Invalid request / invalid input (SDK error) * 401 - Endpoint is being accessed without credentials (SDK error) * 403 - Request can't be served, usually due to insufficient rights (user error) * 404 - Entity not found (user error or outdated data) * 409 - Conflict with the state of target resource (user error) * 50x - Server side error

API version: 0.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// CamerasAPIService CamerasAPI service
type CamerasAPIService service

type ApiAppPrintersPrinterUuidCameraPostRequest struct {
	ctx context.Context
	ApiService *CamerasAPIService
	printerUuid string
	origin *string
}

func (r ApiAppPrintersPrinterUuidCameraPostRequest) Origin(origin string) ApiAppPrintersPrinterUuidCameraPostRequest {
	r.origin = &origin
	return r
}

func (r ApiAppPrintersPrinterUuidCameraPostRequest) Execute() (*CameraResponse, *http.Response, error) {
	return r.ApiService.AppPrintersPrinterUuidCameraPostExecute(r)
}

/*
AppPrintersPrinterUuidCameraPost Register camera to Connect by user.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param printerUuid
 @return ApiAppPrintersPrinterUuidCameraPostRequest
*/
func (a *CamerasAPIService) AppPrintersPrinterUuidCameraPost(ctx context.Context, printerUuid string) ApiAppPrintersPrinterUuidCameraPostRequest {
	return ApiAppPrintersPrinterUuidCameraPostRequest{
		ApiService: a,
		ctx: ctx,
		printerUuid: printerUuid,
	}
}

// Execute executes the request
//  @return CameraResponse
func (a *CamerasAPIService) AppPrintersPrinterUuidCameraPostExecute(r ApiAppPrintersPrinterUuidCameraPostRequest) (*CameraResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CameraResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CamerasAPIService.AppPrintersPrinterUuidCameraPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/app/printers/{printer_uuid}/camera"
	localVarPath = strings.Replace(localVarPath, "{"+"printer_uuid"+"}", url.PathEscape(parameterValueToString(r.printerUuid, "printerUuid")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.origin != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "origin", r.origin, "form", "")
	} else {
		var defaultValue string = "WEB"
		r.origin = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json", "text/plain"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Status
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Status
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 503 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}