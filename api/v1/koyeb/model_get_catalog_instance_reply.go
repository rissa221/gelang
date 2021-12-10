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

// GetCatalogInstanceReply struct for GetCatalogInstanceReply
type GetCatalogInstanceReply struct {
	Instance *CatalogInstance `json:"instance,omitempty"`
}

// NewGetCatalogInstanceReply instantiates a new GetCatalogInstanceReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetCatalogInstanceReply() *GetCatalogInstanceReply {
	this := GetCatalogInstanceReply{}
	return &this
}

// NewGetCatalogInstanceReplyWithDefaults instantiates a new GetCatalogInstanceReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetCatalogInstanceReplyWithDefaults() *GetCatalogInstanceReply {
	this := GetCatalogInstanceReply{}
	return &this
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *GetCatalogInstanceReply) GetInstance() CatalogInstance {
	if o == nil || o.Instance == nil {
		var ret CatalogInstance
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetCatalogInstanceReply) GetInstanceOk() (*CatalogInstance, bool) {
	if o == nil || o.Instance == nil {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *GetCatalogInstanceReply) HasInstance() bool {
	if o != nil && o.Instance != nil {
		return true
	}

	return false
}

// SetInstance gets a reference to the given CatalogInstance and assigns it to the Instance field.
func (o *GetCatalogInstanceReply) SetInstance(v CatalogInstance) {
	o.Instance = &v
}

func (o GetCatalogInstanceReply) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Instance != nil {
		toSerialize["instance"] = o.Instance
	}
	return json.Marshal(toSerialize)
}

type NullableGetCatalogInstanceReply struct {
	value *GetCatalogInstanceReply
	isSet bool
}

func (v NullableGetCatalogInstanceReply) Get() *GetCatalogInstanceReply {
	return v.value
}

func (v *NullableGetCatalogInstanceReply) Set(val *GetCatalogInstanceReply) {
	v.value = val
	v.isSet = true
}

func (v NullableGetCatalogInstanceReply) IsSet() bool {
	return v.isSet
}

func (v *NullableGetCatalogInstanceReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetCatalogInstanceReply(val *GetCatalogInstanceReply) *NullableGetCatalogInstanceReply {
	return &NullableGetCatalogInstanceReply{value: val, isSet: true}
}

func (v NullableGetCatalogInstanceReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetCatalogInstanceReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

