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

// UpdateStackReply struct for UpdateStackReply
type UpdateStackReply struct {
	Stack *Stack `json:"stack,omitempty"`
}

// NewUpdateStackReply instantiates a new UpdateStackReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateStackReply() *UpdateStackReply {
	this := UpdateStackReply{}
	return &this
}

// NewUpdateStackReplyWithDefaults instantiates a new UpdateStackReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateStackReplyWithDefaults() *UpdateStackReply {
	this := UpdateStackReply{}
	return &this
}

// GetStack returns the Stack field value if set, zero value otherwise.
func (o *UpdateStackReply) GetStack() Stack {
	if o == nil || o.Stack == nil {
		var ret Stack
		return ret
	}
	return *o.Stack
}

// GetStackOk returns a tuple with the Stack field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateStackReply) GetStackOk() (*Stack, bool) {
	if o == nil || o.Stack == nil {
		return nil, false
	}
	return o.Stack, true
}

// HasStack returns a boolean if a field has been set.
func (o *UpdateStackReply) HasStack() bool {
	if o != nil && o.Stack != nil {
		return true
	}

	return false
}

// SetStack gets a reference to the given Stack and assigns it to the Stack field.
func (o *UpdateStackReply) SetStack(v Stack) {
	o.Stack = &v
}

func (o UpdateStackReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Stack != nil {
		toSerialize["stack"] = o.Stack
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateStackReply struct {
	value *UpdateStackReply
	isSet bool
}

func (v NullableUpdateStackReply) Get() *UpdateStackReply {
	return v.value
}

func (v *NullableUpdateStackReply) Set(val *UpdateStackReply) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateStackReply) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateStackReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateStackReply(val *UpdateStackReply) *NullableUpdateStackReply {
	return &NullableUpdateStackReply{value: val, isSet: true}
}

func (v NullableUpdateStackReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateStackReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

