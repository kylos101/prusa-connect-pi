# CameraConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CameraId** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** | Path to the driver | [optional] 
**Name** | **string** | Name defined by user | 
**Driver** | Pointer to **string** |  | [optional] 
**TriggerScheme** | Pointer to **string** | Type of snapshot trigger scheme. Manual, layer, gcode only for LINK cameras | [optional] [default to "THIRTY_SEC"]
**Resolution** | Pointer to [**CameraResolution**](CameraResolution.md) |  | [optional] 
**NetworkInfo** | Pointer to [**NetworkInfo**](NetworkInfo.md) |  | [optional] 
**Firmware** | Pointer to **string** | Firmware version of the camera | [optional] 
**Manufacturer** | Pointer to **string** | Manufacturer of the camera | [optional] 
**Model** | Pointer to **string** | Model of the camera | [optional] 

## Methods

### NewCameraConfig

`func NewCameraConfig(name string, ) *CameraConfig`

NewCameraConfig instantiates a new CameraConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCameraConfigWithDefaults

`func NewCameraConfigWithDefaults() *CameraConfig`

NewCameraConfigWithDefaults instantiates a new CameraConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCameraId

`func (o *CameraConfig) GetCameraId() string`

GetCameraId returns the CameraId field if non-nil, zero value otherwise.

### GetCameraIdOk

`func (o *CameraConfig) GetCameraIdOk() (*string, bool)`

GetCameraIdOk returns a tuple with the CameraId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCameraId

`func (o *CameraConfig) SetCameraId(v string)`

SetCameraId sets CameraId field to given value.

### HasCameraId

`func (o *CameraConfig) HasCameraId() bool`

HasCameraId returns a boolean if a field has been set.

### GetPath

`func (o *CameraConfig) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *CameraConfig) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *CameraConfig) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *CameraConfig) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetName

`func (o *CameraConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CameraConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CameraConfig) SetName(v string)`

SetName sets Name field to given value.


### GetDriver

`func (o *CameraConfig) GetDriver() string`

GetDriver returns the Driver field if non-nil, zero value otherwise.

### GetDriverOk

`func (o *CameraConfig) GetDriverOk() (*string, bool)`

GetDriverOk returns a tuple with the Driver field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDriver

`func (o *CameraConfig) SetDriver(v string)`

SetDriver sets Driver field to given value.

### HasDriver

`func (o *CameraConfig) HasDriver() bool`

HasDriver returns a boolean if a field has been set.

### GetTriggerScheme

`func (o *CameraConfig) GetTriggerScheme() string`

GetTriggerScheme returns the TriggerScheme field if non-nil, zero value otherwise.

### GetTriggerSchemeOk

`func (o *CameraConfig) GetTriggerSchemeOk() (*string, bool)`

GetTriggerSchemeOk returns a tuple with the TriggerScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerScheme

`func (o *CameraConfig) SetTriggerScheme(v string)`

SetTriggerScheme sets TriggerScheme field to given value.

### HasTriggerScheme

`func (o *CameraConfig) HasTriggerScheme() bool`

HasTriggerScheme returns a boolean if a field has been set.

### GetResolution

`func (o *CameraConfig) GetResolution() CameraResolution`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *CameraConfig) GetResolutionOk() (*CameraResolution, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *CameraConfig) SetResolution(v CameraResolution)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *CameraConfig) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### GetNetworkInfo

`func (o *CameraConfig) GetNetworkInfo() NetworkInfo`

GetNetworkInfo returns the NetworkInfo field if non-nil, zero value otherwise.

### GetNetworkInfoOk

`func (o *CameraConfig) GetNetworkInfoOk() (*NetworkInfo, bool)`

GetNetworkInfoOk returns a tuple with the NetworkInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInfo

`func (o *CameraConfig) SetNetworkInfo(v NetworkInfo)`

SetNetworkInfo sets NetworkInfo field to given value.

### HasNetworkInfo

`func (o *CameraConfig) HasNetworkInfo() bool`

HasNetworkInfo returns a boolean if a field has been set.

### GetFirmware

`func (o *CameraConfig) GetFirmware() string`

GetFirmware returns the Firmware field if non-nil, zero value otherwise.

### GetFirmwareOk

`func (o *CameraConfig) GetFirmwareOk() (*string, bool)`

GetFirmwareOk returns a tuple with the Firmware field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmware

`func (o *CameraConfig) SetFirmware(v string)`

SetFirmware sets Firmware field to given value.

### HasFirmware

`func (o *CameraConfig) HasFirmware() bool`

HasFirmware returns a boolean if a field has been set.

### GetManufacturer

`func (o *CameraConfig) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *CameraConfig) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *CameraConfig) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *CameraConfig) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetModel

`func (o *CameraConfig) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CameraConfig) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CameraConfig) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CameraConfig) HasModel() bool`

HasModel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


