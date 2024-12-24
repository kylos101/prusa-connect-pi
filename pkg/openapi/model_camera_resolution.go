/*
connect-api

# API Rules * Null values are not sent / received. * Object IDs are **always** sent in list methods, but are **ignored** in create / update methods. * All request and response objects are at the root of the returned structure, if they contain only one entity. * Response of multiple entities is returned as an object that contains the list of entities and a structure `pager`, if necessary. ### Additional documentation: * [Camera registration](../camera_registration/) * [Camera communication](../camera_communication/) ### HTTP Status * 200 - OK, response contains data (usually the entire entity) * 201 - OK, entry created; response contains data as required * 204 - OK, no response * 304 - Response has not been modified * 400 - Invalid request / invalid input (SDK error) * 401 - Endpoint is being accessed without credentials (SDK error) * 403 - Request can't be served, usually due to insufficient rights (user error) * 404 - Entity not found (user error or outdated data) * 409 - Conflict with the state of target resource (user error) * 50x - Server side error

API version: 0.22.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the CameraResolution type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CameraResolution{}

// CameraResolution struct for CameraResolution
type CameraResolution struct {
	Width *int32 `json:"width,omitempty"`
	Height *int32 `json:"height,omitempty"`
}

// NewCameraResolution instantiates a new CameraResolution object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCameraResolution() *CameraResolution {
	this := CameraResolution{}
	return &this
}

// NewCameraResolutionWithDefaults instantiates a new CameraResolution object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCameraResolutionWithDefaults() *CameraResolution {
	this := CameraResolution{}
	return &this
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *CameraResolution) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraResolution) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *CameraResolution) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *CameraResolution) SetWidth(v int32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *CameraResolution) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CameraResolution) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *CameraResolution) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *CameraResolution) SetHeight(v int32) {
	o.Height = &v
}

func (o CameraResolution) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CameraResolution) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	return toSerialize, nil
}

type NullableCameraResolution struct {
	value *CameraResolution
	isSet bool
}

func (v NullableCameraResolution) Get() *CameraResolution {
	return v.value
}

func (v *NullableCameraResolution) Set(val *CameraResolution) {
	v.value = val
	v.isSet = true
}

func (v NullableCameraResolution) IsSet() bool {
	return v.isSet
}

func (v *NullableCameraResolution) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCameraResolution(val *CameraResolution) *NullableCameraResolution {
	return &NullableCameraResolution{value: val, isSet: true}
}

func (v NullableCameraResolution) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCameraResolution) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

