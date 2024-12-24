# CameraResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int32** |  | 
**Name** | Pointer to **string** | Name defined by user | [optional] 
**Config** | Pointer to [**CameraConfig**](CameraConfig.md) |  | [optional] 
**Options** | Pointer to [**CameraOptions**](CameraOptions.md) |  | [optional] 
**Capabilities** | Pointer to **[]string** |  | [optional] 
**TeamId** | **int32** | Team id. 0 is reserved for anonymous. | 
**PrinterUuid** | **string** | Printer UUID - can be found in the url in printer detail or in the printer settings tab | 
**Token** | Pointer to **string** |  | [optional] 
**Origin** | [**CameraOrigin**](CameraOrigin.md) |  | 
**Deleted** | Pointer to **int32** | Timestamp of camera deletion. Deleted cameras are listed only in job detail! | [optional] 
**Registered** | **bool** | True if the registration process of camera is finished | 
**SortOrder** | Pointer to **int32** | Sort order of the cameras per printer! For now it is not possible to change the camera&#39;s sort order | [optional] 

## Methods

### NewCameraResponse

`func NewCameraResponse(id int32, teamId int32, printerUuid string, origin CameraOrigin, registered bool, ) *CameraResponse`

NewCameraResponse instantiates a new CameraResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCameraResponseWithDefaults

`func NewCameraResponseWithDefaults() *CameraResponse`

NewCameraResponseWithDefaults instantiates a new CameraResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CameraResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CameraResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CameraResponse) SetId(v int32)`

SetId sets Id field to given value.


### GetName

`func (o *CameraResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CameraResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CameraResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CameraResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetConfig

`func (o *CameraResponse) GetConfig() CameraConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CameraResponse) GetConfigOk() (*CameraConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CameraResponse) SetConfig(v CameraConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CameraResponse) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetOptions

`func (o *CameraResponse) GetOptions() CameraOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CameraResponse) GetOptionsOk() (*CameraOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CameraResponse) SetOptions(v CameraOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CameraResponse) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetCapabilities

`func (o *CameraResponse) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CameraResponse) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CameraResponse) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *CameraResponse) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetTeamId

`func (o *CameraResponse) GetTeamId() int32`

GetTeamId returns the TeamId field if non-nil, zero value otherwise.

### GetTeamIdOk

`func (o *CameraResponse) GetTeamIdOk() (*int32, bool)`

GetTeamIdOk returns a tuple with the TeamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTeamId

`func (o *CameraResponse) SetTeamId(v int32)`

SetTeamId sets TeamId field to given value.


### GetPrinterUuid

`func (o *CameraResponse) GetPrinterUuid() string`

GetPrinterUuid returns the PrinterUuid field if non-nil, zero value otherwise.

### GetPrinterUuidOk

`func (o *CameraResponse) GetPrinterUuidOk() (*string, bool)`

GetPrinterUuidOk returns a tuple with the PrinterUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrinterUuid

`func (o *CameraResponse) SetPrinterUuid(v string)`

SetPrinterUuid sets PrinterUuid field to given value.


### GetToken

`func (o *CameraResponse) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CameraResponse) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CameraResponse) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CameraResponse) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetOrigin

`func (o *CameraResponse) GetOrigin() CameraOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *CameraResponse) GetOriginOk() (*CameraOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *CameraResponse) SetOrigin(v CameraOrigin)`

SetOrigin sets Origin field to given value.


### GetDeleted

`func (o *CameraResponse) GetDeleted() int32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *CameraResponse) GetDeletedOk() (*int32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *CameraResponse) SetDeleted(v int32)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *CameraResponse) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetRegistered

`func (o *CameraResponse) GetRegistered() bool`

GetRegistered returns the Registered field if non-nil, zero value otherwise.

### GetRegisteredOk

`func (o *CameraResponse) GetRegisteredOk() (*bool, bool)`

GetRegisteredOk returns a tuple with the Registered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistered

`func (o *CameraResponse) SetRegistered(v bool)`

SetRegistered sets Registered field to given value.


### GetSortOrder

`func (o *CameraResponse) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *CameraResponse) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *CameraResponse) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *CameraResponse) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


