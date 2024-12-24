# CameraOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableResolutions** | Pointer to [**[]CameraResolution**](CameraResolution.md) |  | [optional] 

## Methods

### NewCameraOptions

`func NewCameraOptions() *CameraOptions`

NewCameraOptions instantiates a new CameraOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCameraOptionsWithDefaults

`func NewCameraOptionsWithDefaults() *CameraOptions`

NewCameraOptionsWithDefaults instantiates a new CameraOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableResolutions

`func (o *CameraOptions) GetAvailableResolutions() []CameraResolution`

GetAvailableResolutions returns the AvailableResolutions field if non-nil, zero value otherwise.

### GetAvailableResolutionsOk

`func (o *CameraOptions) GetAvailableResolutionsOk() (*[]CameraResolution, bool)`

GetAvailableResolutionsOk returns a tuple with the AvailableResolutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableResolutions

`func (o *CameraOptions) SetAvailableResolutions(v []CameraResolution)`

SetAvailableResolutions sets AvailableResolutions field to given value.

### HasAvailableResolutions

`func (o *CameraOptions) HasAvailableResolutions() bool`

HasAvailableResolutions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


