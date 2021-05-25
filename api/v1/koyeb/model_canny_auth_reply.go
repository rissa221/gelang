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

// CannyAuthReply struct for CannyAuthReply
type CannyAuthReply struct {
	Token *string `json:"token,omitempty"`
}

// NewCannyAuthReply instantiates a new CannyAuthReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCannyAuthReply() *CannyAuthReply {
	this := CannyAuthReply{}
	return &this
}

// NewCannyAuthReplyWithDefaults instantiates a new CannyAuthReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCannyAuthReplyWithDefaults() *CannyAuthReply {
	this := CannyAuthReply{}
	return &this
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *CannyAuthReply) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CannyAuthReply) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *CannyAuthReply) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *CannyAuthReply) SetToken(v string) {
	o.Token = &v
}

func (o CannyAuthReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	return json.Marshal(toSerialize)
}

type NullableCannyAuthReply struct {
	value *CannyAuthReply
	isSet bool
}

func (v NullableCannyAuthReply) Get() *CannyAuthReply {
	return v.value
}

func (v *NullableCannyAuthReply) Set(val *CannyAuthReply) {
	v.value = val
	v.isSet = true
}

func (v NullableCannyAuthReply) IsSet() bool {
	return v.isSet
}

func (v *NullableCannyAuthReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCannyAuthReply(val *CannyAuthReply) *NullableCannyAuthReply {
	return &NullableCannyAuthReply{value: val, isSet: true}
}

func (v NullableCannyAuthReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCannyAuthReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


