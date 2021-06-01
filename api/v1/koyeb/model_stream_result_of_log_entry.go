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

// StreamResultOfLogEntry struct for StreamResultOfLogEntry
type StreamResultOfLogEntry struct {
	Result *LogEntry `json:"result,omitempty"`
	Error *GoogleRpcStatus `json:"error,omitempty"`
}

// NewStreamResultOfLogEntry instantiates a new StreamResultOfLogEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStreamResultOfLogEntry() *StreamResultOfLogEntry {
	this := StreamResultOfLogEntry{}
	return &this
}

// NewStreamResultOfLogEntryWithDefaults instantiates a new StreamResultOfLogEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStreamResultOfLogEntryWithDefaults() *StreamResultOfLogEntry {
	this := StreamResultOfLogEntry{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *StreamResultOfLogEntry) GetResult() LogEntry {
	if o == nil || o.Result == nil {
		var ret LogEntry
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamResultOfLogEntry) GetResultOk() (*LogEntry, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *StreamResultOfLogEntry) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given LogEntry and assigns it to the Result field.
func (o *StreamResultOfLogEntry) SetResult(v LogEntry) {
	o.Result = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *StreamResultOfLogEntry) GetError() GoogleRpcStatus {
	if o == nil || o.Error == nil {
		var ret GoogleRpcStatus
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StreamResultOfLogEntry) GetErrorOk() (*GoogleRpcStatus, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *StreamResultOfLogEntry) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given GoogleRpcStatus and assigns it to the Error field.
func (o *StreamResultOfLogEntry) SetError(v GoogleRpcStatus) {
	o.Error = &v
}

func (o StreamResultOfLogEntry) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableStreamResultOfLogEntry struct {
	value *StreamResultOfLogEntry
	isSet bool
}

func (v NullableStreamResultOfLogEntry) Get() *StreamResultOfLogEntry {
	return v.value
}

func (v *NullableStreamResultOfLogEntry) Set(val *StreamResultOfLogEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableStreamResultOfLogEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableStreamResultOfLogEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStreamResultOfLogEntry(val *StreamResultOfLogEntry) *NullableStreamResultOfLogEntry {
	return &NullableStreamResultOfLogEntry{value: val, isSet: true}
}

func (v NullableStreamResultOfLogEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStreamResultOfLogEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

