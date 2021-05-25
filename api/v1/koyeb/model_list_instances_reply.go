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

// ListInstancesReply struct for ListInstancesReply
type ListInstancesReply struct {
	Instances *[]InstanceListItem `json:"instances,omitempty"`
	Limit *int64 `json:"limit,omitempty"`
	Offset *int64 `json:"offset,omitempty"`
	Count *int64 `json:"count,omitempty"`
}

// NewListInstancesReply instantiates a new ListInstancesReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInstancesReply() *ListInstancesReply {
	this := ListInstancesReply{}
	return &this
}

// NewListInstancesReplyWithDefaults instantiates a new ListInstancesReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInstancesReplyWithDefaults() *ListInstancesReply {
	this := ListInstancesReply{}
	return &this
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *ListInstancesReply) GetInstances() []InstanceListItem {
	if o == nil || o.Instances == nil {
		var ret []InstanceListItem
		return ret
	}
	return *o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstancesReply) GetInstancesOk() (*[]InstanceListItem, bool) {
	if o == nil || o.Instances == nil {
		return nil, false
	}
	return o.Instances, true
}

// HasInstances returns a boolean if a field has been set.
func (o *ListInstancesReply) HasInstances() bool {
	if o != nil && o.Instances != nil {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []InstanceListItem and assigns it to the Instances field.
func (o *ListInstancesReply) SetInstances(v []InstanceListItem) {
	o.Instances = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *ListInstancesReply) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstancesReply) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *ListInstancesReply) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *ListInstancesReply) SetLimit(v int64) {
	o.Limit = &v
}

// GetOffset returns the Offset field value if set, zero value otherwise.
func (o *ListInstancesReply) GetOffset() int64 {
	if o == nil || o.Offset == nil {
		var ret int64
		return ret
	}
	return *o.Offset
}

// GetOffsetOk returns a tuple with the Offset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstancesReply) GetOffsetOk() (*int64, bool) {
	if o == nil || o.Offset == nil {
		return nil, false
	}
	return o.Offset, true
}

// HasOffset returns a boolean if a field has been set.
func (o *ListInstancesReply) HasOffset() bool {
	if o != nil && o.Offset != nil {
		return true
	}

	return false
}

// SetOffset gets a reference to the given int64 and assigns it to the Offset field.
func (o *ListInstancesReply) SetOffset(v int64) {
	o.Offset = &v
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *ListInstancesReply) GetCount() int64 {
	if o == nil || o.Count == nil {
		var ret int64
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstancesReply) GetCountOk() (*int64, bool) {
	if o == nil || o.Count == nil {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *ListInstancesReply) HasCount() bool {
	if o != nil && o.Count != nil {
		return true
	}

	return false
}

// SetCount gets a reference to the given int64 and assigns it to the Count field.
func (o *ListInstancesReply) SetCount(v int64) {
	o.Count = &v
}

func (o ListInstancesReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Instances != nil {
		toSerialize["instances"] = o.Instances
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

type NullableListInstancesReply struct {
	value *ListInstancesReply
	isSet bool
}

func (v NullableListInstancesReply) Get() *ListInstancesReply {
	return v.value
}

func (v *NullableListInstancesReply) Set(val *ListInstancesReply) {
	v.value = val
	v.isSet = true
}

func (v NullableListInstancesReply) IsSet() bool {
	return v.isSet
}

func (v *NullableListInstancesReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListInstancesReply(val *ListInstancesReply) *NullableListInstancesReply {
	return &NullableListInstancesReply{value: val, isSet: true}
}

func (v NullableListInstancesReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListInstancesReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


