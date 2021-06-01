/*
 * Koyeb Rest API
 *
 * The Koyeb API allows you to interact with the Koyeb platform in a simple, programmatic way using conventional HTTP requests. 
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package koyeb

import (
	"encoding/json"
)

// GoogleRpcStatus struct for GoogleRpcStatus
type GoogleRpcStatus struct {
	Code *int32 `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
	Details *[]GoogleProtobufAny `json:"details,omitempty"`
}

// NewGoogleRpcStatus instantiates a new GoogleRpcStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGoogleRpcStatus() *GoogleRpcStatus {
	this := GoogleRpcStatus{}
	return &this
}

// NewGoogleRpcStatusWithDefaults instantiates a new GoogleRpcStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGoogleRpcStatusWithDefaults() *GoogleRpcStatus {
	this := GoogleRpcStatus{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *GoogleRpcStatus) GetCode() int32 {
	if o == nil || o.Code == nil {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleRpcStatus) GetCodeOk() (*int32, bool) {
	if o == nil || o.Code == nil {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *GoogleRpcStatus) HasCode() bool {
	if o != nil && o.Code != nil {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *GoogleRpcStatus) SetCode(v int32) {
	o.Code = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GoogleRpcStatus) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleRpcStatus) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GoogleRpcStatus) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GoogleRpcStatus) SetMessage(v string) {
	o.Message = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *GoogleRpcStatus) GetDetails() []GoogleProtobufAny {
	if o == nil || o.Details == nil {
		var ret []GoogleProtobufAny
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GoogleRpcStatus) GetDetailsOk() (*[]GoogleProtobufAny, bool) {
	if o == nil || o.Details == nil {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *GoogleRpcStatus) HasDetails() bool {
	if o != nil && o.Details != nil {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []GoogleProtobufAny and assigns it to the Details field.
func (o *GoogleRpcStatus) SetDetails(v []GoogleProtobufAny) {
	o.Details = &v
}

func (o GoogleRpcStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Code != nil {
		toSerialize["code"] = o.Code
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Details != nil {
		toSerialize["details"] = o.Details
	}
	return json.Marshal(toSerialize)
}

type NullableGoogleRpcStatus struct {
	value *GoogleRpcStatus
	isSet bool
}

func (v NullableGoogleRpcStatus) Get() *GoogleRpcStatus {
	return v.value
}

func (v *NullableGoogleRpcStatus) Set(val *GoogleRpcStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGoogleRpcStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGoogleRpcStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGoogleRpcStatus(val *GoogleRpcStatus) *NullableGoogleRpcStatus {
	return &NullableGoogleRpcStatus{value: val, isSet: true}
}

func (v NullableGoogleRpcStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGoogleRpcStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

