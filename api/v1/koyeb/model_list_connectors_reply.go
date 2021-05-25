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

// ListConnectorsReply struct for ListConnectorsReply
type ListConnectorsReply struct {
	Connectors *[]ConnectorListItem `json:"connectors,omitempty"`
	Limit *int64 `json:"limit,omitempty"`
	Offset *int64 `json:"offset,omitempty"`
	Count *int64 `json:"count,omitempty"`
}

// NewListConnectorsReply instantiates a new ListConnectorsReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListConnectorsReply() *ListConnectorsReply {
	this := ListConnectorsReply{}
	return &this
}

// NewListConnectorsReplyWithDefaults instantiates a new ListConnectorsReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListConnectorsReplyWithDefaults() *ListConnectorsReply {
	this := ListConnectorsReply{}
	return &this
}

// GetConnectors returns the Connectors field value if set, zero value otherwise.
func (o *ListConnectorsReply) GetConnectors() []ConnectorListItem {
	if o == nil || o.Connectors == nil {
		var ret []ConnectorListItem
		return ret
	}
	return *o.Connectors
}

// GetConnectorsOk returns a tuple with the Connectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectorsReply) GetConnectorsOk() (*[]ConnectorListItem, bool) {
	if o == nil || o.Connectors == nil {
		return nil, false
	}
	return o.Connectors, true
}

// HasConnectors returns a boolean if a field has been set.
func (o *ListConnectorsReply) HasConnectors() bool {
	if o != nil && o.Connectors != nil {
		return true
	}

	return false
}

// SetConnectors gets a reference to the given []ConnectorListItem and assigns it to the Connectors field.
func (o *ListConnectorsReply) SetConnectors(v []ConnectorListItem) {
	o.Connectors = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ListConnectorsReply) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectorsReply) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListConnectorsReply) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ListConnectorsReply) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ListConnectorsReply) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectorsReply) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ListConnectorsReply) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *ListConnectorsReply) SetOffset(v int64) {
	o.Offset = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListConnectorsReply) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListConnectorsReply) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListConnectorsReply) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ListConnectorsReply) SetCount(v int64) {
	o.Count = &v
}

func (o ListConnectorsReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Connectors != nil {
		toSerialize["connectors"] = o.Connectors
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.Offset != nil {
		toSerialize["offset"] = o.Offset
	}
	if o.Count != nil {
		toSerialize["count"] = o.Count
	}
	return json.Marshal(toSerialize)
}

type NullableListConnectorsReply struct {
	value *ListConnectorsReply
	isSet bool
}

func (v NullableListConnectorsReply) Get() *ListConnectorsReply {
	return v.value
}

func (v *NullableListConnectorsReply) Set(val *ListConnectorsReply) {
	v.value = val
	v.isSet = true
}

func (v NullableListConnectorsReply) IsSet() bool {
	return v.isSet
}

func (v *NullableListConnectorsReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListConnectorsReply(val *ListConnectorsReply) *NullableListConnectorsReply {
	return &NullableListConnectorsReply{value: val, isSet: true}
}

func (v NullableListConnectorsReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListConnectorsReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


