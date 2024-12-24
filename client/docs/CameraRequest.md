# CameraRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**CameraConfig**](CameraConfig.md) |  | 
**Options** | Pointer to [**CameraOptions**](CameraOptions.md) |  | [optional] 
**Capabilities** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCameraRequest

`func NewCameraRequest(config CameraConfig, ) *CameraRequest`

NewCameraRequest instantiates a new CameraRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCameraRequestWithDefaults

`func NewCameraRequestWithDefaults() *CameraRequest`

NewCameraRequestWithDefaults instantiates a new CameraRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *CameraRequest) GetConfig() CameraConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CameraRequest) GetConfigOk() (*CameraConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CameraRequest) SetConfig(v CameraConfig)`

SetConfig sets Config field to given value.


### GetOptions

`func (o *CameraRequest) GetOptions() CameraOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CameraRequest) GetOptionsOk() (*CameraOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CameraRequest) SetOptions(v CameraOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CameraRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetCapabilities

`func (o *CameraRequest) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *CameraRequest) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *CameraRequest) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *CameraRequest) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


